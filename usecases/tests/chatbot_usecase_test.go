// package tests

// import (
// 	"evoting/dto"
// 	"evoting/entities"
// 	"evoting/usecases"
// 	"fmt"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// )

// type mockHTTPClient struct {
// 	mock.Mock
// }

// func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
// 	args := m.Called(req)
// 	return args.Get(0).(*http.Response), args.Error(1)
// }

// func TestGetRecommendation(t *testing.T) {
// 	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintln(w, `{"choices":[{"message":{"content":"Recommended place"}}]}`)
// 	}))
// 	defer server.Close()

// 	chatbotUsecase := usecases.NewChatbotUseCase()

// 	query := "Some query"
// 	chatHistory := []dto.ChatHistory{}
// 	allPlaces := &[]entities.Place{
// 		{Name: "Place 1", Description: "Description 1", Weather: entities.WeatherData{Temperature: 25}},
// 		{Name: "Place 2", Description: "Description 2", Weather: entities.WeatherData{Temperature: 30}},
// 	}

// 	recommendation, err := chatbotUsecase.GetRecommendation(query, chatHistory, allPlaces)
// 	if err != nil {
// 		t.Fatalf("GetRecommendation failed: %v", err)
// 	}

// 	assert.NotNil(t, recommendation)
// }
