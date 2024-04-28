package entities

import "time"

type User struct {
	Id        int
	Email     string
	Fullname  string
	Password  string
	Address   *Address
	CreatedAt time.Time
	UpdatedAt time.Time
}
