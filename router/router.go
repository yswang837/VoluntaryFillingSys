package routes

import (
	"github.com/VoluntaryFillingSys/service/user"
	"github.com/VoluntaryFillingSys/test"
	"github.com/VoluntaryFillingSys/utils"
	"github.com/gin-gonic/gin"
)

var DefaultRouter Router

type Router struct {
	userConsumer *user.Consumer
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r := gin.New()
	//r.Use(gin.Recovery(), middleware.Logger(), middleware.Cors())
	//auth := r.Group("api/v1")
	//auth.Use(middleware.JwtToken())
	//{
	//	// 用户模块的路由接口
	//
	//	auth.PUT("user/:id", v1.EditUser)
	//	auth.DELETE("user/:id", v1.DeleteUser)
	//	// 分类模块的路由接口
	//	auth.POST("category/add", v1.AddCate)
	//
	//	auth.PUT("category/:id", v1.EditCate)
	//	auth.DELETE("category/:id", v1.DeleteCate)
	//	// 文章模块的路由接口
	//	auth.POST("article/add", v1.AddArt)
	//
	//	auth.PUT("article/:id", v1.EditArt)
	//	auth.DELETE("article/:id", v1.DeleteArt)
	//	//文件上传
	//	auth.POST("upload", v1.Upload)
	//}
	router := r.Group("/api/v1")
	//router.GET("users", v1.GetUsers)
	//router.GET("category", v1.GetCates)
	//router.GET("article", v1.GetArts)
	//router.GET("article/list/:id", v1.GetCateArt)
	//router.GET("article/info/:id", v1.GetArtInfo)
	//router.POST("login", v1.Login)
	router.POST("user/add", DefaultRouter.userConsumer.AddUser)
	router.GET("ping", test.Pong)
	return r
}

func StartGinService() {
	r := InitRouter()
	if err := r.Run(utils.HttpPort); err != nil {
		panic(err)
	}
}

func init() {
	DefaultRouter.userConsumer = user.NewConsumer()
}
