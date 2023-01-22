package repository

import (
	"lets-go-framework/app/models"
	"lets-go-framework/lets"

	"gorm.io/gorm"
)

// Define repository.
var User = &user{}

// Repository user.
type user struct {
	db *gorm.DB
}

// Implement types.IMySQLRepository.
// Mandatory.
func (tbl *user) SetDriver(db *gorm.DB) {
	tbl.db = db
}

// Get all user data.
// Use on controller.
func (tbl *user) Get() (result []*models.User, err error) {
	response := tbl.db.Find(&result)
	lets.LogI("Get.RowsAffected: %v", response.RowsAffected)

	if err = response.Error; err != nil {
		lets.LogE("Example.Get(): ", err)
		return
	}

	return
}

// Insert into user table.
func (tbl *user) Insert(data *models.User) (result *models.User, err error) {
	response := tbl.db.Create(&data)
	lets.LogI("Insert.RowsAffected: %v", response.RowsAffected)

	if err = response.Error; err != nil {
		lets.LogE("Example.Insert(): ", err.Error())
		return
	}

	result = data
	return
}
