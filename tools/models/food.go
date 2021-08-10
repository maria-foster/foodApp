package models

type Food struct {
	FDCID			int					`json:'fdcId'`
	Description		string				`json:'description'`
	DataType		string				`json:'dataType'`
	PublicationDate	string				`json:'publicationDate'`
	FoodCode		string				`json:'foodCode'`
	FoodNutrients	[]FoodNutrition		`json:'foodNutrients'`
}

type FoodNutrition struct {
	Number		string 		`json:'number'`
	Name		string 		`json:'name'`
	Amount		float64		`json:'amount'`
	UnitName	string		`json:'unitName'`
}