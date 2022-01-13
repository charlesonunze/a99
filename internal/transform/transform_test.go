package transform

import (
	"testing"

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
