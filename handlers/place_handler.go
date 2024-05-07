package handlers

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/helpers"
	"evoting/usecases"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"math"
	"net/http"
	"strconv"
)

type placeHandler struct {
	usecase usecases.PlaceUsecase
}

func NewPlaceHandler(uc usecases.PlaceUsecase) *placeHandler {
	return &placeHandler{uc}
}

func (h *placeHandler) FindAllPlaces(ctx echo.Context) error {
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
	places, totalPtr, err := h.usecase.FindAll(page, limit, sortBy, sortType)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	total := *totalPtr
	lastPage := int(math.Ceil(float64(total) / float64(limit)))
	if page > lastPage {
		page = lastPage
	}

	usersResponse := dto.ToFindAllPlacesResponse(places)
	response := helpers.Response(dto.ResponseParams{
		StatusCode:  http.StatusOK,
		Message:     "Successfully retrieved places data",
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

func (h *placeHandler) FindPlaceById(ctx echo.Context) error {
	id := ctx.Param("id")
	placeId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	place, err := h.usecase.FindById(placeId)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	byIdResponse := dto.ToPlaceByIdResponse(place)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved user data",
		Data:       byIdResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *placeHandler) CreatePlace(ctx echo.Context) error {
	var request dto.PlaceRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create place. please ensure your input correctly",
			Data:       err,
		})
	}
	place, address, err := helpers.GenerateLocationDetail(&request)
	request.Latitude = place.Latitude
	request.Longitude = place.Longitude
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	newPlace, err := h.usecase.Create(&request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	address.PlaceId = newPlace.Id

	newAddress, err := h.usecase.CreateAddress(address)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	placeResponse := dto.ToPlaceResponse(newPlace, newAddress)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Create successful. Place information has been created.",
		Data:       placeResponse,
	})
	return ctx.JSON(http.StatusCreated, response)
}

func (h *placeHandler) Delete(ctx echo.Context) error {
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
		Message:    "Delete successful. User information has been deleted.",
		Data:       nil,
	})

	return ctx.JSON(http.StatusOK, response)
}

func (h *placeHandler) UpdatePlace(ctx echo.Context) error {
	id := ctx.Param("id")
	placeId, err := uuid.Parse(id)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	var request dto.PlaceRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create place. please ensure your input correctly",
			Data:       err,
		})
	}
	place, address, err := helpers.GenerateLocationDetail(&request)
	request.Latitude = place.Latitude
	request.Longitude = place.Longitude
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	newPlace, err := h.usecase.Update(placeId, &request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	address.PlaceId = newPlace.Id

	newAddress, err := h.usecase.UpdateAddress(address)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}

	placeResponse := dto.ToPlaceResponse(newPlace, newAddress)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Update successful. Place information has been updated.",
		Data:       placeResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}
