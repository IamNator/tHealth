package models

import (
	"time"

	"gorm.io/gorm"
)

//User defines a model for user information
type User struct {
	ID    int `json:"id"`
	Price int
	Code  string
}

//patientProfile is a model for patient_profile_table
type patientProfile struct {
	ID              string    `json:"id"`
	FullName        string    `json:"fullname"`
	Gender          string    `json:"gender"`
	Contact         string    `json:"contact"`
	ReligionCulture string    `json:"religion_culture"`
	Telephone  	     string    `json:"telephone"`
	Languages       []string  `json:"languages"`
	Assessment      string    `json:"assessment"`
	History         []string  `json:"history"`
	Objective       []string  `json:"objective"`
	AttachedFiles   []string  `json:"attached_files"`
	ResponseTime    time.Time `json:"response_time"`
	TimeofCompalint time.Time `json:"time_of_compalint"`
}

//NewPatientProfile creates a new patientProfile
func NewPatientProfile()

func (p *patientProfile) Add(db *gorm.DB) {
	db.Model(&patientProfile{}).Create(p)
}

/*	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&User{})

	// Create
	db.Create(&User{Code: "D42", Price: 100})

	// Read
	var product User
	db.First(&product, 1)                 // find product with integer primary key
	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(User{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)*/
