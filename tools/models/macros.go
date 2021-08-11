package models

type Macros struct {
	Fat 				Fat
	Protein				Protein
	Carbohydrates		Carbohydrates

}

type Fat struct {
	FatPercentage		int64
	FatGrams			int64
	FatCalories			int64

}

type Protein struct {
	ProteinPercentage		int64
	ProteinGrams			int64
	ProteinCalories			int64

}

type Carbohydrates struct {
	CarbohydratePercentage		int64
	CarbohydrateGrams			int64
	CarbohydrateCalories		int64

}