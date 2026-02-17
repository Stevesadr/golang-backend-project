package db

import (
	"fmt"
	"time"

	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/Stevesadr/golang-backend-project/pkg/logging"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var logger = logging.NewLogger(config.GetConfig())

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
	logger.Info(logging.Postgres, logging.Startup, "Db connection successful", nil)
	return nil
}

func GetDb() *gorm.DB{
	return dbClient
}

func CloseDb(){
	con, _ := dbClient.DB()
	con.Close()
}