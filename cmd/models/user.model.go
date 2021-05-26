package models

import "time"

type User struct {
	ID        uint
	Name      string
	Phone     string
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
}
