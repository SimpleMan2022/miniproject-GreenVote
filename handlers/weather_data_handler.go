package handlers

import (
	"evoting/usecases"
)

type weatherDataHandler struct {
	usecase usecases.WeatherDataUsecase
}

func NewWeatherDataHandler(usecase usecases.WeatherDataUsecase) *weatherDataHandler {
	return &weatherDataHandler{usecase}
}

//func (h *weatherDataHandler) CreateOrUpdate(ctx echo.Context) error {
//	var request *dto.WeatherDataRequest
//	id := ctx.Param("id")
//	placeid, err := uuid.Parse(id)
//	if err != nil {
//		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
//	}
//}
