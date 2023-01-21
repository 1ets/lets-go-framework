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
func (tbl *user) SetAdapter(db *gorm.DB) {
	tbl.db = db
}

// Get all user data.
// Use on controller.
func (tbl *user) Get() (result []*models.User, err error) {
	response := tbl.db.Find(&result)
	lets.LogI("Get.RowsAffected: %v", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		lets.LogE("Example.Get(): ", err)
		return
	}

	return
}
