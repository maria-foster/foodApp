package parser

import (
	models "foodApp/tools/models"
	"encoding/json"
	"log"
	"io/ioutil"
    "os"
)

func ReadProfileJSON(fileName string) ([]models.Profile){
	jsonFile, err := os.Open(fileName)
	// if we os.Open returns an error then handle it
	if err != nil {
        log.Println(err)
    }
    log.Println("Successfully Opened:" + fileName)
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
        log.Println(err)
    }
	
    var result []models.Profile

    json.Unmarshal([]byte(byteValue), &result)


  	return result
}