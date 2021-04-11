package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	//	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/lib/pq"
	_ "github.com/lib/pq" //using pq
)

//SetupModels set up database connect
func SetupModels() *gorm.DB {

	//Enable Viper to read environmental varibles
	viperUser := "postgres"  // viper.Get("POSTGRES_USER")
	viperPassword := "1234"  // viper.Get("POSTGRES_PASSWORD")
	viperDB := "thealth"     //viper.Get("POSTGRES_DB")
	viperHost := "localhost" //viper.Get("POSTGRES_HOST")
	viperPort := "5432"      //viper.Get("POSTGRES_PORT")

	//setting up connection string
	postgresCon := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		viperHost,
		viperPort,
		viperUser,
		viperDB,
		viperPassword,
	)

	db, err := gorm.Open("postgres", postgresCon)
	if err != nil {
		panic("Failed to connect to database :" + err.Error())
	}

	db.AutoMigrate(&PatientProfiles{})

	//Initialize with values
	m := PatientProfiles{
		//	ID:              1,
		FullName:        "Tese",
		Gender:          "Verinumbe",
		Contact:         "090569844",
		ReligionCulture: "Christain",
		Telephone:       "0804484034",
		Languages:       pq.StringArray{"English", "French", "German", "Mandarin"},
		Assessment:      "Excellent",
		History:         pq.StringArray{"none"},
		Objective:       pq.StringArray{"Excellence"},
		AttachedFiles:   nil,
		ResponseTime:    time.Now(),
	}

	db.Create(&m)

	return db
}
