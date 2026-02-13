package main

import (
	"github.com/Stevesadr/golang-backend-project/api"
	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/Stevesadr/golang-backend-project/data/cache"
)

func main(){
	cfg := config.GetConfig()
	cache.InitRedis(cfg)
	defer cache.CloseRedis()
	api.InitServer(cfg)  
}