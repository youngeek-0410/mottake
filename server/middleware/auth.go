package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/youngeek-0410/mottake/server/config"
	"google.golang.org/api/option"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

var AuthClient *auth.Client

func Auth() gin.HandlerFunc {

	var err error

	opt := option.WithCredentialsFile(config.Config.FirebaseSecret)
	App, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatal(err)
	}
	AuthClient, err = App.Auth(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		idToken := strings.ReplaceAll(authHeader, "Bearer ", "")
		token, err := AuthClient.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			log.Print("Invalid token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}
		c.Set("token", token)
		c.Next()
	}
}

func Dummy() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := &auth.Token{UID: "dummy"}
		c.Set("token", token)
		c.Next()
	}
}
