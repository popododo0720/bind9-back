package main

import (
	"main/controller"
	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	api := r.Group("/api/v1")
    {
        api.GET("/test1", controller.ControllerTest1)
		api.GET("/test2", controller.ControllerTest2)
    }
	
	return r
} 

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
