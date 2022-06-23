package repository

import (
	"dumpro/calculate/domain"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
)

const CalculateHistoryKey = "C-History"

type calculateRepository struct {
	postgresDb *gorm.DB
	redisDb    *redis.Client
}

func NewCalculateRepository(postgresDb *gorm.DB, redisDb *redis.Client) domain.CalculateRepository {
	return calculateRepository{
		postgresDb: postgresDb,
		redisDb:    redisDb,
	}
}

func (r calculateRepository) GetCalculationHistoryRepository(ctx *gin.Context) ([]domain.CalculationHistory, error) {
	var res []domain.CalculationHistory

	result, err := r.redisDb.Get(ctx, CalculateHistoryKey).Result()
	if err == redis.Nil {
		fmt.Printf("%v\n", result)

		dbRes := r.postgresDb.Unscoped().Find(&res)
		if dbRes.Error != nil {
			return nil, dbRes.Error
		}
		marshal, err := json.Marshal(res)
		if err != nil {
			return nil, err
		}
		err = r.redisDb.Set(ctx, CalculateHistoryKey, marshal, 100*time.Second).Err()
		if err != nil {
			return nil, err
		}
		return res, nil
	} else if err != nil {
		return nil, err
	} else {
		err := json.Unmarshal([]byte(result), &res)
		if err != nil {
			fmt.Printf("%v\n", err)
			return nil, err
		}
		return res, nil
	}
}

func (r calculateRepository) GetCalculationRepository(ctx *gin.Context, a int, b int) (int, int, int, float64, error) {
	sum := r.Sum(a, b)
	sub := r.Subtract(a, b)
	times := r.Times(a, b)
	div := r.Divide(a, b)

	calculationHistory := domain.CalculationHistory{
		FirstInteger:  a,
		SecondInteger: b,
		Sum:           sum,
		Subtract:      sub,
		Times:         times,
		Divide:        div,
	}

	resDb := r.postgresDb.Create(&calculationHistory)
	if resDb.Error != nil {
		return 0, 0, 0, 0, resDb.Error
	}

	return sum, sub, times, div, nil
}

func (r calculateRepository) Sum(a int, b int) int {
	return a + b
}

func (r calculateRepository) Subtract(a int, b int) int {
	return a - b
}

func (r calculateRepository) Divide(a int, b int) float64 {
	if a == 0 || b == 0 {
		return float64(0)
	}
	return float64(a) / float64(b)
}

func (r calculateRepository) Times(a int, b int) int {
	return a * b
}
