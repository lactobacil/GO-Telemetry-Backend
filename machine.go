package entity

import "time"

type Machine struct {
	TransactionID    int       `json:"history_id" gorm:"primaryKey"`
	Transaction_Time time.Time `json:"transaction_time" gorm:"not null"`
	Asset_Name       string    `json:"Asset_Name" gorm:"not null"`
	Asset_Image      string    `json:"asset_image_path" gorm:"not null"`
}
