package user

import (
	"advanced-rest-yt/internal/apperror"
	"advanced-rest-yt/internal/http/handlers"
	"advanced-rest-yt/pkg/logging"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	usersUrl      = "/users"
	userUrl       = usersUrl + "/:uuid"
	userCreateUrl = usersUrl
)

var _ handlers.Handler = &handler{}

type handler struct {
	logger *logging.Logger
}

func NewHandler(logger *logging.Logger) handlers.Handler {
	return &handler{logger: logger}
}

func (h *handler) Register(r *httprouter.Router) {
	apperror.SetMiddlewareLogger(h.logger)

	r.HandlerFunc(http.MethodGet, usersUrl, apperror.Middleware(h.GetList))
	r.HandlerFunc(http.MethodGet, userUrl, apperror.Middleware(h.UserByUuid))

	r.HandlerFunc(http.MethodPost, userCreateUrl, apperror.Middleware(h.Create))

	r.HandlerFunc(http.MethodPut, userUrl, apperror.Middleware(h.Update))
	r.HandlerFunc(http.MethodPatch, userUrl, apperror.Middleware(h.ParticiallUpdate))

	r.HandlerFunc(http.MethodDelete, userUrl, apperror.Middleware(h.Delete))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	return apperror.ErrNotFound
}

func (h *handler) UserByUuid(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("api err")
}

func (h *handler) Create(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(201)
	w.Write([]byte("this is Create"))
	return nil

}

func (h *handler) Update(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is Update"))
	return nil

}

func (h *handler) ParticiallUpdate(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(204)
	w.Write([]byte("this is ParticiallUpdate"))
	return nil

}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request) error {
	return apperror.NewAppError(nil, "test", "test", "US-000005")
}
