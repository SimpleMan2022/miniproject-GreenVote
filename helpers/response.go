package helpers

import "evoting/dto"

type ResponseWithData struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}

type ResponseWithPaginate struct {
	Status     bool     `json:"status"`
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       any      `json:"data"`
	Pagination Paginate `json:"pagination"`
	Sorting    Sort     `json:"sorting"`
}

type Paginate struct {
	Total       int64 `json:"total"`
	PerPage     int   `json:"per_page"`
	CurrentPage int   `json:"current_page"`
	LastPage    int   `json:"last_page"`
}

type Sort struct {
	SortBy   string `json:"sort_by"`
	SortType string `json:"sort_type"`
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
		if param.IsPaginate == true {
			response = ResponseWithPaginate{
				Status:     status,
				StatusCode: param.StatusCode,
				Message:    param.Message,
				Data:       param.Data,
				Pagination: Paginate{
					Total:       param.Total,
					PerPage:     param.PerPage,
					CurrentPage: param.CurrentPage,
					LastPage:    param.LastPage,
				},
				Sorting: Sort{
					SortBy:   param.SortBy,
					SortType: param.SortType,
				},
			}

		} else {
			response = ResponseWithData{
				Status:     status,
				StatusCode: param.StatusCode,
				Message:    param.Message,
				Data:       param.Data,
			}
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
