package bizrestaurant

import (
	"context"
	"errors"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	FinRestaurantWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Delete(ctx context.Context, id int) error
}
type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id int) error {
	oldData, err := biz.store.FinRestaurantWithCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}

	if oldData.Status == 0 {
		return errors.New("Can not found restaurant")
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}
