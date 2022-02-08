package main

import (
	"demo/component"
	"demo/modules/restaurant/restauranttransport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name;"`
	Addr string `json:"addr" gorm:"column:addr"`
}
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name;"`
	Addr *string `json:"address" gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := os.Getenv("DBConnectionStr")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	if err := runService(db); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB) error {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	appCtx := component.NewAppContext(db)

	restaurants := r.Group("/restaurants")
	{
		restaurants.POST("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))

		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))

		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))

		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	return r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
