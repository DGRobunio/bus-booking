package util

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var (
	Redis *redis.Client
	DB    *sql.DB
)

func Init() {
	RedisConn()
	DBConn()
}

func RedisConn() {
	Redis = redis.NewClient(&redisConf)
	pong, err := Redis.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf(pong)
	}
}

func DBConn() {
	var err error
	DB, err = sql.Open("mysql", dbConf)
	if err != nil {
		log.Fatalln(err)
	} else {
		err := DB.Ping()
		if err != nil {
			log.Fatalln(err)
		}
	}
}

func Report(err error) {
	if err != nil {
		log.Println(err)
	}
}

func BadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest})
}
