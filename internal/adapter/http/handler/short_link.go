package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"short-link/internal/adapter/http/presenter"
	"short-link/internal/adapter/http/requests"
	"short-link/internal/core/port"
	"short-link/pkg/translation"
)

type ShortLinkHandler struct {
	trans        translation.Translator
	shortService port.ShortLinkService
}

func NewShortLinkHandler(trans translation.Translator, shortService port.ShortLinkService) *ShortLinkHandler {
	return &ShortLinkHandler{
		trans:        trans,
		shortService: shortService,
	}
}

// Generate godoc
// @x-kong {"service": "short-link-service"}
// @Summary Generate ShortLink
// @Description Generate ShortLink
// @Tags ShortLink
// @Accept json
// @Produce json
// @Param language path string true "language 2 abbreviations" default(en)
// @Param request body requests.ShortLink true "ShortLink request"
// @Success 200 {object} presenter.Response{data=presenter.ShortLink} "Successful response"
// @Failure 400 {object} presenter.Error "Failed response"
// @Failure 422 {object} presenter.Response{validationErrors=[]presenter.ValidationError} "Validation error"
// @Failure 500 {object} presenter.Error "Internal server error"
// @ID post_language_v1_create-link
// @Router /{language}/v1/create-link [post]
func (r ShortLinkHandler) Generate(ctx *gin.Context) {
	var req requests.ShortLink
	if err := ctx.ShouldBindJSON(&req); err != nil {
		presenter.NewResponse(ctx, r.trans).Validation(err).Echo(http.StatusUnprocessableEntity)
		return
	}

	shortPath, err := r.shortService.Create(req.Url)
	if err != nil {
		presenter.NewResponse(ctx, r.trans, StatusCodeMapping).Error(err).Echo()
		return
	}
	presenter.NewResponse(ctx, r.trans).Payload(
		presenter.ToResponseShortLink(shortPath),
	).Echo()
}

func (r ShortLinkHandler) Redirect(ctx *gin.Context) {
	var req requests.Redirect
	if err := ctx.ShouldBindUri(&req); err != nil {
		presenter.NewResponse(ctx, r.trans).Validation(err).Echo(http.StatusUnprocessableEntity)
		return
	}

	url, err := r.shortService.GetByShortPath(req.ShortPath)
	if err != nil {
		presenter.NewResponse(ctx, r.trans, StatusCodeMapping).Error(err).Echo()
		return
	}

	presenter.NewResponse(ctx, r.trans).Redirect(url)
}
