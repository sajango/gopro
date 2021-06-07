package restaurantmodel

import "food-delivery/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel
	Name   string        `json:"name" gorm:"column:name;"`
	Addr   string        `json:"addr" gorm:"column:addr;"`
	Status int           `json:"status" gorm:"column:status;"`
	Logo   *common.Image `json:"logo" gorm:"column:logo;"`
	Cover  *common.Image `json:"cover" gorm:"column:logo;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GentUID(common.DbTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel
	Name  string        `json:"name" gorm:"column:name;"`
	Addr  string        `json:"addr" gorm:"column:addr;"`
	Logo  *common.Image `json:"logo" gorm:"column:logo;"`
	Cover *common.Image `json:"cover" gorm:"column:logo;"`
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GentUID(common.DbTypeRestaurant)
}

type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
