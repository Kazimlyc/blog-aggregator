package config

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {

	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	dat, err := io.ReadAll(file)
	if err != nil {
		return Config{}, err
	}

	var configResp Config
	if err := json.Unmarshal(dat, &configResp); err != nil {
		return Config{}, err
	}

	return configResp, nil
}

func getConfigFilePath() (string, error) {
	filePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	home := filePath + "/" + configFileName

	return home, nil
}

func (c *Config) SetUser(username string) error {

	c.CurrentUserName = username
	return write(*c)

}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	file, err := getConfigFilePath()
	if err != nil {
		return err
	}

	if err := os.WriteFile(file, data, 0644); err != nil {
		return err
	}
	return nil
}
