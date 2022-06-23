package domain

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CalculateGetRequest struct {
	First  string `form:"first" binding:"required"`
	Second string `form:"second" binding:"required"`
}

type CalculationHistory struct {
	gorm.Model
	FirstInteger  int
	SecondInteger int
	Sum           int
	Subtract      int
	Times         int
	Divide        float64
}

type CalculateUseCase interface {
	GetCalculationUc(ctx *gin.Context, a string, b string) (int, int, int, float64, error)
	GetCalculationHistoryUc(ctx *gin.Context) ([]CalculationHistory, error)
}

type CalculateRepository interface {
	GetCalculationRepository(ctx *gin.Context, a int, b int) (int, int, int, float64, error)
	Sum(a int, b int) int
	Subtract(a int, b int) int
	Times(a int, b int) int
	Divide(a int, b int) float64
	GetCalculationHistoryRepository(ctx *gin.Context) ([]CalculationHistory, error)
}
