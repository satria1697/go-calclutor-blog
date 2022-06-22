package delivery

import (
	"dumpro/calculate/domain"
)

func MapCalculateResponse(sum int, sub int, times int, div float64) domain.CalculateResponse {
	return domain.CalculateResponse{
		Sum:    sum,
		Sub:    sub,
		Times:  times,
		Divide: div,
	}
}

func MapCalculateHistoryResponse(res []domain.CalculationHistory) []domain.CalculationHistoryResponse {
	var ponse []domain.CalculationHistoryResponse
	for _, re := range res {
		var pon domain.CalculationHistoryResponse
		pon.ID = re.ID
		pon.FirstInteger = re.FirstInteger
		pon.SecondInteger = re.SecondInteger
		pon.Sum = re.Sum
		pon.Subtract = re.Subtract
		pon.Times = re.Times
		pon.Divide = re.Divide
		ponse = append(ponse, pon)
	}
	if len(res) == 0 {
		return []domain.CalculationHistoryResponse{}
	}
	return ponse
}
