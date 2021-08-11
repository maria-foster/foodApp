package models

import (
	"strings"
)
type Profile struct {
	Sex 			string
	Age				float64
	Weight			float64
	HeightInches	float64
	ActiveCalories	float64
	Goal			string
	BMR				float64
	Macros			Macros

}

func (profile Profile) CalculateBMR(){
	if strings.EqualFold("female", profile.Sex) {
		profile.BMR =  655.0 + (4.35 * profile.Weight) + (4.7 * profile.HeightInches) + (4.7 * profile.Age)
	}

}

func (profile Profile) CalculateActiveCalories() {
	
}

func (profile Profile) CalculateMacros() {
	if(strings.EqualFold(profile.Goal, "maintenance" )){
		profile.Macros.Fat.FatPercentage = .3
		profile.Macros.Carbohydrates.CarbohydratePercentage = .3
		profile.Macros.Protein.ProteinPercentage = .3
	}

	profile.Macros.Fat.FatCalories = profile.BMR * profile.Macros.Fat.FatPercentage
	profile.Macros.Fat.FatGrams =  profile.BMR * profile.Macros.Fat.FatPercentage / 9

	profile.Macros.Carbohydrates.CarbohydrateCalories = profile.BMR * profile.Macros.Carbohydrates.CarbohydratePercentage
	profile.Macros.Carbohydrates.CarbohydrateGrams =  profile.BMR * profile.Macros.Carbohydrates.CarbohydratePercentage / 4

	profile.Macros.Protein.ProteinCalories = profile.BMR * profile.Macros.Protein.ProteinPercentage
	profile.Macros.Protein.ProteinGrams =  profile.BMR * profile.Macros.Protein.ProteinPercentage / 4
}