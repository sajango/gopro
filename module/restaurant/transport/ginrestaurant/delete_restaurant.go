package ginrestaurant

import (
	"food-delivery/common"
	bizrestaurant "food-delivery/module/restaurant/biz"
	restaurantstorage "food-delivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteRestaurant(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := bizrestaurant.NewDeleteRestaurantBiz(store)
		if err := biz.DeleteRestaurant(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(200, common.SimpleSuccessResponse(true))
	}
}
