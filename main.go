package main

import (
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
	cfg, err := readConfig(cfgpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	logpath := filepath.Join(os.Getenv("HOME"), ".log", "otcominder.log")
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
