package geocoder

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

const (
	APIEndpoint = "https://map.yahooapis.jp/geocode/V1/geoCoder"
)

type Client struct {
	httpClient *http.Client
	clientID   string
	endpoint   string
}

type Response struct {
	ResultInfo struct {
		Count       int
		Total       int
		Start       int
		Status      int
		Description string
		Copyright   string
		Latency     float32
	}

	Feature []struct {
		Id       string
		Name     string
		Geometry struct {
			Type        string
			Coordinates CoordinateString
		}
		Property struct {
			Uid        string
			CassetteId string
			Yomi       string
			Country    struct {
				Code string
				Name string
			}
			Address              string
			GovernmentCode       int
			AddressMatchingLevel int
		}
	}
}

type RequestParam map[string]string
type CoordinateString string
type Coordinate struct {
	Latitude  float32
	Longitude float32
}

func New(clientID string) (*Client, error) {
	if clientID == "" {
		return nil, errors.New("missing client id")
	}
	c := &Client{
		httpClient: http.DefaultClient,
		clientID:   clientID,
		endpoint:   APIEndpoint,
	}
	return c, nil
}

func (c Client) Search(params RequestParam) (*Response, error) {
	var decodedResponse Response
	request, err := http.NewRequest(http.MethodGet, c.endpoint, nil)
	if err != nil {
		return nil, err
	}
	query := request.URL.Query()
	query.Set("appid", c.clientID)
	for key, value := range params {
		query.Set(key, value)
	}
	request.URL.RawQuery = query.Encode()
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	err = xml.NewDecoder(response.Body).Decode(&decodedResponse)
	if err != nil {
		return nil, err
	}
	return &decodedResponse, nil

}

func (s CoordinateString) ToCoordinate() Coordinate {
	var coordinate Coordinate
	fmt.Sscanf(string(s), "%f,%f", &coordinate.Longitude, &coordinate.Latitude)
	return coordinate
}
