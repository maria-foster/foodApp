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

}

func (profile Profile) CalculateBMR(){
	if strings.EqualFold("female", profile.Sex) {
		profile.BMR =  655.0 + (4.35 * profile.Weight) + (4.7 * profile.HeightInches) + (4.7 * profile.Age)
	}

}

func (profile Profile) CalculateActiveCalories() {
	
}