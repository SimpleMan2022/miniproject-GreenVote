package handlers

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type weatherDataHandler struct {
	usecase usecases.WeatherDataUsecase
}

func NewWeatherDataHandler(usecase usecases.WeatherDataUsecase) *weatherDataHandler {
	return &weatherDataHandler{usecase}
}

func (h *weatherDataHandler) Create(ctx echo.Context) error {
	var request dto.WeatherDataRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	weatherData, err := h.usecase.Create(&request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	place, err := h.usecase.GetPlace(weatherData.PlaceId)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.InternalServerError{Message: err.Error()})
	}

	weatherResponse := dto.ToWeatherDataResponse(weatherData, place)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Create successful. Weather information has been created.",
		Data:       weatherResponse,
	})

	return ctx.JSON(http.StatusCreated, response)
}

func (h *weatherDataHandler) Update(ctx echo.Context) error {
	id := ctx.Param("id")
	placeId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	weatherData, err := h.usecase.Update(placeId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	place, err := h.usecase.GetPlace(placeId)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.InternalServerError{Message: err.Error()})
	}

	weatherResponse := dto.ToWeatherDataResponse(weatherData, place)

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Update successful. Weather information has been updated.",
		Data:       weatherResponse,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *weatherDataHandler) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	placeId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := h.usecase.Delete(placeId); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Delete successful. Weather information has been deleted.",
	})

	return ctx.JSON(http.StatusOK, response)
}
