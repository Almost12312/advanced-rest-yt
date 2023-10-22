package author

import (
	"advanced-rest-yt/internal/apperror"
	"advanced-rest-yt/internal/http/handlers"
	"advanced-rest-yt/pkg/logging"
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
	logger  *logging.Logger
	service *Service
}

func NewHandler(logger *logging.Logger, service *Service) handlers.Handler {
	return &handler{logger: logger, service: service}
}

func (h *handler) Register(r *httprouter.Router) {
	apperror.SetMiddlewareLogger(h.logger)

	r.HandlerFunc(http.MethodGet, authorsUrl, apperror.Middleware(h.GetList))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	all, err := h.service.GetAll(r.Context())
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
