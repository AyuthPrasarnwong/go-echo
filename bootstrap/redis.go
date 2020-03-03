package bootstrap

import (
	"fmt"

	"github.com/go-redis/redis/v7"
	"github.com/spf13/viper"
)

type (
	// RedisDB database management
	RedisDB struct {
	}
)

// dbsRedis variable for define connection
var dbsRedis = make(map[string]*redis.Client)

// CreateRedisConnection make connection
func CreateRedisConnection() {
	// gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	// 	return "jhi_" + defaultTableName
	// }
	for k, v := range viper.Get("redis").(map[string]interface{}) {
		x := v.(map[string]interface{})
		port := 6379
		if val, found := x["port"]; found {
			port = val.(int)
		}

		host := x["host"].(string)
		password := x["password"].(string)
		database := x["db"].(int)

		db := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", host, port),
			Password: password, // no password set
			DB:       database, // use default DB
		})

		if _, err := db.Ping().Result(); err != nil {
			panic(fmt.Sprintf("failed to connect redis of %s connection", k))
		}

		dbsRedis[k] = db
	}
}

// Redis get mysql connection
func (ctl *RedisDB) Redis(x interface{}) *redis.Client {
	if x == nil {
		return dbsRedis["default"]
	}
	if connection, found := dbsRedis[x.(string)]; found {
		return connection
	}
	panic(fmt.Sprintf("connection %s not found", x.(string)))
}
