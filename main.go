package main

import (
	"food-delivery/component/appctx"
	"food-delivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	r := gin.Default()
	r.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	appCtx := appctx.NewAppContext(db)

	v1 := r.Group("/v1")

	restaurants := v1.Group("/restaurants")
	restaurants.POST("/", ginrestaurant.CreateRestaurant(appCtx))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))
	restaurants.PUT("/:id", ginrestaurant.FindRestaurant(db))
	restaurants.GET("/", ginrestaurant.ListRestaurant(appCtx))
	restaurants.GET("/:id", ginrestaurant.FindRestaurant(db))

	r.Run()
}
