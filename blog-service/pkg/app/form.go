package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		for _, realErr := range verrs {
			// 注意 ，这里没有使用国际化中间件
			errs = append(errs, &ValidError{
				Key:     realErr.Namespace(),
				Message: realErr.Error(),
			})
		}
		return false, errs
	}
	return true, nil
}
