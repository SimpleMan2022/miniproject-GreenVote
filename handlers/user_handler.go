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

type userHandler struct {
	usecase usecases.UserUsecase
}

func NewUserHandler(uc usecases.UserUsecase) *userHandler {
	return &userHandler{uc}
}
func (h *userHandler) Create(ctx echo.Context) error {
	var user dto.CreateRequest
	if err := ctx.Bind(&user); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to register. please ensure your input correctly",
			Data:       err,
		})
	}
	newUser, err := h.usecase.Create(&user)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Congratulations! Your registration was successful. Please login to continue",
		Data:       newUser,
	})

	return ctx.JSON(http.StatusCreated, response)
}

func (h *userHandler) Login(ctx echo.Context) error {
	var user dto.LoginRequest
	if err := ctx.Bind(&user); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to login. please ensure your input correctly",
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
		Message:    "You have successfully logged in",
		Data:       result.AccessToken,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) FindUserById(ctx echo.Context) error {
	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	user, err := h.usecase.FindById(userId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved user data",
		Data:       user,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) FindAllUsers(ctx echo.Context) error {
	users, err := h.usecase.FindAll()
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved user data",
		Data:       users,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(ctx echo.Context) error {
	var user dto.UpdateRequest
	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := ctx.Bind(&user); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}

	if err := helpers.ValidateRequest(user); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to update user. please ensure your input correctly",
			Data:       err,
		})
	}
	updateUser, err := h.usecase.Update(userId, &user)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Update successful. User information has been updated.",
		Data:       updateUser,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) DeleteUser(ctx echo.Context) error {
	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}

	if err := h.usecase.Delete(userId); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Delete successful. User information has been deleted.",
		Data:       nil,
	})
	return ctx.JSON(http.StatusOK, response)
}
