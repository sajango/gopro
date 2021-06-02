package bizrestaurant

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type FindRestaurantStore interface {
	FinRestaurantWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}
type findRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewFindRestaurantBiz(store DeleteRestaurantStore) *findRestaurantBiz {
	return &findRestaurantBiz{store: store}
}

func (biz *findRestaurantBiz) FindRestaurant(ctx context.Context, id int) (*restaurantmodel.Restaurant, error) {
	data, err := biz.store.FinRestaurantWithCondition(ctx, map[string]interface{}{"id": id, "status": 1})
	if err != nil {
		return nil, err
	}
	return data, err
}
