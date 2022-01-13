package transform

import "github.com/charlesonunze/a99/internal/model"

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
