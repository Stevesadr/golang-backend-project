package db

import (
	"fmt"
	"log"
	"time"

	"github.com/Stevesadr/golang-backend-project/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func InitDb(cfg *config.Config) error {
	cns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password,
		cfg.Postgres.DbName, cfg.Postgres.SSLMode)

	dbClient, err := gorm.Open(postgres.Open(cns), &gorm.Config{})
	if err != nil{
		return err
	}
	
	sqldb , _ := dbClient.DB()
	
	err = sqldb.Ping()
	if err != nil{
		return err
	}

	sqldb.SetMaxIdleConns(cfg.Postgres.MaxIdleConns)
	sqldb.SetMaxOpenConns(cfg.Postgres.MaxOpenConns)
	sqldb.SetConnMaxLifetime(cfg.Postgres.ConnMaxLifetime * time.Minute)

	log.Println("Db connection successful")
	return nil
}

func GetDb() *gorm.DB{
	return dbClient
}

func CloseDb(){
	con, _ := dbClient.DB()
	con.Close()
}