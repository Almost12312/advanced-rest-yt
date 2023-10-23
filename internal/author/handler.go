package author

import (
	"advanced-rest-yt/internal/apperror"
	service2 "advanced-rest-yt/internal/author/service"
	"advanced-rest-yt/internal/http/handlers"
	"advanced-rest-yt/pkg/api/sort"
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
	service *service2.Service
}

func NewHandler(logger *logging.Logger, service *service2.Service) handlers.Handler {
	return &handler{logger: logger, service: service}
}

func (h *handler) Register(router *httprouter.Router) {
	apperror.SetMiddlewareLogger(h.logger)

	router.HandlerFunc(http.MethodGet, authorsUrl, sort.Middleware(apperror.Middleware(h.GetList), "created_at", sort.ASC))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	var sortOptions sort.Options
	if options, ok := r.Context().Value(sort.OptionsContextKey).(sort.Options); ok {
		sortOptions = options
	}

	all, err := h.service.GetAll(r.Context(), sortOptions)
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
