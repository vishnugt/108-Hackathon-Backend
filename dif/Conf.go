package dif

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	API_KEY         string
	FCM_KEY         string
	Host            string
	Port            string
	Username        string
	Password        string
	DBname          string
	Vehicle_FCM_Key string
	User_FCM_Key    string
}

func ReadConf() Configuration {
	var Conf Configuration
	file, _ := os.Open("/Users/karsinkk/src/Go/src/github.com/karsinkk/108-Hackathon-Backend/dif/conf.json")
	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&Conf)
	return Conf
}
