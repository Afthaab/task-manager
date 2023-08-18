package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", os.Getenv("DBHOST"), os.Getenv("DBUSER"), os.Getenv("DBNAME"), os.Getenv("DBPORT"), os.Getenv("DBPASSWORD"))
	DB, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	return DB, dbErr

}
