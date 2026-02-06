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

func Read() (Config, error) {
	filePath, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}
	home := filePath + "/.gatorconfig.json"

	file, err := os.Open(home)
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
