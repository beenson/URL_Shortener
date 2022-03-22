package cache

import (
	"context"
	"fmt"
	"os"

	util "github.com/beenson/URL_Shortener/pkg/utils"
	"github.com/go-redis/redis/v8"
)

var (
	Instance *redis.Client
	Ctx      = context.Background()
)

func Init() {
	var (
		address  = os.Getenv("CACHE_ADDRESS")
		port     = os.Getenv("CACHE_PORT")
		password = os.Getenv("CACHE_PASSWORD")
		db       int
		poolsize int
	)

	util.ConvertEnvToInt(&db, "CACHE_DB", 0)
	util.ConvertEnvToInt(&poolsize, "CACHE_POOLSIZE", 10)

	Instance = redis.NewClient(&redis.Options{
		Addr:     address + ":" + port,
		Password: password,
		DB:       db,
		PoolSize: poolsize,
	})

	if _, err := Instance.Ping(Ctx).Result(); err != nil {
		fmt.Println("🔴 Connection to Redis failed:", err)
		return
	}

	fmt.Println("🟢 Redis Connection Init Success.")
}
