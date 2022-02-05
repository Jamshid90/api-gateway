package post

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	errapi "github.com/Jamshid90/api-getawey/api/errors"
	importpb "github.com/Jamshid90/api-getawey/genproto/import"
	postpb "github.com/Jamshid90/api-getawey/genproto/post"
	"github.com/Jamshid90/api-getawey/internal/utils"
	"github.com/Jamshid90/api-getawey/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"go.uber.org/zap"
)

type HandlerOption struct {
	Logger  *zap.Logger
	Service services.Service
}

type postHandler struct {
	logger  *zap.Logger
	service services.Service
}

// new handler ...
func NewHandler(option *HandlerOption) chi.Router {
	handler := postHandler{
		logger:  option.Logger,
		service: option.Service,
	}

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Post("/", handler.create())
		r.Get("/{id}", handler.read())
		r.Put("/", handler.update())
		r.Delete("/{id}", handler.delete())
		r.Get("/", handler.list())

		r.Post("/import", handler.importPost())
	})
	return r
}

func (handler *postHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			postCreateRequest postpb.PostCreateRequest
		)

		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}
		defer r.Body.Close()

		if json.Unmarshal(requestBody, &postCreateRequest); err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}

		response, err := handler.service.PostService().Create(r.Context(), &postCreateRequest)
		if err != nil {
			render.Render(w, r, errapi.Error(err))
			return
		}

		render.JSON(w, r, response.Post)
	}
}

func (handler *postHandler) read() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}

		response, err := handler.service.PostService().Read(r.Context(), &postpb.PostReadRequest{Id: id})
		if err != nil {
			render.Render(w, r, errapi.Error(err))
			return
		}
		render.JSON(w, r, response.Post)
	}
}

func (handler *postHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			postUpdateRequest postpb.PostUpdateRequest
		)

		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}
		defer r.Body.Close()

		if json.Unmarshal(requestBody, &postUpdateRequest); err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}

		response, err := handler.service.PostService().Update(r.Context(), &postUpdateRequest)
		if err != nil {
			render.Render(w, r, errapi.Error(err))
			return
		}
		render.JSON(w, r, response.Post)
	}
}

func (handler *postHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}
		response, err := handler.service.PostService().Delete(r.Context(), &postpb.PostDeleteRequest{Id: id})
		if err != nil {
			render.Render(w, r, errapi.Error(err))
			return
		}
		render.JSON(w, r, response.Code)
	}
}

func (handler *postHandler) list() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryParameters := utils.NewQueryParameters(r.URL.Query())
		response, err := handler.service.PostService().List(r.Context(), &postpb.PostListRequest{
			Limit:  queryParameters.GetLimit(),
			Offset: queryParameters.GetOffset(),
			Filter: queryParameters.GetParameters(),
		})
		if err != nil {
			render.Render(w, r, errapi.Error(err))
			return
		}
		render.JSON(w, r, response.List)
	}
}

func (handler *postHandler) importPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			importRequest importpb.ImportRequest
		)

		requestBody, err := io.ReadAll(r.Body)
		if err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}
		defer r.Body.Close()

		if json.Unmarshal(requestBody, &importRequest); err != nil {
			render.Render(w, r, errapi.ErrInvalidArgument)
			return
		}

		response, err := handler.service.PostImportService().Import(r.Context(), &importRequest)
		if err != nil {
			render.Render(w, r, errapi.Error(err))
			return
		}

		render.JSON(w, r, map[string]int64{
			"numberOfCreated": response.NumberOfCreated,
		})
	}
}
