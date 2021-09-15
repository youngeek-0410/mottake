package controllers

import (
	"regexp"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/config"
	"github.com/youngeek-0410/mottake/server/geocoder"
)

var (
	NameRegexp        = regexp.MustCompile(`[^£$§¡€¢§¶ªº«\\/]{1,30}`)
	DescriptionRegexp = regexp.MustCompile(`[^£$§¡€¢§¶ªº«\\/]{1,255}`)
)

type APIError struct {
	StatusCode   int
	ErrorMessage string
}

func getUID(c *gin.Context) string {
	token, _ := c.Get("token")
	uid := token.(*auth.Token).UID
	return uid
}

func AddressToCoordinate(address string) (geocoder.Coordinate, error) {
	c := config.Config
	clientID := c.GeocoderClientID
	client, err := geocoder.New(clientID)
	var init geocoder.Coordinate
	if err != nil {
		return init, err
	}
	response, err := client.Search(geocoder.RequestParam{"query": address})
	if err != nil {
		return init, err
	}
	result := response.Feature[0]
	return result.Geometry.Coordinates.ToCoordinate(), nil
}
