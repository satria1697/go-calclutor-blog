// Code generated by mockery v2.13.1. DO NOT EDIT.

package domain

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// MockCalculateRepository is an autogenerated mock type for the CalculateRepository type
type MockCalculateRepository struct {
	mock.Mock
}

// Divide provides a mock function with given fields: a, b
func (_m *MockCalculateRepository) Divide(a int, b int) float64 {
	ret := _m.Called(a, b)

	var r0 float64
	if rf, ok := ret.Get(0).(func(int, int) float64); ok {
		r0 = rf(a, b)
	} else {
		r0 = ret.Get(0).(float64)
	}

	return r0
}

// GetCalculationHistoryRepository provides a mock function with given fields: ctx
func (_m *MockCalculateRepository) GetCalculationHistoryRepository(ctx *gin.Context) ([]CalculationHistory, error) {
	ret := _m.Called(ctx)

	var r0 []CalculationHistory
	if rf, ok := ret.Get(0).(func(*gin.Context) []CalculationHistory); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]CalculationHistory)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*gin.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCalculationRepository provides a mock function with given fields: ctx, a, b
func (_m *MockCalculateRepository) GetCalculationRepository(ctx *gin.Context, a int, b int) (int, int, int, float64, error) {
	ret := _m.Called(ctx, a, b)

	var r0 int
	if rf, ok := ret.Get(0).(func(*gin.Context, int, int) int); ok {
		r0 = rf(ctx, a, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(*gin.Context, int, int) int); ok {
		r1 = rf(ctx, a, b)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 int
	if rf, ok := ret.Get(2).(func(*gin.Context, int, int) int); ok {
		r2 = rf(ctx, a, b)
	} else {
		r2 = ret.Get(2).(int)
	}

	var r3 float64
	if rf, ok := ret.Get(3).(func(*gin.Context, int, int) float64); ok {
		r3 = rf(ctx, a, b)
	} else {
		r3 = ret.Get(3).(float64)
	}

	var r4 error
	if rf, ok := ret.Get(4).(func(*gin.Context, int, int) error); ok {
		r4 = rf(ctx, a, b)
	} else {
		r4 = ret.Error(4)
	}

	return r0, r1, r2, r3, r4
}

// Subtract provides a mock function with given fields: a, b
func (_m *MockCalculateRepository) Subtract(a int, b int) int {
	ret := _m.Called(a, b)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(a, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Sum provides a mock function with given fields: a, b
func (_m *MockCalculateRepository) Sum(a int, b int) int {
	ret := _m.Called(a, b)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(a, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Times provides a mock function with given fields: a, b
func (_m *MockCalculateRepository) Times(a int, b int) int {
	ret := _m.Called(a, b)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, int) int); ok {
		r0 = rf(a, b)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

type mockConstructorTestingTNewMockCalculateRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockCalculateRepository creates a new instance of MockCalculateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockCalculateRepository(t mockConstructorTestingTNewMockCalculateRepository) *MockCalculateRepository {
	mock := &MockCalculateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
