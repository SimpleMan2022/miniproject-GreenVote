package handlers

import (
	"evoting/dto"
	"evoting/errorHandlers"
	"evoting/usecases"
	"github.com/labstack/echo/v4"
)

type PlaceUsecase struct {
	usecase usecases.PlaceUsecase
}

func NewPlaceUsacase(uc usecases.PlaceUsecase) *PlaceUsecase {
	return &PlaceUsecase{uc}
}

func (h *PlaceUsecase) CreatePlace(ctx echo.Context) error {
	var user dto.PlaceRequest
	if err := ctx.Bind(&user); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}

}
