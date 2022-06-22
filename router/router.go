package router

import (
	calculatehandler "dumpro/calculate/delivery/http"
	calculaterepository "dumpro/calculate/repository"
	calculateusecase "dumpro/calculate/usecase"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupRouter(postgrestDb *gorm.DB, redisDb *redis.Client) *gin.Engine {
	r := gin.Default()
	calculateRepo := calculaterepository.NewCalculateRepository(postgrestDb, redisDb)
	calculateUseCase := calculateusecase.NewCalculateUseCase(calculateRepo)
	calculatehandler.NewCalculateHandler(r, calculateUseCase)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
