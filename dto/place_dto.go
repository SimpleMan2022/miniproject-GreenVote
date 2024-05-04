package dto

type PlaceRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Longnitude  float64 `json:"longnitude"`
	Latitude    float64 `json:"latitude"`
}
