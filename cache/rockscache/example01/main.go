package main

import (
	"github.com/dtm-labs/rockscache"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	rockscache.SetVerbose(true)
	// new a client for rockscache using the default options
	rc := rockscache.NewClient(redisClient, rockscache.NewDefaultOptions())

	// use Fetch to fetch data
	// 1. the first parameter is the key of the data
	// 2. the second parameter is the data expiration time
	// 3. the third parameter is the data fetch function which is called when the cache does not exist
	v, err := rc.Fetch("key1", 300 * time.Second, func()(string, error) {
		// fetch data from database or other sources
		return "value1", nil
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(v)
}
