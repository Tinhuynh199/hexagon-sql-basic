package app

import (
	"hexrestapi/pkg/client"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig        `yaml: "server"`
	Sql    SqlConfig           `yaml: "sql"`
	Client client.ClientConfig `yaml: "client"`
}

type ServerConfig struct {
	Name string `yaml: "name"`
	Port *int64 `yaml: "port"`
}

type SqlConfig struct {
	Driver     string `yaml: "driver"`
	Data_Source string `yaml: "data_source"`
}

func (c *Config) getConf() *Config {

	yamlFile, err := ioutil.ReadFile("./configs/config.yml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func GetConfig() *Config {
	var conf Config
	conf.getConf()

	server := conf.Server
	sql := conf.Sql
	client := conf.Client

	return &Config{
		Server: server,
		Sql:    sql,
		Client: client,
	}
}
