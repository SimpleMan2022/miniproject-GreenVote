package helpers

import "evoting/dto"

type ResponseWithData struct {
	Status     bool
	StatusCode int
	Message    string
	Data       any
}

type ResponseWithoutData struct {
	Status     bool
	StatusCode int
	Message    string
}

func Response(param dto.ResponseParams) any {
	var status bool
	var response any
	if param.StatusCode >= 200 && param.StatusCode < 300 {
		status = true
	} else {
		status = false
	}

	if param.Data != nil {
		response = ResponseWithData{
			Status:     status,
			StatusCode: param.StatusCode,
			Message:    param.Message,
			Data:       param.Data,
		}
	} else {
		response = ResponseWithoutData{
			Status:     status,
			StatusCode: param.StatusCode,
			Message:    param.Message,
		}
	}
	return response

}
