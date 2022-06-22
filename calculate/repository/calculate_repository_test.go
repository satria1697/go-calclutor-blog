package repository_test

//
//import (
//	"dumpro/calculate/repository"
//	"github.com/stretchr/testify/assert"
//	"testing"
//)
//
//func TestDivide(t *testing.T) {
//	a := 10
//	b := 5
//
//	rep := repository.NewCalculateRepository()
//	t.Run("whole number", func(t *testing.T) {
//		res := rep.Divide(a, b)
//		assert.Equal(t, float64(2), res)
//	})
//
//	t.Run("float number", func(t *testing.T) {
//		res := rep.Divide(b, a)
//		assert.Equal(t, 0.5, res)
//	})
//	t.Run("divide by 0", func(t *testing.T) {
//		a = 0
//		b = 0
//		res := rep.Divide(a, b)
//		assert.Equal(t, float64(0), res)
//	})
//}
//
//func TestSubtract(t *testing.T) {
//	rep := repository.NewCalculateRepository()
//	t.Run("whole number", func(t *testing.T) {
//		t.Run("positive positive", func(t *testing.T) {
//			a := 10
//			b := 5
//			t.Run("ab", func(t *testing.T) {
//				res := rep.Subtract(a, b)
//				assert.Equal(t, 5, res)
//			})
//			t.Run("ba", func(t *testing.T) {
//				res := rep.Subtract(b, a)
//				assert.Equal(t, -5, res)
//			})
//		})
//		t.Run("positive negative", func(t *testing.T) {
//			a := 10
//			b := -5
//			t.Run("ab", func(t *testing.T) {
//				res := rep.Subtract(a, b)
//				assert.Equal(t, 15, res)
//			})
//			t.Run("ba", func(t *testing.T) {
//				res := rep.Subtract(b, a)
//				assert.Equal(t, -15, res)
//			})
//		})
//		t.Run("negative positive", func(t *testing.T) {
//			a := -10
//			b := 5
//			t.Run("ab", func(t *testing.T) {
//				res := rep.Subtract(a, b)
//				assert.Equal(t, -15, res)
//			})
//			t.Run("ba", func(t *testing.T) {
//				res := rep.Subtract(b, a)
//				assert.Equal(t, 15, res)
//			})
//		})
//		t.Run("negative negative", func(t *testing.T) {
//			a := -10
//			b := -5
//			t.Run("ab", func(t *testing.T) {
//				res := rep.Subtract(a, b)
//				assert.Equal(t, -5, res)
//			})
//			t.Run("ba", func(t *testing.T) {
//				res := rep.Subtract(b, a)
//				assert.Equal(t, 5, res)
//			})
//		})
//	})
//}
//
//func TestSum(t *testing.T) {
//	rep := repository.NewCalculateRepository()
//	t.Run("whole number", func(t *testing.T) {
//		t.Run("positive number", func(t *testing.T) {
//			a := 10
//			b := 5
//			res := rep.Sum(a, b)
//			assert.Equal(t, 15, res)
//		})
//		t.Run("negative number", func(t *testing.T) {
//			a := 10
//			b := -5
//			res := rep.Sum(b, a)
//			assert.Equal(t, 5, res)
//		})
//	})
//}
//
//func TestTimes(t *testing.T) {
//	rep := repository.NewCalculateRepository()
//	t.Run("whole number", func(t *testing.T) {
//		t.Run("positive positive", func(t *testing.T) {
//			a := 10
//			b := 5
//			res := rep.Times(a, b)
//			assert.Equal(t, 50, res)
//		})
//		t.Run("positive negative", func(t *testing.T) {
//			a := 10
//			b := -5
//			res := rep.Times(b, a)
//			assert.Equal(t, -50, res)
//		})
//		t.Run("negative positive", func(t *testing.T) {
//			a := -10
//			b := 5
//			res := rep.Times(b, a)
//			assert.Equal(t, -50, res)
//		})
//		t.Run("negative negative", func(t *testing.T) {
//			a := -10
//			b := -5
//			res := rep.Times(b, a)
//			assert.Equal(t, 50, res)
//		})
//	})
//}
