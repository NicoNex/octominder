package main

import "github.com/BurntSushi/toml"

type Reminder struct {
	Msg  string `toml:"message"`
	Tick int    `toml:"tick"`
}

type Config struct {
	Reminders map[string]Reminder `toml:"reminder"`
}

func readConfig(path string) (Config, error) {
	var cfg Config

	if _, err := toml.DecodeFile(path, &cfg); err != nil {
		return cfg, err
	}
	return cfg, nil
}
