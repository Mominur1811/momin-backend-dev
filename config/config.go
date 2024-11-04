package config

import (
	"encoding/json"
	"flag"
	"github.com/go-playground/validator"
	"log/slog"
	"os"
)

type Config struct {
	Mode        string `json:"mode"`
	ServiceName string `json:"service_name"`
	HttpPort    int    `json:"http_port"`
	Db          DB     `json:"db"`
}

type DB struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
	User     string `json:"user"`
	Password string `json:"password"`
	SslMode  string `json:"ssl_mode"`
}

var config *Config
var confFlag = flag.String("c", "/home/mominur/go prog/LibraryManagement/config.json", "Configuration file path")

func init() {
	flag.Parse()
	var data []byte
	var err error

	exit := func(err error) {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if data, err = os.ReadFile(*confFlag); err != nil {
		exit(err)
	}

	if err = json.Unmarshal(data, &config); err != nil {
		exit(err)
	}

	v := validator.New()
	if err = v.Struct(config); err != nil {
		exit(err)
	}

}
