package models

type Food struct {
	Number		string 		`json:'number'`
	Name		string 		`json:'name'`
	Amount		int			`json:'amount'`
	UnitName	string		`json:'unitName'`
}
