package main

import (
	USDA "foodApp/USDA"
	parser "foodApp/tools/parser"
	// MacroCalculator "foodApp/macroCalculator"
	// "log"
)


func main(){ 
	USDA.GetAllFoods()
	profiles := parser.ReadProfileJSON("./data/profiles.json")
	
	for _, p := range profiles {
		log.Print(p.CalculateBMR())
	}
}