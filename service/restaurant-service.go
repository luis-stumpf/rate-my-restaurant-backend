package service

import (
	"rate-my-restaurant/domain"
	"rate-my-restaurant/entity"
)

type RestaurantService interface {
	FindAll() *[]entity.Restaurant
	Save(entity.Restaurant) entity.Restaurant
	CreateRestaurant(*entity.Restaurant) (*entity.Restaurant, error)
	FindRestaurant(restaurantName string) (*entity.Restaurant, error)
}

type restaurantService struct {
	restaurants []entity.Restaurant
}

func New() RestaurantService {
	return &restaurantService{}
}

func (service *restaurantService) FindAll() *[]entity.Restaurant {
	restaurant := domain.FindAll()
	return restaurant
}

func (service *restaurantService) Save(restaurant entity.Restaurant) entity.Restaurant {

	service.restaurants = append(service.restaurants, restaurant)
	return restaurant
}

func (service *restaurantService) CreateRestaurant(restaurant *entity.Restaurant) (*entity.Restaurant, error) {
	restaurant, err := domain.Create(restaurant)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (service *restaurantService) FindRestaurant(restaurantName string) (*entity.Restaurant, error) {
	restaurant, err := domain.Find(restaurantName)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}
