package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	webpush "github.com/SherClockHolmes/webpush-go"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const vapidPrivateKey = "KMYVkbkzylo5XATcUcnX2VNr7P68sENwDw3w7eSEmrI"

func main() {
	server := gin.Default()
	server.Use(cors.Default())

	server.GET("/ping", getPong)
	server.GET("/notify", getNotify)

	server.Run() // listen and serve on 0.0.0.0:8080
}

func getPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func getNotify(c *gin.Context) {
	body := c.Query("body")
	sub := c.Query("sub")

	subJSON := webpush.Subscription{}
	if err := json.NewDecoder(bytes.NewBufferString(sub)).Decode(&subJSON); err != nil {
		log.Fatal(err)
	}

	// Send Notification
	_, err := webpush.SendNotification([]byte(body), &subJSON, &webpush.Options{
		VAPIDPrivateKey: vapidPrivateKey,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "failed",
		})
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "sent",
	})
}
