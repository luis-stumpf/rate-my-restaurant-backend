package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"rate-my-restaurant/controller"
	"rate-my-restaurant/domain"
	"rate-my-restaurant/service"
)

var (
	restaurantService    service.RestaurantService       = service.New()
	restaurantController controller.RestaurantController = controller.New(restaurantService)
)

func main() {

	domain.ConnDB()

	server := gin.Default()

	config := cors.DefaultConfig()
	config.AddAllowHeaders("restaurantid")

	server.Use(cors.Default())

	server.LoadHTMLGlob("templates/*.html")

	apiRoutes := server.Group("/api")
	{

		apiRoutes.GET("/restaurants", func(ctx *gin.Context) {
			ctx.JSON(200, restaurantController.FindAll())
		})

		apiRoutes.POST("/restaurants", func(ctx *gin.Context) {
			err := restaurantController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Valid"})

			}

		})
		apiRoutes.POST("/create", func(ctx *gin.Context) {
			err := restaurantController.CreateRestaurant(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "User created"})

			}
		})

		apiRoutes.POST("/addDish", func(ctx *gin.Context) {
			err := restaurantController.AddDish(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Dish added"})
			}
		})

		apiRoutes.GET("/find", func(ctx *gin.Context) {
			restaurantController.FindRestaurant(ctx)
		})
	}

	server.GET("/main", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	server.GET("/lorenz", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "lorenz.html", nil)
	})

	server.Run(":8080")
}
