package main

import (
	"demo/common"
	"demo/component"
	"demo/component/uploadprovider"
	"demo/middleware"
	"demo/modules/restaurant/restauranttransport/ginrestaurant"
	"demo/modules/restaurantlike/transport/ginresturantlike"
	"demo/modules/upload/uploadtransport/ginupload"
	"demo/modules/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	dsn := os.Getenv("DBConnectionStr")

	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3APIKey := os.Getenv("S3APIKey")
	s3SecretKey := os.Getenv("S3SecretKey")
	s3Domain := os.Getenv("S3Domain")
	secretKey := os.Getenv("SYSTEM_SECRET")

	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db = db.Debug()

	if err := runService(db, s3Provider, secretKey); err != nil {
		log.Fatalln(err)
	}
}

func runService(db *gorm.DB, upProvider uploadprovider.UploadProvider, secretKey string) error {

	appCtx := component.NewAppContext(db, upProvider, secretKey)
	r := gin.Default()

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// CRUD

	v1 := r.Group("/v1")

	v1.POST("/upload", ginupload.Upload(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))
	v1.POST("/login", ginuser.Login(appCtx))
	v1.GET("/profile", middleware.RequiredAuth(appCtx), ginuser.GetProfile(appCtx))

	restaurants := v1.Group("/restaurants", middleware.RequiredAuth(appCtx))
	{
		restaurants.POST("", ginrestaurant.CreateRestaurant(appCtx))
		restaurants.GET("/:id", ginrestaurant.GetRestaurant(appCtx))
		restaurants.GET("", ginrestaurant.ListRestaurant(appCtx))
		restaurants.PATCH("/:id", ginrestaurant.UpdateRestaurant(appCtx))
		restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))

		restaurants.GET("/:id/liked-users", ginresturantlike.ListUser(appCtx))
		restaurants.POST("/:id/like", ginresturantlike.UserLikeRestaurant(appCtx))
		restaurants.DELETE("/:id/unlike", ginresturantlike.UserUnlikeRestaurant(appCtx))
	}

	v1.GET("/encode-uid", func(c *gin.Context) {
		type reqData struct {
			DbType int `form:"type"`
			RealId int `form:"id"`
		}

		var d reqData
		c.ShouldBind(&d)

		c.JSON(http.StatusOK, gin.H{
			"id": common.NewUID(uint32(d.RealId), d.DbType, 1),
		})
	})

	return r.Run()
}
