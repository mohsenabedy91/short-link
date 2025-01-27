package middlewares

import (
	"github.com/gin-gonic/gin"
	"short-link/internal/adapter/http/presenter"
	"short-link/pkg/serviceerror"
	"short-link/pkg/translation"
)

type LanguageUri struct {
	Language string `uri:"language" binding:"required"`
}

func LocaleMiddleware(trans translation.Translator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var langUri LanguageUri
		if err := ctx.ShouldBindUri(&langUri); err != nil {
			serviceErr := serviceerror.NewServerError()
			presenter.NewResponse(ctx, trans).Error(serviceErr).Echo()
		}
		_ = trans.GetLocalizer(langUri.Language)

		ctx.Next()
	}
}
