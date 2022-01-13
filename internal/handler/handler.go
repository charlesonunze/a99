package handler

import (
	"context"

	"github.com/charlesonunze/a99/internal/model"
	"github.com/charlesonunze/a99/internal/repo"
	"github.com/charlesonunze/a99/internal/service"
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

	var features []model.Feature

	for _, f := range req.Features {
		features = append(features, model.Feature{
			Name: f,
		})
	}

	svc := s.GetService()
	car, err := svc.RegisterCar(ctx, model.Car{
		CarType:    req.Type,
		Name:       req.Name,
		Color:      req.Color,
		SpeedRange: req.SpeedRange,
		Features:   features,
	})
	if err != nil {
		return nil, err
	}

	var stringArr []string

	for _, f := range car.Features {
		stringArr = append(stringArr, f.Name)
	}

	return &pb.CarResponse{
		Type:       car.CarType,
		Name:       car.Name,
		Color:      car.Color,
		SpeedRange: car.SpeedRange,
		Features:   stringArr,
	}, nil
}

func (s *server) GetCarByID(ctx context.Context, req *pb.GetCarRequest) (*pb.CarResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return &pb.CarResponse{
		Type:       "Sedan",
		Name:       "Lucid Air",
		Color:      "blue",
		SpeedRange: 1,
		Features:   []string{"better than tesla"},
	}, nil
}

func (s *server) ListCars(ctx context.Context, req *pb.ListCarsRequest) (*pb.CarsResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	return &pb.CarsResponse{
		Cars: []*pb.CarResponse{
			{
				Type:       "Sedan",
				Name:       "Lucid Air",
				Color:      "blue",
				SpeedRange: 1,
				Features:   []string{"better than tesla"},
			},
		},
	}, nil
}
