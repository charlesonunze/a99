package handler

import (
	"context"

	"github.com/charlesonunze/a99/pb"
)

type server struct{}

// New - returns an instance of the CarServiceServer
func New() pb.CarServiceServer {
	return &server{}
}

func (s *server) RegisterCar(ctx context.Context, req *pb.RegisterCarRequest) (*pb.CarResponse, error) {
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
