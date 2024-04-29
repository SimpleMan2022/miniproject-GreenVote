package dto

type ResponseParams struct {
	StatusCode int
	Message    string
	Data       any
}

type ResponseError struct {
	Status     bool
	StatusCode int
	Message    string
	Data       any
}
