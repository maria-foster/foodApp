package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

func init() {
	var Config Config 

    yamlFile, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, Config)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

}