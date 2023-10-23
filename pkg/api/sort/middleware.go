package sort

import (
	"context"
	"net/http"
	"strings"
)

const (
	ASC               = "ASC"
	DESC              = "DESC"
	OptionsContextKey = "sort_options"
)

func Middleware(handlerFunc http.HandlerFunc, defaultSortField, defaultSortOrder string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		field := request.URL.Query().Get("sort_by")
		order := request.URL.Query().Get("sort_order")

		if field == "" {
			field = defaultSortField
		}

		if order == "" {
			order = defaultSortOrder
		} else {
			u := strings.ToUpper(order)
			if u != ASC && u != DESC {
				writer.WriteHeader(http.StatusBadRequest)
				// TODO: app error
				return
			}
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
