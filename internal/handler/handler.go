package handler

import (
	"context"

	"github.com/charlesonunze/a99/internal/model"
	"github.com/charlesonunze/a99/internal/repo"
	"github.com/charlesonunze/a99/internal/service"
	"github.com/charlesonunze/a99/internal/transform"
	"github.com/charlesonunze/a99/pb"
	"gorm.io/gorm"
)

type server struct {
	db *gorm.DB
}

// New - returns an instance of the CarServiceServer
func New(db *gorm.DB) pb.CarServiceServer {
	return &server{db}
}

func (s *server) GetService() service.CarService {
	carRepo := repo.New(s.db)
	return service.New(carRepo)
}

func (s *server) RegisterCar(ctx context.Context, req *pb.RegisterCarRequest) (*pb.CarResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	svc := s.GetService()
	car, err := svc.RegisterCar(ctx, model.Car{
		CarType:    req.Type,
		Name:       req.Name,
		Color:      req.Color,
		SpeedRange: req.SpeedRange,
		Features:   transform.MapToFeaturesArray(req.Features),
	})
	if err != nil {
		return nil, err
	}

	return &pb.CarResponse{
		Type:       car.CarType,
		Name:       car.Name,
		Color:      car.Color,
		SpeedRange: car.SpeedRange,
		Features:   transform.MapFeaturesToStringArray(car.Features),
	}, nil
}

func (s *server) GetCarByID(ctx context.Context, req *pb.GetCarRequest) (*pb.CarResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	carID := req.Id

	svc := s.GetService()
	car, err := svc.FetchCar(ctx, carID)
	if err != nil {
		return nil, err
	}

	return &pb.CarResponse{
		Type:       car.CarType,
		Name:       car.Name,
		Color:      car.Color,
		SpeedRange: car.SpeedRange,
		Features:   transform.MapFeaturesToStringArray(car.Features),
	}, nil
}

func (s *server) ListCars(ctx context.Context, req *pb.ListCarsRequest) (*pb.CarsResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	query := model.Car{
		CarType: req.Type,
		Color:   req.Color,
	}

	svc := s.GetService()
	cars, err := svc.FetchCars(ctx, query)
	if err != nil {
		return nil, err
	}

	return &pb.CarsResponse{
		Cars: transform.MapCarsToResponseArray(cars),
	}, nil
}
