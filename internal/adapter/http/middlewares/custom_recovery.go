package middlewares

import (
	"github.com/gin-gonic/gin"
	"short-link/internal/adapter/http/presenter"
	"short-link/pkg/serviceerror"
	"short-link/pkg/translation"
)

func ErrorHandler(trans translation.Translator) func(ctx *gin.Context, err interface{}) {
	return func(ctx *gin.Context, err interface{}) {
		serviceErr := serviceerror.NewServerError()
		presenter.NewResponse(ctx, trans).Error(serviceErr).Echo()
	}
}
