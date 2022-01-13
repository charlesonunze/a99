package transform

import (
	"github.com/charlesonunze/a99/internal/model"
	"github.com/charlesonunze/a99/pb"
)

// MapToFeaturesArray - maps an array of strings to []model.Feature
func MapToFeaturesArray(data []string) []model.Feature {
	var features []model.Feature

	for _, f := range data {
		features = append(features, model.Feature{
			Name: f,
		})
	}

	return features
}

// MapFeaturesToStringArray - maps an array of model.Feature to []string
func MapFeaturesToStringArray(data []model.Feature) []string {
	var features []string

	for _, f := range data {
		features = append(features, f.Name)
	}

	return features
}

// MapCarsToResponseArray - maps an array of cars to []*pb.CarResponse
func MapCarsToResponseArray(data []model.Car) []*pb.CarResponse {
	var carsRes []*pb.CarResponse

	for _, c := range data {
		carsRes = append(carsRes, &pb.CarResponse{
			Type:       c.CarType,
			Name:       c.Name,
			Color:      c.Color,
			SpeedRange: c.SpeedRange,
			Features:   MapFeaturesToStringArray(c.Features),
		})
	}

	return carsRes
}
