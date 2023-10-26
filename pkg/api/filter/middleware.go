package filter

import (
	"context"
	"net/http"
)

const (
	ASC               = "ASC"
	DESC              = "DESC"
	OptionsContextKey = "sort_options"
)

func Middleware(handlerFunc http.HandlerFunc, defaultLimitField int) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		field := request.URL.Query().Get("limit")

		if field == "" {
			field = defaultLimitField
		}

		opts := Options{
			Field: field,
			Order: order,
		}

		optCtx := context.WithValue(request.Context(), OptionsContextKey, opts)
		request = request.WithContext(optCtx)

		handlerFunc(writer, request)
	}
}

type Options struct {
	Field, Order string
}
