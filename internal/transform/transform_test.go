package transform

import (
	"testing"
	"time"

	"github.com/charlesonunze/a99/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestMapToFeaturesArray(t *testing.T) {
	assert := assert.New(t)
	data := []string{"sunroof", "panorama", "auto-parking"}
	featuresArr := MapToFeaturesArray(data)

	assert.Equal(len(featuresArr), 3)
	assert.Equal(featuresArr[0].Name, "sunroof")
	assert.Equal(featuresArr[1].Name, "panorama")
	assert.Equal(featuresArr[2].Name, "auto-parking")
}

func TestFeaturesMapToStringArray(t *testing.T) {
	assert := assert.New(t)
	features := []model.Feature{
		{Name: "sunroof"},
		{Name: "panorama"},
		{Name: "auto-parking"},
	}
	stringArr := MapFeaturesToStringArray(features)

	assert.Equal(len(stringArr), 3)
	assert.Equal(stringArr[0], "sunroof")
	assert.Equal(stringArr[1], "panorama")
	assert.Equal(stringArr[2], "auto-parking")
}

func TestMapCarsToResponseArray(t *testing.T) {
	assert := assert.New(t)
	cars := []model.Car{
		{
			ID:          "4a2fdc1e-8bea-44d1-a381-ba91548387dg",
			CarType:     "Van",
			Name:        "Mercedes benz X463",
			Color:       "red",
			CreateTime:  time.Now(),
			LastUpdated: time.Now(),
			SpeedRange:  1,
			Features: []model.Feature{
				{
					ID:    "4a2fdc1e-8bea-44d1-a381-ba91548387dz",
					CarID: "4a2fdc1e-8bea-44d1-a381-ba91548387dg",
					Name:  "sunroof",
				},
			},
		},
		{
			ID:          "4a2fdc1e-8bea-44d1-a381-ba91548387da",
			CarType:     "Van",
			Name:        "Mercedes benz X462",
			Color:       "red",
			CreateTime:  time.Now(),
			LastUpdated: time.Now(),
			SpeedRange:  12,
			Features: []model.Feature{
				{
					ID:    "4a2fdc1e-8bea-44d1-a381-ba91548387dw",
					CarID: "4a2fdc1e-8bea-44d1-a381-ba91548387da",
					Name:  "sunroof",
				},
			},
		},
	}

	response := MapCarsToResponseArray(cars)

	assert.Equal(len(response), 2)
	assert.Equal(response[0].Type, "Van")
	assert.Equal(response[0].Color, "red")
	assert.Equal(response[0].Name, "Mercedes benz X463")
	assert.Equal(len(response[0].Features), 1)

	assert.Equal(response[1].Type, "Van")
	assert.Equal(response[1].Color, "red")
	assert.Equal(response[1].Name, "Mercedes benz X462")
	assert.Equal(len(response[0].Features), 1)
}
