package models

type Employee struct {
	ID       int     `json:"id" gorm:"primary_key"`
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}
