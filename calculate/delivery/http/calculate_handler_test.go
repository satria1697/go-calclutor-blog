package http_test

import (
	calculatehandler "dumpro/calculate/delivery/http"
	"dumpro/calculate/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCalculateHandler_GetCalculation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := errors.New("error")

	t.Run("Get Data", func(t *testing.T) {
		_, engine := gin.CreateTestContext(httptest.NewRecorder())
		mockCalculateUseCase := domain.NewMockCalculateUseCase(t)
		mockCalculateUseCase.On("GetCalculationUc", mock.Anything, "10", "10").Return(10, 10, 10, float64(10), nil)
		calculatehandler.NewCalculateHandler(engine, mockCalculateUseCase)
		req, err := http.NewRequest(http.MethodGet, "/calculate?first=10&second=10", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
	t.Run("error parse", func(t *testing.T) {
		_, engine := gin.CreateTestContext(httptest.NewRecorder())
		mockCalculateUseCase := domain.NewMockCalculateUseCase(t)
		mockCalculateUseCase.On("GetCalculationUc", mock.Anything, "10", "asd").Return(0, 0, 0, float64(0), err)
		calculatehandler.NewCalculateHandler(engine, mockCalculateUseCase)
		req, err := http.NewRequest(http.MethodGet, "/calculate?first=10&second=asd", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})

	t.Run("error Calculate", func(t *testing.T) {
		_, engine := gin.CreateTestContext(httptest.NewRecorder())
		mockCalculateUseCase := domain.NewMockCalculateUseCase(t)
		mockCalculateUseCase.On("GetCalculationUc", mock.Anything, "10", "10").Return(0, 0, 0, float64(0), err)
		calculatehandler.NewCalculateHandler(engine, mockCalculateUseCase)
		req, err := http.NewRequest(http.MethodGet, "/calculate?first=10&second=10", nil)
		assert.NoError(t, err)
		w := httptest.NewRecorder()
		engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}
