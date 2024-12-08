package config

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
)

var (
	ErrInvalidInput       = fmt.Errorf("cfg must be a pointer to struct and initialized")
	ErrConfigFileNotFound = fmt.Errorf("config file not found")
)

var (
	DefaultRemoteRefreshInterval = time.Minute * 10
	DefaultRemoteMaxAttempt      = 5
	DefaultFileName              = "config"
	DefaultTagName               = "json"
	DefaultSearchPaths           = []string{"."}
	DefaultRemoteIsAutoRefresh   = true
)

type Provider string

var ProviderLocal Provider = "local"
var ProviderRemote Provider = "remote"

type Config struct {
	*viper.Viper
	options ConfigOptions
	*Watcher
}

func New(ctx context.Context, env string, opts ...ConfigOption) *Config {
	v := &Config{
		Viper: viper.New(),
		options: ConfigOptions{
			ConfigLocalOptions: ConfigLocalOptions{
				fileName:        DefaultFileName,
				fileSearchPaths: DefaultSearchPaths,
				tagName:         DefaultTagName,
			},
			remoteMaxAttempt:      DefaultRemoteMaxAttempt,
			env:                   env,
			remoteRefreshInterval: DefaultRemoteRefreshInterval,
			isAutoRefresh:         DefaultRemoteIsAutoRefresh,
		},
	}
	for _, opt := range opts {
		opt(v)
	}
	return v
}

func (v *Config) Load(ctx context.Context, cfg interface{}) (stopper func(ctx context.Context) error, err error) {
	var (
		remoteLoader func(ctx context.Context) (interface{}, error)
		rv           = reflect.ValueOf(cfg)
	)
	stopper = func(ctx context.Context) error { return nil }

	if rv.Kind() != reflect.Pointer || rv.IsNil() || rv.Elem().Kind() != reflect.Struct {
		err = ErrInvalidInput
		return
	}

	if v.options.consulURL != "" {
		remoteLoader = func(ctx context.Context) (data interface{}, err error) {
			err = v.loadFromConsul()
			if err == nil {
				err = v.Unmarshal(cfg)
				return
			}
			data = cfg
			return
		}
	}
	if v.options.vaultURL != "" && v.options.vaultToken != "" {
		remoteLoader = func(ctx context.Context) (data interface{}, err error) {
			dataCfg, err := v.loadFromVault()
			if err != nil {
				return
			}
			b, err := json.Marshal(dataCfg)
			if err != nil {
				return
			}
			err = json.Unmarshal(b, cfg)
			if err != nil {
				return
			}
			data = cfg
			return
		}

		return
	}

	if v.options.provider == ProviderRemote && v.options.isAutoRefresh {
		//watcher
		bg, cancel := context.WithCancel(ctx)
		v.Watcher = &Watcher{
			refreshTimer: time.NewTicker(v.options.remoteRefreshInterval),
			cancel:       cancel,
			Load:         remoteLoader,
		}

		if err = v.Watcher.load(bg, cfg); err != nil {
			cancel()
		}

		go v.Watcher.autoRefresh(bg, cfg)

		stopper = v.Watcher.Stop

		return
	} else {
		if remoteLoader != nil {
			cfg, err = remoteLoader(ctx)
			if err != nil {
				return nil, err
			}
		}
	}

	err = v.loadFromFile()
	if _, ok := err.(viper.ConfigFileNotFoundError); ok {
		err = fmt.Errorf("%w: no '%s' file found on search paths", ErrConfigFileNotFound, v.options.fileName)
		return
	}
	err = v.Unmarshal(cfg)
	return
}

func (v *Config) loadFromFile() error {
	v.options.provider = ProviderLocal
	v.SetConfigName(v.options.fileName)
	for _, path := range v.options.fileSearchPaths {
		v.AddConfigPath(path)
	}
	v.SetEnvPrefix(v.options.env)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	return v.ReadInConfig()
}

func (v *Config) loadFromConsul() (err error) {
	v.options.provider = ProviderRemote
	err = v.AddRemoteProvider("consul", v.options.consulURL, v.options.consulKey)
	if err != nil {
		return
	}

	v.SetConfigType("yaml")
	stop := false
	attempt := 0
	for !stop {
		err = v.ReadRemoteConfig()
		attempt++
		stop = err == nil || attempt >= v.options.remoteMaxAttempt
		if !stop {
			time.Sleep(500 * time.Millisecond)
		}
	}

	log.Printf("Initializing remote config, consul endpoint: %s,number of attempt: %d", v.options.consulURL, attempt)
	return
}

func (v *Config) loadFromVault() (data interface{}, err error) {
	var (
		attempt = 0
		config  = vault.DefaultConfig()
	)

	v.options.provider = ProviderRemote
	config.Address = v.options.vaultURL
	client, err := vault.NewClient(config)
	if err != nil {
		err = fmt.Errorf("unable to initialize Vault client: %v", err)
		return
	}

	// Authenticate
	client.SetToken(v.options.vaultToken)

	// Read a secret from the default mount path for KV v2 in dev mode, "secret"
	secret, err := client.KVv2("appsecret").Get(context.Background(), v.options.vaultKey)
	if err != nil {
		log.Fatalf("unable to read secret: %v", err)
	}

	data = secret.Data["data"].(map[string]interface{})

	log.Printf("Initializing remote config, vault endpoint: %s, number of attempt: %d", v.options.vaultURL, attempt)
	return
}
