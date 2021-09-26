package controllers

import (
	"errors"
	"regexp"

	"firebase.google.com/go/v4/auth"
	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/config"
	"github.com/youngeek-0410/mottake/server/geocoder"
)

var (
	NameRegexp        = regexp.MustCompile(`[^£$§¡€¢§¶ªº«\\/]{1,30}`)
	DescriptionRegexp = regexp.MustCompile(`[^£$§¡€¢§¶ªº«\\/]{1,255}`)

	errInvalidJSONRequest = "invalid json request"
	errCouldNotCreateMenu = "could not create menu"
	errCouldNotGetMenus   = "could not get menus"
	errCouldNotGetMenu    = "could not get menu"
	errInvalidMenuID      = "invalid menu_id"
	errCouldNotDeleteMenu = "could not delete menu"
	errCouldNotUpdateMenu = "could not update menu"

	errCouldNotCreateReceipt = "could not create receipt"
	errCouldNotQueryReceipts = "could not query receipts"

	errNotFound           = "Not Found"
	errCouldNotCreateShop = "Could Not Create Shop"
	errInvalidName        = "Invalid character or length (1~30)"
	errInvalidDescription = "Invalid character of length (1~255)"
	errInvalidAddress     = "Invalid Address. Please enter a proper address in Japanese"
	errCouldNotUpdateShop = "Could Not Update Shop"
	errNotAuthorized      = "Not Authorized"
	errCouldNotDeleteShop = "Could Not Delete Shop"

	errCouldNotUpdateGenre = "Could Not Update Genre"

	errCouldNotCreateCustomer = "Could Not Create Customer"
	errCouldNotUpdateCustomer = "Could Not Update Customer"
	errCouldNotDeleteCustomer = "Could Not Delete Customer"
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
	if response.ResultInfo.Count <= 0 {
		return init, errors.New("Result infomation is zero.")
	}
	result := response.Feature[0]
	return result.Geometry.Coordinates.ToCoordinate(), nil
}
