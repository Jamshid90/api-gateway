package api

import (
	"context"
	"net/http"
	"time"

	"github.com/Jamshid90/api-getawey/api/handlers/v1/post"
	"github.com/Jamshid90/api-getawey/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type Option struct {
	Logger         *zap.Logger
	Service        services.Service
	ContextTimeout time.Duration
}

//New ...
func New(option Option) http.Handler {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(option.ContextTimeout))
	//r.Use(contentTypeJson)

	r.Route("/v1", func(r chi.Router) {
		r.Use(apiVersionCtx("v1"))

		// post handlers initialization
		r.Mount("/post", post.NewHandler(&post.HandlerOption{Logger: option.Logger, Service: option.Service}))

	})

	return r
}

func apiVersionCtx(version string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(context.WithValue(r.Context(), "api.version", version))
			next.ServeHTTP(w, r)
		})
	}
}

func contentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
