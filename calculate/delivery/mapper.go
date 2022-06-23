package delivery

import (
	"dumpro/calculate/domain"
)

type CalculationHistoryResponse struct {
	ID            uint    `json:"ID"`
	FirstInteger  int     `json:"firstInteger"`
	SecondInteger int     `json:"secondInteger"`
	Sum           int     `json:"sum"`
	Subtract      int     `json:"subtract"`
	Times         int     `json:"times"`
	Divide        float64 `json:"divide"`
}

type CalculateResponse struct {
	Sum    int     `json:"sum"`
	Sub    int     `json:"sub"`
	Times  int     `json:"times"`
	Divide float64 `json:"divide"`
}

func MapCalculateResponse(sum int, sub int, times int, div float64) CalculateResponse {
	return CalculateResponse{
		Sum:    sum,
		Sub:    sub,
		Times:  times,
		Divide: div,
	}
}

func MapCalculateHistoryResponse(res []domain.CalculationHistory) []CalculationHistoryResponse {
	var ponse []CalculationHistoryResponse
	if res == nil {
		return []CalculationHistoryResponse{}
	}
	for _, re := range res {
		var pon CalculationHistoryResponse
		pon.ID = re.ID
		pon.FirstInteger = re.FirstInteger
		pon.SecondInteger = re.SecondInteger
		pon.Sum = re.Sum
		pon.Subtract = re.Subtract
		pon.Times = re.Times
		pon.Divide = re.Divide
		ponse = append(ponse, pon)
	}
	return ponse
}
