package handlers

import (
	"evoting/dto"
	"evoting/usecases"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type chatbotHandler struct {
	ChatbotUseCase usecases.ChatbotUsecase
	PlaceUsecae    usecases.PlaceUsecase
}

func NewChatbotHandler(usecase usecases.ChatbotUsecase, place usecases.PlaceUsecase) *chatbotHandler {
	return &chatbotHandler{usecase, place}
}

func (h *chatbotHandler) HandleWebSocket(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Println("upgrade:", err)
		return err
	}
	defer ws.Close()

	allPlaces, _, err := h.PlaceUsecae.FindAll(1, 10, "", "", "")
	if err != nil {
		log.Println("read:", err)
		return err
	}

	var chatHistory []dto.ChatHistory
	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		chatHistory = append(chatHistory, dto.ChatHistory{PreviousMessages: string(msg)})
		fmt.Println("msg:", string(msg))

		recommendation, err := h.ChatbotUseCase.GetRecommendation(string(msg), chatHistory, allPlaces)
		if err != nil {
			log.Println("get recommendation:", err)
			break
		}

		if err := ws.WriteMessage(websocket.TextMessage, []byte(recommendation)); err != nil {
			log.Println("write:", err)
			break
		}
	}

	return nil
}
