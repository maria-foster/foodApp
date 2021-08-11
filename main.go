package main

import (
	USDA "foodApp/USDA"
	parser "foodApp/tools/parser"

	"log"
)


func main(){ 
	USDA.GetAllFoods()
	profiles := parser.ReadProfileJSON("./data/profiles.json")
	
	for _, p := range profiles {
		p.CalculateBMR()
		p.CalculateMacros()
	}
}