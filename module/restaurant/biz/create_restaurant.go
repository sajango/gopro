package bizrestaurant

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}
type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if data.Name == "" {
		return errors.New("name cannot empty")
	}
	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}
	return nil
}
