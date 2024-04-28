package entities

import "github.com/google/uuid"

type Address struct {
	Id          uuid.UUID
	Province    string
	City        string
	SubDistrict string
	StreetName  string
	ZipCode     string
	CreatedAt   string
	UpdatedAt   string
}
