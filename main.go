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

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/TheCreeper/go-notify"
)

var (
	logger  *log.Logger
	cfgpath = filepath.Join(os.Getenv("HOME"), ".config", "octominder.toml")
)

func tick(summary, body string, tch <-chan time.Time) {
	var ntf = notify.NewNotification(summary, body)

	for _ = range tch {
		if _, err := ntf.Show(); err != nil {
			logger.Println("tick", summary, err)
		}
	}
}

func main() {
	flag.StringVar(&cfgpath, "c", cfgpath, "Config file.")
	flag.Usage = usage
	flag.Parse()

	cfg, err := readConfig(cfgpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	logpath := filepath.Join(os.Getenv("HOME"), ".log", "octominder.log")
	logfile, err := os.OpenFile(logpath, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	logger = log.New(logfile, "", log.LstdFlags)

	for k, rem := range cfg.Reminders {
		go tick(k, rem.Msg, time.Tick(time.Minute*time.Duration(rem.Tick)))
		logger.Println(fmt.Sprintf("%q reminder enabled..", k))
	}
	select {}
}

func usage() {
	fmt.Printf(`octominder - A simple reminder for humans.
Octominder reminds you about stuff.
Usage:
    %s [OPTIONS]
Options:
    -c  string
        Specify a config file to use (default %s).
    -h, --help
        Prints this help message and exits.
Copyright (C) 2021 Nicolò Santamaria
`, os.Args[0], filepath.Join(os.Getenv("HOME"), ".config", "octominder.toml"))
}
