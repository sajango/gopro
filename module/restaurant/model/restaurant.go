package restaurantmodel

import "food-delivery/common"

type Restaurant struct {
	common.SQLModel
	Name   string `json:"name" gorm:"column:name;"`
	Addr   string `json:"addr" gorm:"column:addr;"`
	Status int    `json:"status" gorm:"column:status;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
