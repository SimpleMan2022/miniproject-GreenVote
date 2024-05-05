package helpers

import (
	"encoding/json"
	"evoting/dto"
	"evoting/entities"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type ResponseMapsApi struct {
	ResourceSets []struct {
		Resources []struct {
			Name    string `json:"name"`
			Address struct {
				AdminDistrict    string `json:"adminDistrict"`
				AdminDistrict2   string `json:"adminDistrict2"`
				CountryRegion    string `json:"countryRegion"`
				FormattedAddress string `json:"formattedAddress"`
				Locality         string `json:"locality"`
			} `json:"address"`
			Point struct {
				Coordinates []float64 `json:"coordinates"`
			} `json:"point"`
		} `json:"resources"`
	} `json:"resourceSets"`
}

func init() {
	viper.AutomaticEnv()
}

func GenerateLocationDetail(request *dto.PlaceRequest) (*entities.Place, *entities.PlaceAddress, error) {
	apikey := viper.GetString("MAPS_API_KEY")
	geocodingUrl := viper.GetString("MAPS_GEOCODING_URL")
	query := request.Name

	url := fmt.Sprintf(geocodingUrl, query, apikey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	var response ResponseMapsApi
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, nil, err
	}

	var res *struct {
		Name    string `json:"name"`
		Address struct {
			AdminDistrict    string `json:"adminDistrict"`
			AdminDistrict2   string `json:"adminDistrict2"`
			CountryRegion    string `json:"countryRegion"`
			FormattedAddress string `json:"formattedAddress"`
			Locality         string `json:"locality"`
		} `json:"address"`
		Point struct {
			Coordinates []float64 `json:"coordinates"`
		} `json:"point"`
	}
	for _, r := range response.ResourceSets[0].Resources {
		if r.Address.AdminDistrict2 != "" && r.Address.Locality != "" {
			res = &r
			break
		}
	}

	if res == nil {
		if len(response.ResourceSets) > 0 && len(response.ResourceSets[0].Resources) > 0 {
			res = &response.ResourceSets[0].Resources[0]
		} else {
			return nil, nil, fmt.Errorf("no address data found")
		}
	}

	place := &entities.Place{
		Name:      res.Name,
		Longitude: res.Point.Coordinates[0],
		Latitude:  res.Point.Coordinates[1],
	}
	address := &entities.PlaceAddress{
		Province:    res.Address.AdminDistrict,
		City:        res.Address.AdminDistrict2,
		SubDistrict: res.Address.Locality,
	}
	return place, address, nil
}

func GenerateImageLocation(place *dto.PlaceRequest) string {
	longLat := fmt.Sprintf("%.8f,%.8f", place.Longitude, place.Latitude)
	apikey := viper.GetString("MAPS_API_KEY")
	imageryUrl := viper.GetString("MAPS_IMAGERY_URL")
	url := fmt.Sprintf(imageryUrl, longLat, longLat, apikey)

	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	image, err := UploadPlaceImage(resp.Body)
	if err != nil {
		return ""
	}
	return image
}
