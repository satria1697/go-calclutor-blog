package usecase_test

import (
	"dumpro/calculate/domain"
	"dumpro/calculate/usecase"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCalculateUseCase_GetCalculationHistoryUc(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	err := errors.New("error")
	t.Run("Get Empty Data", func(t *testing.T) {
		mockCalculateRepo := domain.NewMockCalculateRepository(t)
		mockCalculateRepo.On("GetCalculationHistoryRepository", c).Return([]domain.CalculationHistory{}, nil).Once()
		u := usecase.NewCalculateUseCase(mockCalculateRepo)
		uc, err := u.GetCalculationHistoryUc(c)
		assert.NoError(t, err)
		assert.NotNil(t, uc)
	})

	t.Run("Error", func(t *testing.T) {
		mockCalculateRepo := domain.NewMockCalculateRepository(t)
		mockCalculateRepo.On("GetCalculationHistoryRepository", c).Return(nil, err).Once()
		u := usecase.NewCalculateUseCase(mockCalculateRepo)
		uc, err := u.GetCalculationHistoryUc(c)
		assert.Error(t, err)
		assert.Nil(t, uc)
	})

	t.Run("Get all data", func(t *testing.T) {
		data := []domain.CalculationHistory{
			{
				FirstInteger:  1,
				SecondInteger: 2,
				Sum:           3,
				Subtract:      -1,
				Times:         2,
				Divide:        0.5,
			},
			{
				FirstInteger:  10,
				SecondInteger: 20,
				Sum:           30,
				Subtract:      -10,
				Times:         200,
				Divide:        0.5,
			},
		}
		mockCalculateRepo := domain.NewMockCalculateRepository(t)
		mockCalculateRepo.On("GetCalculationHistoryRepository", c).Return(data, nil).Once()
		u := usecase.NewCalculateUseCase(mockCalculateRepo)
		uc, err := u.GetCalculationHistoryUc(c)
		assert.NoError(t, err)
		assert.NotNil(t, uc)
		assert.Len(t, uc, 2)
	})
}

func TestCalculateUseCase_GetCalculationUc(t *testing.T) {
	gin.SetMode(gin.TestMode)
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	err := errors.New("error")

	t.Run("Get Calculation", func(t *testing.T) {
		mockCalculateRepo := domain.NewMockCalculateRepository(t)
		mockCalculateRepo.On("GetCalculationRepository", c, 10, 10).Return(1, 1, 1, float64(1), nil).Once()
		u := usecase.NewCalculateUseCase(mockCalculateRepo)
		sum, sub, times, div, err := u.GetCalculationUc(c, "10", "10")
		assert.NoError(t, err)
		assert.NotZero(t, sum)
		assert.NotZero(t, sub)
		assert.NotZero(t, times)
		assert.NotZero(t, div)
	})

	t.Run("error on repository", func(t *testing.T) {
		mockCalculateRepo := domain.NewMockCalculateRepository(t)
		mockCalculateRepo.On("GetCalculationRepository", c, 10, 10).Return(0, 0, 0, float64(0), err).Once()
		u := usecase.NewCalculateUseCase(mockCalculateRepo)
		sum, sub, times, div, err := u.GetCalculationUc(c, "10", "10")
		assert.Error(t, err)
		assert.Zero(t, sum)
		assert.Zero(t, sub)
		assert.Zero(t, times)
		assert.Zero(t, div)
	})

	t.Run("error on parse", func(t *testing.T) {
		mockCalculateRepo := domain.NewMockCalculateRepository(t)
		u := usecase.NewCalculateUseCase(mockCalculateRepo)
		sum, sub, times, div, err := u.GetCalculationUc(c, "10", "asd")
		assert.Error(t, err)
		assert.Zero(t, sum)
		assert.Zero(t, sub)
		assert.Zero(t, times)
		assert.Zero(t, div)
	})
}
