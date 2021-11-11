package main

import(
	c "app/controllers"
	"github.com/gin-gonic/gin"
)

func main(){

	router := gin.Default()

	router.POST("/signup", c.SignUp)
	router.POST("/login", c.Login)
	router.POST("/change", c.ChangePassword)
	router.GET("/echo", c.WebSocket)

	router.Run("localhost:9090")
}