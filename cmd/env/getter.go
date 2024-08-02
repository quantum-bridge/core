package env

import (
	"github.com/pkg/errors"
	bridgeErrors "github.com/quantum-bridge/core/pkg/errors"
	"github.com/spf13/viper"
	"os"
	"sync"
)

// envConfigFile is the environment variable that holds the path to the config file
const envConfigFile = "ENV_CONFIG_FILE"

// Getter is the interface that implements getting a string map from a key.
type Getter interface {
	// GetStringMap gets a string map from the given key.
	GetStringMap(key string) (map[string]interface{}, error)
}

// viperBackend is the struct that holds the viperInstance backend to get the configuration.
type viperBackend struct {
	viperInstance *viper.Viper
	read          sync.Once
	readErr       error
}

// getter is the struct that holds viper backends for the Getter interface.
type getter struct {
	backends []Getter
}

// NewViperFile creates a new viper backend with the given file name.
func NewViperFile(fn string) Getter {
	// Create a new viper instance with the given file name.
	viperInstance := viper.New()
	viperInstance.SetConfigFile(fn)

	return &viperBackend{
		viperInstance: viperInstance,
	}
}

// MustFromEnv returns a Getter from the environment variables.
func MustFromEnv() Getter {
	// Get the getter from the environment variables.
	getter, err := FromEnv()
	if err != nil {
		panic(errors.Wrap(err, "kv.FromEnv panicked"))
	}

	return getter
}

// ensureRead ensures that the config file is read.
func (v *viperBackend) ensureRead() error {
	// Read the config file only once.
	v.read.Do(func() {
		// Read the config file.
		if err := v.viperInstance.ReadInConfig(); err != nil {
			v.readErr = errors.Wrapf(err, "failed to read config file %s", v.viperInstance.ConfigFileUsed())
		}
	})

	// Return the read error.
	return v.readErr
}

// GetStringMap gets a string map from the given key.
func (v *viperBackend) GetStringMap(key string) (map[string]interface{}, error) {
	// Ensure that the config file is read.
	if err := v.ensureRead(); err != nil {
		return nil, err
	}

	// Get the string map from the given key.
	return v.viperInstance.GetStringMap(key), nil
}

// FromEnv returns a Getter from the environment variables.
func FromEnv() (Getter, error) {
	var getters []Getter
	if viperFn := os.Getenv(envConfigFile); viperFn != "" {
		getter := NewViperFile(viperFn)
		_, err := getter.GetStringMap("ping")
		if err != nil {
			return nil, errors.Wrap(err, "viper backend seems unavailable")
		}
		getters = append(getters, getter)
	}

	// If there are no backends, return an error.
	if len(getters) == 0 {
		return nil, bridgeErrors.ErrNoBackends
	}

	// Return the merged getters.
	return MergeGetters(getters...), nil
}

// MergeGetters merges the given backends into a single Getter.
func MergeGetters(backends ...Getter) Getter {
	return &getter{
		backends: backends,
	}
}

// GetStringMap gets a string map from the given key.
func (g getter) GetStringMap(key string) (map[string]interface{}, error) {
	// Get the string map from the given key from the backends.
	for _, backend := range g.backends {
		// Get the string map from the given key.
		value, err := backend.GetStringMap(key)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get key %s", key)
		}

		// If the value is not nil, return it.
		if value != nil {
			return value, nil
		}
	}

	return nil, nil
}

// GetterFunc is a function that implements the Getter interface.
type GetterFunc func(key string) (map[string]interface{}, error)

// GetStringMap gets a string map from the given key.
func (f GetterFunc) GetStringMap(key string) (map[string]interface{}, error) {
	return f(key)
}
