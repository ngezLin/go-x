package configloader

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	vault "github.com/hashicorp/vault/api"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

type (
	ViperLoader struct {
		*viper.Viper
		consulKey             string
		consulURL             string
		vaultURL              string
		vaultToken            string
		vaultKey              string
		envPrefix             string
		configFileName        string
		remoteMaxAttempt      int
		tagName               string
		configFileSearchPaths []string
	}
	Option func(*ViperLoader)
)

var (
	ErrInvalidInput       = fmt.Errorf("cfg must be a pointer to struct and initialized")
	ErrConfigFileNotFound = fmt.Errorf("config file not found")
)

// WithConfigFileSearchPaths will add paths to where the loader will search the configuration files
func WithConfigFileSearchPaths(paths ...string) Option {
	return func(v *ViperLoader) {
		if len(paths) > 0 {
			v.configFileSearchPaths = append(v.configFileSearchPaths, paths...)
		}
	}
}

func WithLoadFromConsulMaxAttempt(n int) Option {
	return func(v *ViperLoader) {
		if n > 0 {
			v.remoteMaxAttempt = n
		}
	}
}

func WithConfigFileName(fileName string) Option {
	return func(v *ViperLoader) {
		v.configFileName = fileName
	}
}

func WithStructTagName(name string) Option {
	return func(v *ViperLoader) {
		v.tagName = name
	}
}

func WithConsule(url, key string) Option {
	return func(v *ViperLoader) {
		v.consulURL = url
		v.consulKey = key
	}
}

func WithVault(url, key, token string) Option {
	return func(v *ViperLoader) {
		v.vaultURL = url
		v.vaultToken = token
		v.vaultKey = key
	}
}

func New(envPrefix string, opts ...Option) *ViperLoader {
	v := &ViperLoader{
		Viper:                 viper.New(),
		configFileName:        "config",
		remoteMaxAttempt:      5,
		tagName:               "json",
		envPrefix:             envPrefix,
		configFileSearchPaths: []string{"."},
	}
	for _, opt := range opts {
		opt(v)
	}
	return v
}

func (v *ViperLoader) Load(cfg interface{}) (err error) {
	rv := reflect.ValueOf(cfg)
	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		return ErrInvalidInput
	}
	if v.consulURL != "" {
		err = v.loadFromConsul()
		if err == nil {
			err = v.Unmarshal(cfg)
			return
		}
	}
	if v.vaultURL != "" && v.vaultToken != "" {
		dataCfg, err := v.loadFromVault()
		if err != nil {
			return err
		}
		b, err := json.Marshal(dataCfg)
		if err != nil {
			return err
		}
		err = json.Unmarshal(b, cfg)
		if err != nil {
			return err
		}

		return nil
	}

	log.Printf("Can not load from consule, either consul url is not set, or an error occured: %+v. Will load configuration from file and environment variables.\n", err)
	err = v.loadFromFileAndEnv()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		err = fmt.Errorf("%w: no '%s' file found on search paths.", ErrConfigFileNotFound, v.configFileName)
		return
	}
	err = v.Unmarshal(cfg)
	return
}

func (v *ViperLoader) loadFromFileAndEnv() error {
	v.SetConfigName(v.configFileName)
	for _, path := range v.configFileSearchPaths {
		v.AddConfigPath(path)
	}
	v.SetEnvPrefix(v.envPrefix)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()
	return v.ReadInConfig()
}

func (v *ViperLoader) loadFromConsul() error {
	err := v.AddRemoteProvider("consul", v.consulURL, v.consulKey)
	if err != nil {
		return err
	}

	v.SetConfigType("yaml")
	stop := false
	attempt := 0
	for !stop {
		err = v.ReadRemoteConfig()
		attempt++
		stop = err == nil || attempt >= v.remoteMaxAttempt
		if !stop {
			time.Sleep(500 * time.Millisecond)
		}
	}

	log.Printf("Initializing remote config, consul endpoint: %s, consul key: %s, number of attempt: %d", v.consulURL, v.consulKey, attempt)
	return err
}

func (v *ViperLoader) loadFromVault() (data interface{}, err error) {
	var (
		attempt = 0
		config  = vault.DefaultConfig()
	)

	config.Address = v.vaultURL
	client, err := vault.NewClient(config)
	if err != nil {
		err = fmt.Errorf("unable to initialize Vault client: %v", err)
		return
	}
	// Authenticate
	client.SetToken(v.vaultToken)

	// Read a secret from the default mount path for KV v2 in dev mode, "secret"
	secret, err := client.KVv2("appsecret").Get(context.Background(), v.vaultKey)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	data = secret.Data["data"].(map[string]interface{})

	log.Printf("Initializing remote config, vault endpoint: %s, vault token: %s, number of attempt: %d", v.vaultURL, v.vaultToken, attempt)
	return
}
