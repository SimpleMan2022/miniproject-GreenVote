package handlers

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
	"net/http"
)

type adminHandler struct {
	usecase usecases.AdminUsecase
}

func NewAdminHandle(uc usecases.AdminUsecase) *adminHandler {
	return &adminHandler{uc}
}

func (h *adminHandler) Login(ctx echo.Context) error {
	var admin dto.LoginAdminRequest
	if err := ctx.Bind(&admin); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(admin); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to login. Please ensure your input correctly",
			Data:       err,
		})
	}

	result, err := h.usecase.Login(&admin)
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
	loginResponse := dto.ToLoginAdminResponse(result)

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "You have successfully logged in",
		Data:       loginResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *adminHandler) Logout(ctx echo.Context) error {
	cookie, err := ctx.Cookie("refreshToken")
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.ForbiddenError{Message: err.Error()})
	}
	if err := h.usecase.Logout(cookie.Value); err != nil {
		return &errorHandlers.ForbiddenError{Message: err.Error()}
	}
	ctx.SetCookie(&http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		Domain:   "",
		MaxAge:   0,
		Secure:   true,
		HttpOnly: true,
	})
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Logout Success",
	})
	return ctx.JSON(http.StatusOK, response)
}
