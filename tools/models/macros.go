package models

type Macros struct {
	Fat 				Fat
	Protein				Protein
	Carbohydrates		Carbohydrates

}

type Fat struct {
	FatPercentage		float64
	FatGrams			float64
	FatCalories			float64

}

type Protein struct {
	ProteinPercentage		float64
	ProteinGrams			float64
	ProteinCalories			float64

}

type Carbohydrates struct {
	CarbohydratePercentage		float64
	CarbohydrateGrams			float64
	CarbohydrateCalories		float64

}
