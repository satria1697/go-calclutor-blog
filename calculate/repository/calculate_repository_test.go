package repository_test

import (
	"database/sql/driver"
	"dumpro/calculate/domain"
	"dumpro/calculate/repository"
	"encoding/json"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http/httptest"
	"testing"
	"time"
)

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

var histories = []domain.CalculationHistory{
	{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		FirstInteger:  10,
		SecondInteger: 10,
		Sum:           10,
		Subtract:      10,
		Times:         10,
		Divide:        10,
	},
	{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		FirstInteger:  20,
		SecondInteger: 20,
		Sum:           20,
		Subtract:      20,
		Times:         20,
		Divide:        20,
	},
}

func initDb(t *testing.T) (*gorm.DB, *redis.Client, sqlmock.Sqlmock, redismock.ClientMock, *sqlmock.Rows) {
	db, sMock, err := sqlmock.New()
	assert.NoError(t, err)
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	client, rMock := redismock.NewClientMock()

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at", "deleted_at", "first_integer", "second_integer", "sum", "subtract", "times", "divide"}).
		AddRow(histories[0].ID, histories[0].CreatedAt, histories[0].UpdatedAt, histories[0].DeletedAt, histories[0].FirstInteger, histories[0].SecondInteger, histories[0].Sum, histories[0].Subtract, histories[0].Times, histories[0].Divide).
		AddRow(histories[1].ID, histories[1].CreatedAt, histories[1].UpdatedAt, histories[1].DeletedAt, histories[1].FirstInteger, histories[1].SecondInteger, histories[1].Sum, histories[1].Subtract, histories[1].Times, histories[1].Divide)

	return gormDb, client, sMock, rMock, rows
}

func TestCalculateRepository_GetCalculationRepository(t *testing.T) {
	gin.SetMode(gin.TestMode)

	context, _ := gin.CreateTestContext(httptest.NewRecorder())

	t.Run("Error", func(t *testing.T) {
		db, client, sMock, _, _ := initDb(t)

		rep := repository.NewCalculateRepository(db, client)
		err := errors.New("error")
		sMock.ExpectBegin()
		sMock.ExpectQuery("INSERT INTO (.*?)").WillReturnError(err)
		sMock.ExpectRollback()

		sum, sub, times, div, err := rep.GetCalculationRepository(context, 10, 10)
		assert.Error(t, err)
		assert.Zero(t, sum)
		assert.Zero(t, sub)
		assert.Zero(t, times)
		assert.Zero(t, div)
	})

	t.Run("Success", func(t *testing.T) {
		db, client, sMock, _, _ := initDb(t)

		rep := repository.NewCalculateRepository(db, client)

		sMock.ExpectBegin()
		sMock.ExpectQuery("INSERT INTO (.*?) VALUES (.*?) RETURNING (.*?)").
			WithArgs(AnyTime{}, AnyTime{}, nil, 10, 20, 30, -10, 200, 0.5).
			WillReturnRows(sqlmock.NewRows([]string{"ID"}).AddRow(uint(2)))
		sMock.ExpectCommit()
		sum, sub, times, div, err := rep.GetCalculationRepository(context, 10, 20)
		assert.NoError(t, err)
		assert.NotZero(t, sum)
		assert.NotZero(t, sub)
		assert.NotZero(t, times)
		assert.NotZero(t, div)
	})

}

func TestCalculateRepository_GetCalculationHistoryRepository(t *testing.T) {
	gin.SetMode(gin.TestMode)

	context, _ := gin.CreateTestContext(httptest.NewRecorder())
	t.Run("redis nil", func(t *testing.T) {
		t.Run("error db", func(t *testing.T) {
			err := errors.New("error")

			db, client, sMock, rMock, _ := initDb(t)

			rep := repository.NewCalculateRepository(db, client)
			defer rMock.ClearExpect()
			rMock.ExpectGet(repository.CalculateHistoryKey).RedisNil()
			sMock.ExpectQuery("SELECT (.*?)").WillReturnError(err)

			res, err := rep.GetCalculationHistoryRepository(context)

			assert.Error(t, err)
			assert.Nil(t, res)
		})
		t.Run("error set on redis", func(t *testing.T) {

			db, client, sMock, rMock, rows := initDb(t)

			rep := repository.NewCalculateRepository(db, client)
			defer rMock.ClearExpect()
			rMock.ExpectGet(repository.CalculateHistoryKey).RedisNil()
			sMock.ExpectQuery("SELECT (.*?)").
				WillReturnRows(rows)
			marshal, err := json.Marshal(histories)
			assert.NoError(t, err)

			err = errors.New("error")

			rMock.ExpectSet(repository.CalculateHistoryKey, marshal, 100*time.Second).SetErr(err)
			res, err := rep.GetCalculationHistoryRepository(context)
			assert.Error(t, err)
			assert.Nil(t, res)
		})
		t.Run("success", func(t *testing.T) {

			db, client, sMock, rMock, rows := initDb(t)

			rep := repository.NewCalculateRepository(db, client)
			defer rMock.ClearExpect()
			rMock.ExpectGet(repository.CalculateHistoryKey).RedisNil()
			sMock.ExpectQuery("SELECT (.*?)").
				WillReturnRows(rows)
			marshal, err := json.Marshal(histories)
			assert.NoError(t, err)

			rMock.ExpectSet(repository.CalculateHistoryKey, marshal, 100*time.Second).SetVal("success")
			res, err := rep.GetCalculationHistoryRepository(context)
			assert.NoError(t, err)
			assert.Len(t, res, 2)
		})
	})
}

