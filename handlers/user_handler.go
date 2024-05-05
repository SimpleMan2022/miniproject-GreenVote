package handlers

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/usecases"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
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
	createResponse := dto.ToCreateResponse(newUser)

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Congratulations! Your registration was successful. Please login to continue",
		Data:       createResponse,
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
			Message:    "Failed to login. Please ensure your input correctly",
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
	loginResponse := dto.ToLoginResponse(result)

	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "You have successfully logged in",
		Data:       loginResponse,
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

	byIdResponse := dto.ToByIdResponse(user)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved user data",
		Data:       byIdResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) FindAllUsers(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	sortBy := ctx.QueryParam("sort_by")
	sortType := ctx.QueryParam("sort_type")
	if sortBy == "" {
		sortBy = "updated_at"
		sortType = "desc"
	}
	if sortType == "" {
		sortType = "asc"
	}
	users, totalPtr, err := h.usecase.FindAll(page, limit, sortBy, sortType)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	total := *totalPtr
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		page = lastPage
	}
	fmt.Println(total, limit, page, lastPage, sortBy, sortType)
	usersResponse := dto.ToFindAllResponse(users)
	response := helpers.Response(dto.ResponseParams{
		StatusCode:  http.StatusOK,
		Message:     "Successfully retrieved user data",
		Data:        usersResponse,
		IsPaginate:  true,
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
		SortBy:      sortBy,
		SortType:    sortType,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) FindAllUserWithSoftDelete(ctx echo.Context) error {
	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	if page == 0 {
		page = 1
	}
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	sortBy := ctx.QueryParam("sort_by")
	sortType := ctx.QueryParam("sort_type")
	if sortBy == "" {
		sortBy = "updated_at"
		sortType = "desc"
	}
	if sortType == "" {
		sortType = "asc"
	}
	users, totalPtr, err := h.usecase.FindSoftDelete(page, limit, sortBy, sortType)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	total := *totalPtr
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		page = lastPage
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode:  http.StatusOK,
		Message:     "Successfully retrieved user data with soft delete",
		Data:        users,
		IsPaginate:  true,
		Total:       total,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
		SortBy:      sortBy,
		SortType:    sortType,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *userHandler) UpdateUser(ctx echo.Context) error {
	id := ctx.Param("id")
	userId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	image, err := ctx.FormFile("image")

	user := dto.UpdateRequest{
		Email:    ctx.FormValue("email"),
		Fullname: ctx.FormValue("fullname"),
		Password: ctx.FormValue("password"),
		Image:    image,
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

	updateResponse := dto.ToByIdResponse(updateUser)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Update successful. User information has been updated.",
		Data:       updateResponse,
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
