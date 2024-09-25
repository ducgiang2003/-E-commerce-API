package main

import (
	"golang_api/jwt-authentication-golang/controllers"
	"golang_api/jwt-authentication-golang/database"
	"golang_api/jwt-authentication-golang/milldleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	database.Connect("root:@tcp(localhost:3306)/jwt_demo?parseTime=true")
	database.Migrate()

	// Initialize Router
	router := initRouter()
	router.Run(":8080")
}
func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api", controllers.RegisterUser)
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(milldleware.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}
