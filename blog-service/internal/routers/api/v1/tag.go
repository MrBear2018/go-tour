package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-tour/blog-service/global"
	"github.com/go-tour/blog-service/internal/service"
	"github.com/go-tour/blog-service/pkg/app"
	"github.com/go-tour/blog-service/pkg/error_code"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) List(c *gin.Context) {
	param := service.CountTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		response.ToErrorResponse(error_code.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	response.ToResponse(gin.H{})
	return
}
func (t Tag) Create(c *gin.Context) {}
func (t Tag) Update(c *gin.Context) {}
func (t Tag) Delete(c *gin.Context) {}
