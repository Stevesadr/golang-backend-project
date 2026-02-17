package main

import (
	"github.com/Stevesadr/golang-backend-project/api"
	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/Stevesadr/golang-backend-project/data/cache"
	"github.com/Stevesadr/golang-backend-project/data/db"
	"github.com/Stevesadr/golang-backend-project/pkg/logging"
)

// @contact.name Steve  Sadr
// @contact.url https://github.com/stevesadr
// @contact.email steve.sadr@gmail.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main(){
	cfg := config.GetConfig()
	logger := logging.NewLogger(cfg)
	err := cache.InitRedis(cfg)
	if err != nil {
		logger.Fatal(logging.Redis, logging.Startup, err.Error(), nil)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil{
		logger.Fatal(logging.Postgres, logging.Startup, err.Error(), nil)
	}
	defer db.CloseDb()
	
	api.InitServer(cfg)  
}