package main

import (
	"os"

	"github.com/aicomylleville/formula-mars_go-api/controllers"
	"github.com/aicomylleville/formula-mars_go-api/middlewares"
	"github.com/aicomylleville/formula-mars_go-api/models"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDatabase()

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
	}))

	public := router.Group("/api")
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	users := router.Group("/api/user")
	users.Use(middlewares.JwtAuthMiddleware())
	users.GET("/", controllers.CurrentUser)

	races := router.Group("/api/races")
	races.Use(middlewares.JwtAuthMiddleware())
	races.GET("/", controllers.GetRaces)
	races.POST("/", controllers.AddRace)
	races.GET("/:id", controllers.GetRaceByID)
	races.PUT("/:id", controllers.UpdateRace)
	races.DELETE("/:id", controllers.DeleteRace)

	championships := router.Group("/api/championships")
	championships.Use(middlewares.JwtAuthMiddleware())
	championships.GET("/", controllers.GetChampionships)
	championships.POST("/", controllers.AddChampionship)
	championships.GET("/:id", controllers.GetChampionshipById)
	championships.PUT("/:id", controllers.UpdateChampionship)
	championships.DELETE("/:id", controllers.DeleteChampionship)

	drivers := router.Group("/api/drivers")
	drivers.Use(middlewares.JwtAuthMiddleware())
	drivers.GET("/", controllers.GetDrivers)
	drivers.POST("/", controllers.AddDriver)
	drivers.GET("/:id", controllers.GetDriverById)
	drivers.PUT("/:id", controllers.UpdateDriver)
	drivers.DELETE("/:id", controllers.DeleteDriver)

	if os.Getenv("PORT") != "" {
		router.Run("0.0.0.0:" + os.Getenv("PORT"))
	} else {
		router.Run("0.0.0.0:8000")
	}
}
