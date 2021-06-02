package main

import (
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
	log.Println(db, err)
	r := gin.Default()
	r.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	restaurants := v1.Group("/restaurants")

	restaurants.POST("/", ginrestaurant.CreateRestaurant(db))
	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(db))
	restaurants.GET("/", ginrestaurant.FindRestaurant(db))

	r.Run()
}
