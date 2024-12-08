package config

import "time"

type (
	ConfigConsulOptions struct {
		consulKey string
		consulURL string
	}

	ConfigVaultOptions struct {
		vaultURL        string
		vaultToken      string
		vaultSecretPath string
		vaultMountPath  string
	}

	ConfigLocalOptions struct {
		fileName        string
		fileSearchPaths []string
		tagName         string
	}
)

type (
	ConfigOptions struct {
		ConfigConsulOptions
		ConfigVaultOptions
		ConfigLocalOptions
		remoteMaxAttempt      int
		env                   string
		provider              Provider
		isAutoRefresh         bool
		remoteRefreshInterval time.Duration
	}
	ConfigOption func(*Config)
)

func WithRemoteRefreshInterval(remoteRefreshInterval time.Duration) ConfigOption {
	return func(v *Config) {
		v.options.remoteRefreshInterval = remoteRefreshInterval
	}
}
func WithRemoteIsAutoRefresh(isAutoRefresh bool) ConfigOption {
	return func(v *Config) {
		v.options.isAutoRefresh = isAutoRefresh
	}
}

func WithConfigFileSearchPaths(paths ...string) ConfigOption {
	return func(v *Config) {
		if len(paths) > 0 {
			v.options.fileSearchPaths = append(v.options.fileSearchPaths, paths...)
		}
	}
}

func WithLoadFromConsulMaxAttempt(n int) ConfigOption {
	return func(v *Config) {
		if n > 0 {
			v.options.remoteMaxAttempt = n
		}
	}
}

func WithConfigFileName(fileName string) ConfigOption {
	return func(v *Config) {
		v.options.fileName = fileName
	}
}

func WithStructTagName(name string) ConfigOption {
	return func(v *Config) {
		v.options.tagName = name
	}
}

func WithConsul(url, key string) ConfigOption {
	return func(v *Config) {
		v.options.consulURL = url
		v.options.consulKey = key
	}
}

func WithVault(url, token, mountPath, secretPath string) ConfigOption {
	return func(v *Config) {
		v.options.vaultURL = url
		v.options.vaultToken = token
		v.options.vaultMountPath = mountPath
		v.options.vaultSecretPath = secretPath
	}
}
