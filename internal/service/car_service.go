package service

import (
	"context"

	"github.com/charlesonunze/a99/internal/model"
	"github.com/charlesonunze/a99/internal/repo"
)

// CarService - interface for the car service
type CarService interface {
	RegisterCar(ctx context.Context, car model.Car) (model.Car, error)
	FetchCars(ctx context.Context, car model.Car) ([]model.Car, error)
	FetchCar(ctx context.Context, carID string) (model.Car, error)
}

type carService struct {
	repo repo.CarRepo
}

// New - returns an instance of the CarService
func New(repo repo.CarRepo) CarService {
	return &carService{
		repo,
	}
}

func (s *carService) RegisterCar(ctx context.Context, car model.Car) (model.Car, error) {
	car, err := s.repo.InsertOne(car)
	if err != nil {
		return car, err
	}

	return car, nil
}

func (s *carService) FetchCars(ctx context.Context, car model.Car) ([]model.Car, error) {
	cars, err := s.repo.Find(car)
	if err != nil {
		return cars, err
	}

	return cars, nil
}

func (s *carService) FetchCar(ctx context.Context, carID string) (model.Car, error) {
	car, err := s.repo.FindByID(carID)
	if err != nil {
		return car, err
	}

	return car, nil
}
