package models

import (
	"time"
)

type Receipt struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	CustomerUid string    `json:"customer_uid"`
	CreatedAt   time.Time `json:"created_at"`
}
