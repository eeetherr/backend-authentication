package utils

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strings"
)

const (
	LOCAL        = "local"
	EMPTY_STRING = ""
	YML          = "yml"
	HOST_TYPE    = "HOST_TYPE"
)

// SetupConfig sets the name for the configs file,
// takes the path name of the yaml where Viper can search for the configs file in, and
// calls LoadConfig to pull out the configs from ENV to store in Viper.
// accepts array of paths to allow for docker & test environments.
func SetupConfig(paths []string) {

	configName := os.Getenv(HOST_TYPE)

	if configName == EMPTY_STRING {
		configName = LOCAL
	}

	viper.SetConfigName(configName)

	for _, path := range paths {
		viper.AddConfigPath(path)
	}

	viper.SetConfigType(YML)

	logrus.Infof("Config file name: %s\n", configName)

	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Error reading configs file, %s\n", err)
	}

	logrus.Infof("Running in %s environment\n", configName)

	//This pulls out configs from ENV and stores in VIPER
	LoadConfig()
}

// LoadConfig This pulls the configs from environment variable and saves it inside viper.
// Make sure viper is initialized with configs from yaml file before calling
// only keys starting with $ will be considered ex. $FERNET_KEY
func LoadConfig() {

	if viper.AllSettings() == nil {
		logrus.Errorf("Viper has not been initialized with a configuration")
		return
	}

	for _, key := range viper.AllKeys() {
		value := viper.GetString(key)
		logrus.Debugf("Setting value corresponding to key: %s", key)

		if strings.HasPrefix(value, "$") {
			value = strings.TrimPrefix(value, "$")
			param := os.Getenv(value)
			viper.Set(key, param)
		} else {
			newValue := viper.Get(key)
			viper.Set(key, newValue)
		}
	}
}
