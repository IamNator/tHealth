package models

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	//	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/lib/pq"
	_ "github.com/lib/pq" //using pq
)

//SetupModels set up database connect
func SetupModels() *gorm.DB {

	//Enable Viper to read environmental varibles
	viper.AutomaticEnv()

	viperUser := viper.Get("POSTGRES_USER")
	viperPassword := viper.Get("POSTGRES_PASSWORD")
	viperDB := viper.Get("POSTGRES_DB")
	viperHost := viper.Get("POSTGRES_HOST")
	viperPort := viper.Get("POSTGRES_PORT")

	//setting up connection string
	postgresCon := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		viperHost,
		viperPort,
		viperUser,
		viperDB,
		viperPassword,
	)

	fmt.Println(postgresCon)
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
