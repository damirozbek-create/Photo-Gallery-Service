package models

type Photo struct {
	ID     uint   `gorm:"primaryKey"`
	UserID uint   `json:"user_id"`
	URL    string `json:"url"`
}
