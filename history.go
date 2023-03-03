package entity

type History struct {
	HistoryId int    `json:"history_id" gorm:"primaryKey"`
	Country   string `json:"country" gorm:"not null"`
	Value     int    `json:"value" gorm:"not null"`
}
