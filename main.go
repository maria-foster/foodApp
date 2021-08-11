package main

import (
	USDA "foodApp/USDA"
	MacroCalculator "foodApp/macroCalculator"
	// "log"
)


func main(){ 
	USDA.GetAllFoods()
	MacroCalculator.CollectProfileInfo()
	// for i, food:= range allFoods {
	// 	log.Println(i, food)
	// }
}