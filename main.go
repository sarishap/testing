package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

var m = "var/log/config/firewall"
var j = "syslog"

type LogFile struct {
	LogDirectories map[string]string `json:"log-directories"`
}

func main() {
	body, err := os.ReadFile("config.json")
	ft := LogFile{LogDirectories: make(map[string]string)}
	err = json.Unmarshal(body, &ft)

	if err != nil {
		log.Println("error")
		return
	}
	value, ok := ft.LogDirectories[m]
	if ok {
		v := strings.Split(m, "/")
		if v[len(v)-1] == "firewall" {
			if j != value {
				fmt.Println("Filetype doesn't match")

			} else {
				fmt.Println("Filetype matches")
			}
		}
	}
}
