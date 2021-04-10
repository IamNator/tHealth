package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // using postgres sql
	"github.com/spf13/viper"
)

//SetupModels set up database connect
func SetupModels() *gorm.DB {

	//Enable Viper to read environmental varibles
	viperUser := viper.Get("POSTGRES_USER")
	viperPassword := viper.Get("POSTGRES_PASSWORD")
	viperDB := viper.Get("POSTGRES_DB")
	viperHost := viper.Get("POSTGRES_HOST")
	viperPort := viper.Get("POSTGRES_PORT")

	//setting up connection string
	postgresConnName := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		viperHost,
		viperPort,
		viperUser,
		viperDB,
		viperPassword,
	)

	db, err := gorm.Open("postgres", postgresConnName)
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
