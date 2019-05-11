package config

import (
			"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

const FILE_NAME = "config.json"

type DataBase struct {
	Name string `json:"name"`
	User string `json:"user"`
	Password string `json:"password"`
}


type Config struct {
	DataBase *DataBase `json:"database"`
	//Name string `json:"name"`
	//User string `json:"user"`
	//Password string `json:"password"`
}



var config = new(Config)
func init() {
	fmt.Println("init config...")
	content, err := ioutil.ReadFile(FILE_NAME)
	if err != nil {
		log.Fatal("can't find config.json")
	}
	if err := json.Unmarshal(content, config); err == nil {
		fmt.Println("json -> struct")
		fmt.Println(config)
	} else {
		fmt.Println(err)
	}
}

func Get() *Config {
	return config
}
