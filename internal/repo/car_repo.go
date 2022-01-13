package repo

import (
	"github.com/charlesonunze/a99/internal/model"
	"gorm.io/gorm"
)

// CarRepo - car repository
type CarRepo interface {
	InsertOne(car model.Car) (model.Car, error)
	FindByID(carID string) (model.Car, error)
	Find(car model.Car) ([]model.Car, error)
}

type carRepo struct {
	db *gorm.DB
}

// New - returns a new car repo
func New(db *gorm.DB) CarRepo {
	return &carRepo{db}
}

func (r *carRepo) InsertOne(car model.Car) (model.Car, error) {
	if err := r.db.Create(&car).Error; err != nil {
		return car, err
	}

	return car, nil
}

func (r *carRepo) FindByID(carID string) (model.Car, error) {
	var car model.Car

	if err := r.db.Preload("Features").Find(&car, "id = ?", carID).Error; err != nil {
		return car, err
	}

	return car, nil
}

func (r *carRepo) Find(car model.Car) ([]model.Car, error) {
	var cars []model.Car
	if err := r.db.Preload("Features").Find(&cars, car).Error; err != nil {
		return cars, err
	}

	return cars, nil
}
