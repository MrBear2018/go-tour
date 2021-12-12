package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-tour/blog-service/internal/routers/api/v1"
)

func NewCusRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// TODO 这里目前是人工维护，后续可以改成放在某一个结构体中，便于统一管理；
	// TODO 改为依赖注入框架，避免人工维护操作
	article := v1.NewArticle()
	tag := v1.NewTag()

	apiV1 := r.Group("/api/v1")
	{
		// 标签管理
		apiV1.POST("/tags", tag.Create)            // 批量新增 tag
		apiV1.DELETE("/tags/:id", tag.Delete)      // 根据 id 删除指定的 tag
		apiV1.PUT("/tags/:id", tag.Update)         // 根据 id 更新指定的标签
		apiV1.PATCH("/tags/:id/state", tag.Update) // 根据 id 修改指定标签的状态
		apiV1.GET("/tags", tag.List)               // 获取标签列表

		// 文章管理
		apiV1.POST("/articles", article.Create)            // 新增文章
		apiV1.DELETE("/articles/:id", article.Delete)      // 根据 id 删除指定的文章
		apiV1.PUT("/articles/:id", article.Update)         // 更新文章
		apiV1.PATCH("/articles/:id/state", article.Update) // 根据 id 更新文章的状态
		apiV1.GET("/articles/:id", article.Get)            // 根据 id 获取文章
		apiV1.GET("/articles", article.List)               // 获取文章列表
	}

	return r
}
