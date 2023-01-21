package models

type User struct {
	ID       uint   `json:"id" gorm:"column:id;type:autoIncrement;primaryKey"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}
