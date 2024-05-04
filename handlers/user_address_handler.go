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

type UserAddressHandler struct {
	usecaseAddress usecases.UserAddressUsecase
}

func NewUserAddressHandler(usecaseAddress usecases.UserAddressUsecase) *UserAddressHandler {
	return &UserAddressHandler{usecaseAddress}
}

func (h *UserAddressHandler) Create(ctx echo.Context) error {
	var request dto.CreateUserAddressRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.InternalServerError{err.Error()})
	}
	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create address. please ensure your input correctly",
			Data:       err,
		})
	}

	id := ctx.Get("userId")
	userId := id.(*uuid.UUID)
	request.UserId = *userId
	newAddress, err := h.usecaseAddress.Create(&request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	addressResponse := dto.ToUserAddressResponse(newAddress)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Create successful. User address information has been created.",
		Data:       addressResponse,
	})
	return ctx.JSON(http.StatusCreated, response)
}

func (h *UserAddressHandler) Update(ctx echo.Context) error {

	var request dto.CreateUserAddressRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.InternalServerError{err.Error()})
	}
	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update address. please ensure your input correctly",
			Data:       err,
		})
	}

	id := ctx.Get("userId")
	userId := id.(*uuid.UUID)
	request.UserId = *userId
	newAddress, err := h.usecaseAddress.Update(&request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	addressResponse := dto.ToUserAddressResponse(newAddress)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Update successful. User address information has been updated.",
		Data:       addressResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *UserAddressHandler) Delete(ctx echo.Context) error {

	id := ctx.Get("userId")
	userId := id.(*uuid.UUID)
	err := h.usecaseAddress.Delete(*userId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Delete successful. User address information has been deleted.",
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *UserAddressHandler) GetDetailUser(ctx echo.Context) error {
	id := ctx.Get("userId")
	userId := id.(*uuid.UUID)
	user, err := h.usecaseAddress.GetDetailUser(*userId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	byIdResponse := dto.ToByIdResponse(user)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved user data",
		Data:       byIdResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}
