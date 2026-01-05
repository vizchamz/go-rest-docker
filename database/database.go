package database

import (
	"fmt"
	"log"
	"os"

	"github.com/vizchamz/go-rest-docker/models"
	"gorm.io/driver/postgres"
    oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	if os.Getenv("APPLICATION_DB") == "postgres" {
		dsn := fmt.Sprintf(
			"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Colombo",
			os.Getenv("POSTGRES_DB_USER"),
			os.Getenv("POSTGRES_DB_PASSWORD"),
			os.Getenv("POSTGRES_DB_NAME"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatal("Failed to connect to database. \n", err)
			os.Exit(2)
		}

		log.Println("connected")
		db.Logger = logger.Default.LogMode(logger.Info)

		log.Println("running migrations")
		db.AutoMigrate(&models.Fact{})

		DB = Dbinstance{
			Db: db,
		}
	} else if os.Getenv("APPLICATION_DB") == "oracle" {
		dsn := fmt.Sprintf(
			"oracle://%s:%s@%s:%s/%s",
			os.Getenv("ORACLE_DB_USER"),
			os.Getenv("ORACLE_DB_PASSWORD"),
			os.Getenv("ORACLE_DB_HOST"),
			os.Getenv("ORACLE_DB_PORT"),
			os.Getenv("ORACLE_DB_SID"),
		)

		db, err := gorm.Open(oracle.Open(dsn), &gorm.Config{		
			Logger: logger.Default.LogMode(logger.Info),
		})

		if err != nil {
			log.Fatal("Failed to connect to database. \n", err)
			os.Exit(2)
		}

		log.Println("connected")
		db.Logger = logger.Default.LogMode(logger.Info)

		log.Println("running migrations")
		db.AutoMigrate(&models.Fact{})

		DB = Dbinstance{
			Db: db,
		}
	}
}
