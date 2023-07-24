package config

import (
	"auth_audit/config/configReader"
)

type Config struct {
	configReader configReader.ConfigReader
}

func NewConfig(configReader configReader.ConfigReader) *Config {
	return &Config{configReader: configReader}
}

func (c *Config) LoadConfig(pathToConfigFile, configFileType, configFileName string) error {
	c.configReader.SetConfigName(configFileName)
	c.configReader.SetConfigType(configFileType)
	c.configReader.AddConfigPath(pathToConfigFile)

	if err := c.configReader.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) GetValueByKeys(key string) string {
	value := c.configReader.GetString(key)
	if value == "" {
		panic("value with such key was not found or is empty, which is impermissible")
	}

	return value
}
