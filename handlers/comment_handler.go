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

type commentHandler struct {
	usecase usecases.CommentUsecase
}

func NewCommentHandler(usecase usecases.CommentUsecase) *commentHandler {
	return &commentHandler{usecase}
}

func (h *commentHandler) CreateComment(ctx echo.Context) error {
	var request dto.CommentRequest
	if err := ctx.Bind(&request); err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	idPlace := ctx.Param("id")
	placeId, err := uuid.Parse(idPlace)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := helpers.ValidateRequest(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, dto.ResponseError{
			Status:     false,
			StatusCode: http.StatusBadRequest,
			Message:    "Failed to create comment. please ensure your input correctly",
			Data:       err,
		})
	}
	idUser := ctx.Get("userId")
	userId := idUser.(*uuid.UUID)
	newComment, err := h.usecase.Create(*userId, placeId, &request)
	if err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	createResponse := dto.ToCommentCreateResponse(newComment)
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusCreated,
		Message:    "Create successful. Comment has been created.",
		Data:       createResponse,
	})

	return ctx.JSON(http.StatusCreated, response)
}

func (h *commentHandler) GetAllCommentsInPlace(ctx echo.Context) error {
	idPlace := ctx.Param("id")
	placeId, err := uuid.Parse(idPlace)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	comments, placeDetail, err := h.usecase.GetAllCommentInPlace(placeId)

	allResponse := dto.CommentFindAllResponse{
		PlaceDetail: dto.CommentDetail{
			PlaceName:   placeDetail.PlaceName,
			Province:    placeDetail.Province,
			City:        placeDetail.City,
			SubDistrict: placeDetail.SubDistrict,
			StreetName:  placeDetail.StreetName,
		},
		Comments: comments,
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Successfully retrieved place comments",
		Data:       allResponse,
	})
	return ctx.JSON(http.StatusOK, response)
}

func (h *commentHandler) DeleteComment(ctx echo.Context) error {
	idPlace := ctx.Param("placeId")
	placeId, err := uuid.Parse(idPlace)

	idComment := ctx.Param("commentId")
	commentId, err := uuid.Parse(idComment)

	idUser := ctx.Get("userId")
	userId := idUser.(*uuid.UUID)
	if err != nil {
		return errorHandlers.HandleError(ctx, &errorHandlers.BadRequestError{err.Error()})
	}
	if err := h.usecase.Delete(commentId, *userId, placeId); err != nil {
		return errorHandlers.HandleError(ctx, err)
	}
	response := helpers.Response(dto.ResponseParams{
		StatusCode: http.StatusOK,
		Message:    "Delete successful. Comment has been deleted",
	})
	return ctx.JSON(http.StatusOK, response)
}
