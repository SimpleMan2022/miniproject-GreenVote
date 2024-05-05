package dto

type PlaceAddress struct {
	Province    string `json:"province"`
	City        string `json:"city"`
	SubDistrict string `json:"sub_district"`
	StreetName  string `json:"street_name"`
	ZipCode     string `json:"zip_code"`
}

type PlaceAddressRequest struct {
	Province    string `json:"province"`
	City        string `json:"city"`
	SubDistrict string `json:"sub_district"`
	StreetName  string `json:"street_name"`
	ZipCode     string `json:"zip_code"`
}
