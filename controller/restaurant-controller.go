package controller

import (
	"net/http"
	"rate-my-restaurant/entity"
	"rate-my-restaurant/service"

	"github.com/gin-gonic/gin"
)

type RestaurantController interface {
	FindAll() *[]entity.Restaurant
	Save(ctx *gin.Context) error
	ShowAll(ctx *gin.Context)
	CreateRestaurant(ctx *gin.Context) error
	FindRestaurant(ctx *gin.Context) (result *gin.Context)
}

type controller struct {
	service service.RestaurantService
}

func New(service service.RestaurantService) RestaurantController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() *[]entity.Restaurant {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) error {
	var restaurant entity.Restaurant
	err := ctx.ShouldBindJSON(&restaurant)
	if err != nil {
		return err
	}
	c.service.Save(restaurant)
	return nil
}

func (c *controller) ShowAll(ctx *gin.Context) {
	restaurants := c.service.FindAll()
	data := gin.H{
		"title":       "Restaurant Page",
		"restaurants": restaurants,
	}
	ctx.HTML(http.StatusOK, "restaurants.html", data)
}

func (c *controller) CreateRestaurant(ctx *gin.Context) error {

	var newRestaurant entity.Restaurant
	err := ctx.ShouldBindJSON(&newRestaurant)
	if err != nil {
		return err
	}
	c.service.CreateRestaurant(&newRestaurant)
	return nil
}

func (c *controller) FindRestaurant(ctx *gin.Context) (result *gin.Context) {
	restaurantName := ctx.Request.URL.Query()
	/*if restaurantName.Get("restaurant") == "" {
		ctx.JSON(http.StatusBadRequest, "no restaurant")
		return
	}

	*/
	restaurant, err := c.service.FindRestaurant(restaurantName.Get("restaurant"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, restaurant)
	return ctx
}
