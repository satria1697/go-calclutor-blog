package main

import (
	"dumpro/database"
	redisClient "dumpro/database/redis"
	_ "dumpro/docs"
	"dumpro/router"
	"dumpro/utils"
	"github.com/sirupsen/logrus"
)

// @title           Calculator API
// @version         1.0
// @description     Calculate your data and get history with it
// @termsOfService  http://swagger.io/terms/

func main() {
	config := utils.GetConfigs()

	postgrestDb, err := database.InitDatabase(config.PostgresConfig)
	if err != nil {
		logrus.Fatalf("Error Database %v", err)
	}
	redisDb := redisClient.InitRedis(config.RedisConfig)

	r := router.SetupRouter(postgrestDb, redisDb)
	err = r.Run(config.Port)
	if err != nil {
		logrus.Fatalf("Error Gin %v", err)
	}
}
