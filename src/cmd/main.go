package main

import (
	"log"

	"github.com/Stevesadr/golang-backend-project/api"
	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/Stevesadr/golang-backend-project/data/cache"
	"github.com/Stevesadr/golang-backend-project/data/db"
)

func main(){
	cfg := config.GetConfig()

	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cache.CloseRedis()

	err = db.InitDb(cfg)
	if err != nil{
		log.Fatal(err)
	}
	defer db.CloseDb()
	
	api.InitServer(cfg)  
}