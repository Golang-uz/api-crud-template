package api

import (
	"github.com/gin-gonic/gin"
	"github.com/realtemirov/api-crud-template/api/handler"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Route(r *gin.Engine, h *handler.Handler) {

	r.GET("/", h.Default)
	r.GET("/ping", h.Ping)

	u := h.UserHandler
	user := r.Group("/user")
	{
		user.POST("/", u.CreateUser)
		user.GET("/id/:id", u.GetUserByID)
		user.GET("/email/:email", u.GetUserByEmail)
		user.GET("username/:username", u.GetUserByUserName)
		user.GET("/", u.GetUsers)
		user.PUT("/:id", u.UpdateUser)
		user.DELETE("/:id", u.DeleteUser)
	}

	p := h.PostHandler
	post := r.Group("post")
	{
		post.POST("/", p.CreatePost)
		post.GET("/:id", p.GetPostByID)
		post.GET("/user/:id", p.GetPostByUserID)
		post.GET("/", p.GetPosts)
		post.PUT("/:id", p.UpdatePost)
		post.DELETE("/:id", p.DeletePost)
	}
	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
