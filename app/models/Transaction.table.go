package models

type Transaction struct {
	ID        uint    `json:"id" gorm:"column:id;type:autoIncrement;primaryKey"`
	AccountId uint    `json:"account_id" gorm:"column:account_id"`
	Flow      string  `json:"flow" gorm:"column:flow"`
	Amount    float64 `json:"amount" gorm:"column:amount;default:0"`
}
