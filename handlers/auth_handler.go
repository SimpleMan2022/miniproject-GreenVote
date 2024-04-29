package handlers

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

type authHandler struct {
	usecase usecases.AuthUsecase
}

func NewAuthHandler(uc usecases.AuthUsecase) *authHandler {
	return &authHandler{uc}
}
func (h *authHandler) Register(ctx echo.Context) error {
	var user dto.UserRequest
	if err := ctx.Bind(&user); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Register failed, please fill input correctly",
			Data:       err,
		})
	}
	newUser, err := h.usecase.Register(&user)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Register user successfully",
		Data:       newUser,
	})

	return ctx.JSON(http.StatusCreated, response)
}

func (h *authHandler) Login(ctx echo.Context) error {
	var user dto.LoginRequest
	if err := ctx.Bind(&user); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Login failed, please fill input correctly",
			Data:       err,
		})
	}

	result, err := h.usecase.Login(&user)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	ctx.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    result.RefreshToken,
		Path:     "/",
		Domain:   "",
		MaxAge:   24 * 60 * 60,
		Secure:   true,
		HttpOnly: true,
	})
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Login successfully",
		Data:       result.AccessToken,
	})
	return ctx.JSON(http.StatusOK, response)
}
