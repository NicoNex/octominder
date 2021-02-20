/*
 * Octominder - a simple reminder.
 * Copyright (C) 2021  Nicolò Santamaria
 *
 * Octominder is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * Octominder is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <https://www.gnu.org/licenses/>.
 */

package main

import "github.com/BurntSushi/toml"

type Reminder struct {
	Msg  string `toml:"message"`
	Tick string `toml:"repeat"`
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
