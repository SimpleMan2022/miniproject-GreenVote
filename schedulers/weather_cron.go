package schedulers

import (
	"evoting/config"
	"evoting/repositories"
	"evoting/usecases"
	"fmt"
	"github.com/go-co-op/gocron"
	"log"
	"time"
)

func StartScheduler() {
	jakartaTime, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		jakartaTime = time.UTC
	}
	s := gocron.NewScheduler(jakartaTime)
	s.Every(1).Day().At("16:02").Do(updateAllPlaces)
	s.StartBlocking()
}

func updateAllPlaces() {
	placeRepository := repositories.NewPlaceRepository(config.DB)
	placeUsecase := usecases.NewPlaceUsecase(placeRepository)

	weatherRepository := repositories.NewWeatherDataRepository(config.DB)
	weatherUsecase := usecases.NewWeatherDataUsecase(weatherRepository)

	allPlaces, _, err := placeUsecase.FindAll(1, 10, "", "", "")
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, place := range *allPlaces {
		_, err := weatherUsecase.Update(place.Id)
		if err != nil {
			log.Printf("Failed to update weather data for place %s: %v", place.Name, err)
			continue
		}

	}

	fmt.Println("Weather data update for all places completed")
}
