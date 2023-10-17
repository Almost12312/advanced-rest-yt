package author

import (
	"advanced-rest-yt/internal/apperror"
	"advanced-rest-yt/internal/http/handlers"
	"advanced-rest-yt/pkg/logging"
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

const (
	authorsUrl      = "/authors"
	authorUrl       = authorsUrl + "/:uuid"
	authorCreateUrl = authorsUrl
)

var _ handlers.Handler = &handler{}

type handler struct {
	logger     *logging.Logger
	repository Repository
}

func NewHandler(logger *logging.Logger, repository Repository) handlers.Handler {
	return &handler{logger: logger, repository: repository}
}

func (h *handler) Register(r *httprouter.Router) {
	apperror.SetMiddlewareLogger(h.logger)

	r.HandlerFunc(http.MethodGet, authorsUrl, apperror.Middleware(h.GetList))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	all, err := h.repository.FindAll(context.TODO())
	if err != nil {
		w.WriteHeader(400)
		return nil
	}

	bytes, err := json.Marshal(all)
	if err != nil {
		w.WriteHeader(500)
		return nil
	}

	w.WriteHeader(200)
	_, _ = w.Write(bytes)

	return nil
}
