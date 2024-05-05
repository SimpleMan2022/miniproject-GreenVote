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

type placeAddressHandler struct {
	usecase usecases.PlaceAddressUsecase
}

func NewPlaceAddressHandler(usecase usecases.PlaceAddressUsecase) *placeAddressHandler {
	return &placeAddressHandler{usecase}
}

func (h *placeAddressHandler) CreateAddress(ctx echo.Context) error {
	var request dto.PlaceAddressRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}

	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create place address. please ensure your input correctly",
			Data:       err,
		})
	}

	newAddress, err := h.usecase.Create(&request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	addressResponse := dto.ToPlaceAddressResponse(newAddress)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Create successful. place address information has been created.",
		Data:       addressResponse,
	})
	return ctx.JSON(http.StatusCreated, response)

}

func (h *placeAddressHandler) UpdateAddress(ctx echo.Context) error {
	var request dto.PlaceAddressRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	id := ctx.Param("id")
	placeId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}

	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update place address. please ensure your input correctly",
			Data:       err,
		})
	}

	updateAddress, err := h.usecase.Update(placeId, &request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	addressResponse := dto.ToPlaceAddressResponse(updateAddress)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Update successful. place address information has been updated.",
		Data:       addressResponse,
	})
	return ctx.JSON(http.StatusOK, response)

}

func (h *placeAddressHandler) DeleteAddress(ctx echo.Context) error {
	var request dto.PlaceAddressRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
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
		Message:    "Delete successful. place address information has been deleted.",
	})
	return ctx.JSON(http.StatusOK, response)
}