func TestCalculateRepository_Divide(t *testing.T) {
	a := 10
	b := 5

	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	client, _ := redismock.NewClientMock()

	rep := repository.NewCalculateRepository(gormDb, client)
	t.Run("whole number", func(t *testing.T) {
		res := rep.Divide(a, b)
		assert.Equal(t, float64(2), res)
	})

	t.Run("float number", func(t *testing.T) {
		res := rep.Divide(b, a)
		assert.Equal(t, 0.5, res)
	})
	t.Run("divide by 0", func(t *testing.T) {
		a = 0
		b = 0
		res := rep.Divide(a, b)
		assert.Equal(t, float64(0), res)
	})
}

func TestCalculateRepository_Subtract(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	client, _ := redismock.NewClientMock()

	rep := repository.NewCalculateRepository(gormDb, client)
	t.Run("whole number", func(t *testing.T) {
		t.Run("positive positive", func(t *testing.T) {
			a := 10
			b := 5
			t.Run("ab", func(t *testing.T) {
				res := rep.Subtract(a, b)
				assert.Equal(t, 5, res)
			})
			t.Run("ba", func(t *testing.T) {
				res := rep.Subtract(b, a)
				assert.Equal(t, -5, res)
			})
		})
		t.Run("positive negative", func(t *testing.T) {
			a := 10
			b := -5
			t.Run("ab", func(t *testing.T) {
				res := rep.Subtract(a, b)
				assert.Equal(t, 15, res)
			})
			t.Run("ba", func(t *testing.T) {
				res := rep.Subtract(b, a)
				assert.Equal(t, -15, res)
			})
		})
		t.Run("negative positive", func(t *testing.T) {
			a := -10
			b := 5
			t.Run("ab", func(t *testing.T) {
				res := rep.Subtract(a, b)
				assert.Equal(t, -15, res)
			})
			t.Run("ba", func(t *testing.T) {
				res := rep.Subtract(b, a)
				assert.Equal(t, 15, res)
			})
		})
		t.Run("negative negative", func(t *testing.T) {
			a := -10
			b := -5
			t.Run("ab", func(t *testing.T) {
				res := rep.Subtract(a, b)
				assert.Equal(t, -5, res)
			})
			t.Run("ba", func(t *testing.T) {
				res := rep.Subtract(b, a)
				assert.Equal(t, 5, res)
			})
		})
	})
}

func TestCalculateRepository_Sum(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	client, _ := redismock.NewClientMock()

	rep := repository.NewCalculateRepository(gormDb, client)
	t.Run("whole number", func(t *testing.T) {
		t.Run("positive number", func(t *testing.T) {
			a := 10
			b := 5
			res := rep.Sum(a, b)
			assert.Equal(t, 15, res)
		})
		t.Run("negative number", func(t *testing.T) {
			a := 10
			b := -5
			res := rep.Sum(b, a)
			assert.Equal(t, 5, res)
		})
	})
}

func TestCalculateRepository_Times(t *testing.T) {
	db, _, err := sqlmock.New()
	assert.NoError(t, err)
	gormDb, err := gorm.Open(postgres.New(postgres.Config{Conn: db}))
	assert.NoError(t, err)

	client, _ := redismock.NewClientMock()

	rep := repository.NewCalculateRepository(gormDb, client)
	t.Run("whole number", func(t *testing.T) {
		t.Run("positive positive", func(t *testing.T) {
			a := 10
			b := 5
			res := rep.Times(a, b)
			assert.Equal(t, 50, res)
		})
		t.Run("positive negative", func(t *testing.T) {
			a := 10
			b := -5
			res := rep.Times(b, a)
			assert.Equal(t, -50, res)
		})
		t.Run("negative positive", func(t *testing.T) {
			a := -10
			b := 5
			res := rep.Times(b, a)
			assert.Equal(t, -50, res)
		})
		t.Run("negative negative", func(t *testing.T) {
			a := -10
			b := -5
			res := rep.Times(b, a)
			assert.Equal(t, 50, res)
		})
	})
}
