package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Bid struct for holding auction details
type Bid struct {
	gorm.Model
	BidID     int       `json:"product_id"`
	Status    string    `json:"status"`
	StartTime time.Time `json:"start_time"`
	StopTime  time.Time `json:"stop_time"`
}
