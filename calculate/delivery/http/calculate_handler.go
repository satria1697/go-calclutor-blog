package http

import (
	"dumpro/calculate/delivery"
	"dumpro/calculate/domain"
	"dumpro/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CalculateHandler struct {
	calculateUseCase domain.CalculateUseCase
}

func NewCalculateHandler(r *gin.Engine, calculateUc domain.CalculateUseCase) {
	handler := &CalculateHandler{
		calculateUseCase: calculateUc,
	}
	r.GET("/calculate", handler.GetCalculation)
	r.GET("/calculate/history", handler.GetCalculationHistory)
}

// GetCalculation godoc
// @Summary      Calculate between 2 integer
// @ID           get-calculate
// @Tags         Calculate
// @Produce      json
// @Param        first   query      string  true "string"
// @Param        second   query      string  true "string"
// @Success      200  {object}  delivery.CalculateResponse
// @Router       /calculate [get]
func (h CalculateHandler) GetCalculation(c *gin.Context) {
	var request domain.CalculateGetRequest
	err := c.ShouldBindQuery(&request)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	sum, sub, times, div, err := h.calculateUseCase.GetCalculationUc(c, request.First, request.Second)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Data:  delivery.MapCalculateResponse(sum, sub, times, div),
		Error: nil,
	})
}

// GetCalculationHistory godoc
// @Summary      Get Calculate history
// @ID           get-calculate-history
// @Tags         Calculate
// @Produce      json
// @Success      200  {object}  delivery.CalculationHistoryResponse
// @Router       /calculate/history [get]
func (h CalculateHandler) GetCalculationHistory(c *gin.Context) {
	res, err := h.calculateUseCase.GetCalculationHistoryUc(c)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.Response{Error: err})
		return
	}
	c.JSON(http.StatusOK, utils.Response{
		Data:  delivery.MapCalculateHistoryResponse(res),
		Error: nil,
	})
}
