package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {

	log.Println("in helloHandler")
	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func MiddleWareAuthen(c *gin.Context) {
	log.Println("start middleware") // Do Midleware

	//Find header: Authorization
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
		c.Abort()

		return
	}

	c.Next()                        // Do Handler
	log.Println("end middleware")   // Do Middleware
}

func Kaioo(c *gin.Context) {
	log.Println("in KaiooHandler")
	c.JSON(http.StatusOK, gin.H{
		"name": "Kaioo",
	})
}

func main() {

	r := gin.Default()

	//Do only helloHandler
	r.GET("/login", helloHandler)

	//Do both Middleware and Kaioo
	r.Use(MiddleWareAuthen)
	r.GET("/statements", Kaioo)

	r.Run(":1234")
}
