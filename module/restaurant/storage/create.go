package restaurantstorage

import (
	"context"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStore) Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Table(restaurantmodel.RestaurantCreate{}.TableName()).Create(&data).Error; err != nil {
		return err
	}
	return nil
}
