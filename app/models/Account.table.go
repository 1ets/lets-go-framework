package models

type Account struct {
	ID      uint    `json:"id" gorm:"column:id;type:autoIncrement;primaryKey"`
	Name    string  `json:"name" gorm:"column:name"`
	Balance float64 `json:"balance" gorm:"column:balance;default:0"`
}
