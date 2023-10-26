package author

import (
	"advanced-rest-yt/internal/apperror"
	service2 "advanced-rest-yt/internal/author/service"
	"advanced-rest-yt/internal/http/handlers"
	"advanced-rest-yt/pkg/api/filter"
	"advanced-rest-yt/pkg/api/sort"
	"advanced-rest-yt/pkg/logging"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
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

	router.HandlerFunc(http.MethodGet, authorsUrl, filter.Middleware(sort.Middleware(apperror.Middleware(h.GetList), "created_at", sort.ASC), 10))
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) error {
	filterOptions := r.Context().Value(filter.OptionsContextKey).(filter.Options)

	name := r.URL.Query().Get("name")
	if name != "" {
		err := filterOptions.AddField("name", filter.OperatorLike, name, filter.DataTypeStr)
		if err != nil {
			return err
		}
	}

	age := r.URL.Query().Get("age")
	if age != "" {
		operator := filter.OperatorEq
		v := age
		if index := strings.Index(age, ":"); index != -1 {
			split := strings.Split(age, ":")
			operator = split[0]
			v = split[1]
		}
		err := filterOptions.AddField("age", operator, v, filter.DataTypeInt)
		if err != nil {
			return err
		}
	}

	isAlive := r.URL.Query().Get("is_alive")
	if isAlive != "" {
		_, err := strconv.ParseBool(isAlive)
		if err != nil {
			bad := apperror.BadRequest("filter params incorrect", "bool value wrong")
			bad.WithFields(map[string]string{
				"is_alive": "this value must be true or false",
			})
		}

		err = filterOptions.AddField("is_alive", filter.OperatorEq, isAlive, filter.DataTypeBool)
		if err != nil {
			return err
		}
	}

	createdAt := r.URL.Query().Get("created_at")
	if createdAt != "" {
		var operator string
		if ix := strings.Index(createdAt, ":"); ix != 1 {
			// range
			operator = filter.OperatorEq
		} else {
			// single date
			operator = filter.OperatorBetween
		}

		err := filterOptions.AddField("is_alive", operator, createdAt, filter.DataTypeDate)
		if err != nil {
			return err
		}
	}

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
