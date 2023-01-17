package repository

import (
	"fmt"
	"lets-go-framework/app/models"

	"gorm.io/gorm"
)

var Account account = account{}

type account struct {
	adapter *gorm.DB // Database adapter
}

func (account *account) SetAdapter(adapter *gorm.DB) {
	account.adapter = adapter
}

// Retrieve all data
func (account *account) Get() (result []*models.Account, err error) {
	db := account.adapter

	response := db.Find(&result)
	fmt.Printf("Get.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		fmt.Println(err)
		return
	}

	return
}

// Search by id and return as model
func (account *account) Find(id uint) (result *models.Account, err error) {
	fmt.Println("repository.account.Find()")

	db := account.adapter
	model := models.Account{}

	response := db.First(&model, id)
	fmt.Printf("Find.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {

		fmt.Println(err)
		return
	}

	result = &model
	return
}

// Insert a record
func (account *account) Insert(row models.Account) (result *models.Account, err error) {
	db := account.adapter

	response := db.Create(&row)
	fmt.Printf("Insert.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		fmt.Println(err)
		return
	}

	result = &row
	return
}

// Update record by id
func (account *account) UpdateById(id uint, row models.Account) (result *models.Account, err error) {
	db := account.adapter

	response := db.Updates(&row)
	fmt.Printf("UpdateById.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		fmt.Println(err)
		return
	}

	result = &row

	return
}

// Delete record by id
func (account *account) DeleteById(id uint) (err error) {
	db := account.adapter

	response := db.Delete(&models.Account{}, id)
	fmt.Printf("DeleteById.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		fmt.Println(err)
		return
	}

	return
}
