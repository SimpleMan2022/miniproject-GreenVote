package usecases

import (
	"encoding/json"
	"evoting/dto"
	"evoting/entities"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type ChatbotUsecase interface {
	GetRecommendation(query string, chatHistory []dto.ChatHistory, allPlaces *[]entities.Place) (string, error)
}

type chatbotUsecase struct {
}

func NewChatbotUseCase() *chatbotUsecase {
	return &chatbotUsecase{}
}

func (uc *chatbotUsecase) GetRecommendation(query string, chatHistory []dto.ChatHistory, allPlaces *[]entities.Place) (string, error) {
	messages := append([]dto.ChatHistory{}, chatHistory...)
	messages = append(messages, dto.ChatHistory{PreviousMessages: query})

	var prompt string
	prompt = "You are assistant for give all information about the places. given the location data to be voted:"
	for i, place := range *allPlaces {
		prompt += fmt.Sprintf("%d. %s, description: %s, temperature: %d. ", i+1, place.Name, place.Description, place.Weather.Temperature)
	}
	prompt += fmt.Sprintf("question: %s", messages)

	url := "https://wgpt-production.up.railway.app/v1/chat/completions"
	payload := fmt.Sprintf(`{"model": "gpt-3.5-turbo", "messages": [{"role": "system", "content": "%s"}], "stream": false}`, prompt)
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var aiResponse dto.ChatCompletionResponse
	if err := json.Unmarshal(body, &aiResponse); err != nil {
		return "", err
	}

	if len(aiResponse.Choices) > 0 {
		return aiResponse.Choices[0].Message.Content, nil
	}

	return "", fmt.Errorf("no completions found")
}
