package USDA

import (
	// data "foodApp/tools/config"
	models "foodApp/tools/models"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
)


func GetAllFoods() []models.Food { 
	var food []models.Food
	resp, err := http.Get("https://api.nal.usda.gov/fdc/v1/foods/list?api_key=dhNl9zeD9ugZmrR2Ru0neMEe8qWsPKCSwWfxJ49p")
	if err != nil {
	   log.Fatalln(err)
	}
 //We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}

	err = json.Unmarshal(body, &food)
	if err != nil {
		log.Fatalln(err)
	 }

	return food
	
}