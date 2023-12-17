package routes

import (
	"github.com/Chandan050/blogwebsite/controller"
	"github.com/Chandan050/blogwebsite/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {
	app.POST("/app/register", controller.Register())
	app.POST("/app/login", controller.Login())
	app.Use(middleware.IsAunthenticate())
	app.POST("/app/post", controller.CreatePost())
	app.GET("/app/allpost", controller.AllPost())
	app.GET("/app/allpost/:id", controller.DetailsPost())
	app.PUT("/app/updatepost/:id", controller.UpdatePost())
	app.GET("/app/uniquepost", controller.UniuePost())
	app.DELETE("/app/deletepost/:id", controller.DeletePost())
	app.POST("app/upload-image", controller.Upload())

}
