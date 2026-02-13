package cache
import (
	"fmt"
	"time"
	"github.com/Stevesadr/golang-backend-project/config"
	"github.com/go-redis/redis"
)
var redisClient *redis.Client
func InitRedis(cfg *config.Config)  {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port),
		Password: cfg.Redis.Password,
		DB: 0,
		ReadTimeout: cfg.Redis.ReadTimeout * time.Second,
		DialTimeout: cfg.Redis.DialTimeout *time.Second,
		WriteTimeout: cfg.Redis.WriteTimeout *time.Second,
		PoolSize: cfg.Redis.PoolSize,
		PoolTimeout: cfg.Redis.PoolTimeOut,
		IdleTimeout: cfg.Redis.IdleTimeout * time.Microsecond,
		IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Microsecond,
	})
}

func GetRedis() *redis.Client{
	return redisClient
}

func CloseRedis() {
	redisClient.Close()
}