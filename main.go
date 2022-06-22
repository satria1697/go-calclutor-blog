package main

import (
	calculatehandler "dumpro/calculate/delivery/http"
	calculaterepository "dumpro/calculate/repository"
	calculateusecase "dumpro/calculate/usecase"
	"dumpro/database/postgresql"
	"dumpro/database/redis"
	_ "dumpro/docs"
	"dumpro/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Calculator API
// @version         1.0
// @description     Calculate your data and get history with it
// @termsOfService  http://swagger.io/terms/

func main() {
	config := utils.GetConfigs()

	postgrestDb, err := postgresql.InitDatabase(config.PostgresConfig)
	if err != nil {
		logrus.Fatalf("Error Database %v", err)
	}
	redisDb := redis.InitRedis(config.RedisConfig)

	r := gin.Default()
	calculateRepo := calculaterepository.NewCalculateRepository(postgrestDb, redisDb)
	calculateUseCase := calculateusecase.NewCalculateUseCase(calculateRepo)
	calculatehandler.NewCalculateHandler(r, calculateUseCase)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	err = r.Run(config.Port)
	if err != nil {
		logrus.Fatalf("Error Gin %v", err)
	}
}
