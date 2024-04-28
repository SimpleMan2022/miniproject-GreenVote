package entities

import "time"

type User struct {
	Id        int
	Email     string
	Fullname  string
	Password  string
	AddressId *Address
	Image     *string
	CreatedAt time.Time
	UpdatedAt time.Time
}
