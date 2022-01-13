package model

import (
	"time"
)

type (
	// Car - the car model
	Car struct {
		ID          string    `gorm:"type:uuid;default:uuid_generate_v4()"`
		CarType     string    `gorm:"car_type"`
		Name        string    `gorm:"name"`
		Color       string    `gorm:"color"`
		SpeedRange  int32     `gorm:"speed_range"`
		CreateTime  time.Time `gorm:"create_time"`
		LastUpdated time.Time `gorm:"last_updated"`
		Features    []Feature
	}

	// Feature - features of a car
	Feature struct {
		ID    string `gorm:"type:uuid;default:uuid_generate_v4()"`
		Name  string `gorm:"name"`
		CarID string `gorm:"car_id"`
	}
)
