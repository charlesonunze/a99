package repo

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/charlesonunze/a99/internal/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	carID   = "4a2fdc1e-8bea-44d1-a381-ba91548387dg"
	featID  = "4a2fdc1e-8bea-44d1-a381-ba91548387dq"
	testCar = model.Car{
		ID:          carID,
		CarType:     "Van",
		Name:        "Mercedes benz X463",
		Color:       "red",
		CreateTime:  time.Now(),
		LastUpdated: time.Now(),
		SpeedRange:  1,
		Features: []model.Feature{
			{
				ID:    featID,
				CarID: carID,
				Name:  "sunroof",
			},
		},
	}
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatal(err)
	}
	return db, mock
}

func NewGormDB(db *sql.DB) *gorm.DB {
	dialector := postgres.New(postgres.Config{
		DriverName: "postgres",
		Conn:       db,
	})
	gdb, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return gdb
}

func TestCarRepo_InsertOne(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()

	gdb := NewGormDB(db)
	carRepo := New(gdb)

	mock.ExpectBegin()
	mock.ExpectQuery(insertCar).
		WithArgs(testCar.CarType, testCar.Name, testCar.Color, testCar.SpeedRange, testCar.CreateTime, testCar.LastUpdated, testCar.ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(carID))
	mock.ExpectQuery(insertFeature).
		WithArgs(testCar.Features[0].Name, testCar.Features[0].CarID, testCar.Features[0].ID).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(featID))
	mock.ExpectCommit()

	car, err := carRepo.InsertOne(testCar)
	assert.Equal(car, testCar)
	assert.NoError(err)
}

func TestCarRepo_InsertOne_Error(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()

	gdb := NewGormDB(db)
	carRepo := New(gdb)

	mock.ExpectBegin()
	mock.ExpectQuery(insertCar).
		WithArgs(testCar.CarType, testCar.Name, testCar.Color, testCar.SpeedRange, testCar.CreateTime, testCar.LastUpdated, testCar.ID).
		WillReturnError(errors.New("something went wrong"))
	mock.ExpectQuery(insertFeature).
		WithArgs(testCar.Features[0].Name, testCar.Features[0].CarID, testCar.Features[0].ID).
		WillReturnError(errors.New("something went wrong"))
	mock.ExpectCommit()

	_, err := carRepo.InsertOne(testCar)
	assert.Error(err)
}

func TestCarRepo_FindByID(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()

	gdb := NewGormDB(db)
	carRepo := New(gdb)

	carRows := sqlmock.NewRows([]string{"car_type", "name", "color", "speed_range", "create_time", "last_updated", "id"}).
		AddRow(testCar.CarType, testCar.Name, testCar.Color, testCar.SpeedRange, testCar.CreateTime, testCar.LastUpdated, testCar.ID)
	featRows := sqlmock.NewRows([]string{"name", "car_id", "id"}).
		AddRow(testCar.Features[0].Name, testCar.Features[0].CarID, testCar.Features[0].ID)

	mock.ExpectQuery(selectFromCarsWithID).WithArgs(carID).WillReturnRows(carRows)
	mock.ExpectQuery(selectFromFeaturesWithID).WithArgs(carID).WillReturnRows(featRows)

	car, err := carRepo.FindByID(carID)
	assert.NoError(err)
	assert.Equal(car, testCar)
}

func TestCarRepo_FindByID_Error(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()

	gdb := NewGormDB(db)
	carRepo := New(gdb)

	mock.ExpectQuery(selectFromCarsWithID).WithArgs(carID).WillReturnError(errors.New("something went wrong"))
	mock.ExpectQuery(selectFromFeaturesWithID).WithArgs(carID).WillReturnError(errors.New("something went wrong"))

	car, err := carRepo.FindByID(carID)
	assert.Error(err)
	assert.NotEqual(car, testCar)
}

func TestCarRepo_Find(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()

	gdb := NewGormDB(db)
	carRepo := New(gdb)

	carRows := sqlmock.NewRows([]string{"car_type", "name", "color", "speed_range", "create_time", "last_updated", "id"}).
		AddRow(testCar.CarType, testCar.Name, testCar.Color, testCar.SpeedRange, testCar.CreateTime, testCar.LastUpdated, testCar.ID)
	featRows := sqlmock.NewRows([]string{"name", "car_id", "id"}).
		AddRow(testCar.Features[0].Name, testCar.Features[0].CarID, testCar.Features[0].ID)

	query := model.Car{
		CarType:    "Van",
		Color:      "red",
		Name:       "Mercedes benz X463",
		SpeedRange: 1,
	}
	mock.ExpectQuery(selectFromCars).WithArgs(query.CarType, query.Name, query.Color, query.SpeedRange).WillReturnRows(carRows)
	mock.ExpectQuery(selectFromFeaturesWithID).WithArgs(carID).WillReturnRows(featRows)

	cars, err := carRepo.Find(query)
	testCars := []model.Car{testCar}
	assert.NoError(err)
	assert.Equal(cars, testCars)
}

func TestCarRepo_Find_Error(t *testing.T) {
	assert := assert.New(t)
	db, mock := NewMock()
	defer db.Close()

	gdb := NewGormDB(db)
	carRepo := New(gdb)

	mock.ExpectQuery(selectFromCars).WithArgs(carID).WillReturnError(errors.New("something went wrong"))
	mock.ExpectQuery(selectFromFeaturesWithID).WithArgs(carID).WillReturnError(errors.New("something went wrong"))

	query := model.Car{
		CarType:    "Van",
		Color:      "red",
		Name:       "Mercedes benz X463",
		SpeedRange: 1,
	}

	cars, err := carRepo.Find(query)
	testCars := []model.Car{testCar}
	assert.Error(err)
	assert.NotEqual(cars, testCars)
}
