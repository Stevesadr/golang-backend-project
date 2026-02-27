package migration

import (
	"fmt"

	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/Stevesadr/golang-backend-project/data/db"
	"github.com/Stevesadr/golang-backend-project/data/models"
	"github.com/Stevesadr/golang-backend-project/pkg/logging"
)

var logger =  logging.NewLogger(config.GetConfig())

func Up_1(){
	database := db.GetDb()

	tables := []interface{}{}

	country := models.Country{}
	fmt.Println("Has country table:", database.Migrator().HasTable(country))
	if !database.Migrator().HasTable(country){
		tables = append(tables, country)
	}

	city := models.City{}
	fmt.Println("Has city table:", database.Migrator().HasTable(city))
	if !database.Migrator().HasTable(city){
		tables = append(tables, city)
	}

	fmt.Println("Tables to create:", len(tables))
	err := database.Migrator().CreateTable(tables...)
	if err != nil{
		logger.Error(logging.Postgres, logging.Migration, "", nil)
	}
	
	logger.Info(logging.Postgres, logging.Migration, "tables created", nil)
}