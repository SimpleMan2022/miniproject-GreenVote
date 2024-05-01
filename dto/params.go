package dto

type ResponseParams struct {
	StatusCode  int    `json:"status_code"`
	Message     string `json:"message"`
	Data        any    `json:"data"`
	IsPaginate  bool   `json:"is_paginate"`
	Total       int64  `json:"Total"`
	PerPage     int    `json:"per_page"`
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	SortBy      string `json:"sort_by"`
	SortType    string `json:"sort_type"`
}

type ResponseError struct {
	Status     bool   `json:"status"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
}
