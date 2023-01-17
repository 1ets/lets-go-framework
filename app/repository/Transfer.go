package repository

import (
	"fmt"
	"lets-go-framework/app/models"

	"gorm.io/gorm"
)

var Transfer transfer = transfer{}

type transfer struct {
	adapter *gorm.DB // Database adapter
}

func (transfer *transfer) SetAdapter(adapter *gorm.DB) {
	transfer.adapter = adapter
}

// Retrieve all data
func (transfer *transfer) Get() (result []*models.Transaction, err error) {
	db := transfer.adapter

	response := db.Find(&result)
	fmt.Printf("Get.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		fmt.Println(err)
		return
	}

	return
}

// Search by id and return as model
func (transfer *transfer) Find(id uint) (result *models.Transaction, err error) {
	fmt.Println("repository.transfer.Find()")

	db := transfer.adapter
	model := models.Transaction{}

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
func (transfer *transfer) Insert(row models.Transaction) (result *models.Transaction, err error) {
	fmt.Println("Insert(row models.Transaction)")

	db := transfer.adapter
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
func (transfer *transfer) UpdateById(id uint, row models.Transaction) (result *models.Transaction, err error) {
	db := transfer.adapter

	response := db.Where("id = ?", id).Updates(&row)
	fmt.Printf("UpdateById.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		fmt.Println(err)
		return
	}

	result = &row

	return
}

// Delete record by id
func (transfer *transfer) DeleteById(id uint) (err error) {
	db := transfer.adapter

	response := db.Delete(&models.Transaction{}, id)
	fmt.Printf("DeleteById.RowsAffected: %v\n", response.RowsAffected)

	if err = response.Error; response.Error != nil {
		fmt.Println(err)
		return
	}

	return
}
