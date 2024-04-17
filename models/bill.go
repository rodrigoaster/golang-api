package models

type Bill struct {
	Title    string
	Products []Product
	TotalKwh float64
}
