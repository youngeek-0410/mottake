package models

import (
	"time"
)

type Receipt struct {
	ID          int `json:"id" gorm:"primaryKey"`
	Customer    Customer
	CustomerUID string    `json:"customer_uid"`
	CreatedAt   time.Time `json:"created_at"`
}
