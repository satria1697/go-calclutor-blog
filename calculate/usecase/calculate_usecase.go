package usecase

import (
	"dumpro/calculate/domain"
	"github.com/gin-gonic/gin"
	"strconv"
)

type calculateUseCase struct {
	calculateRepo domain.CalculateRepository
}

func NewCalculateUseCase(repo domain.CalculateRepository) domain.CalculateUseCase {
	return calculateUseCase{
		calculateRepo: repo,
	}
}

func (c calculateUseCase) GetCalculationHistoryUc(ctx *gin.Context) ([]domain.CalculationHistory, error) {
	return c.calculateRepo.GetCalculationHistoryRepository(ctx)
}

func (c calculateUseCase) GetCalculationUc(ctx *gin.Context, a string, b string) (int, int, int, float64, error) {
	first, err := strconv.ParseInt(a, 10, 32)
	second, err := strconv.ParseInt(b, 10, 32)
	if err != nil {
		return 0, 0, 0, 0, err
	}
	sum, sub, times, div, err := c.calculateRepo.GetCalculationRepository(ctx, int(first), int(second))
	if err != nil {
		return 0, 0, 0, 0, err
	}
	return sum, sub, times, div, nil
}
