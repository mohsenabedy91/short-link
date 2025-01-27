package routes

import (
	"short-link/internal/adapter/http/handler"
	"short-link/internal/adapter/http/middlewares"
)

// NewShortLinkRouter creates a new HTTP router
func (r *Router) NewShortLinkRouter(shortLinkHandler handler.ShortLinkHandler) *Router {
	v1 := r.Engine.Group(":language/v1", middlewares.LocaleMiddleware(r.trans))
	{
		v1.POST("create-link", shortLinkHandler.Generate)
	}
	r.Engine.GET(":shortPath", shortLinkHandler.Redirect)

	return &Router{
		Engine: r.Engine,
		conf:   r.conf,
		trans:  r.trans,
	}
}
