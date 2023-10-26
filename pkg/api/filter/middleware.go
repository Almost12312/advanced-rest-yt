package filter

import (
	"context"
	"net/http"
	"strconv"
)

const (
	OptionsContextKey = "filter_options"
)

func Middleware(handlerFunc http.HandlerFunc, defaultLimitField int) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		limitQuery := request.URL.Query().Get("limit")

		limit := defaultLimitField
		var limitErr error

		if limitQuery != "" {
			if limit, limitErr = strconv.Atoi(limitQuery); limitErr != nil {
				writer.WriteHeader(http.StatusBadRequest)
				writer.Write([]byte("bad request"))
				return
			}
		}

		opts := NewOption(limit)

		optCtx := context.WithValue(request.Context(), OptionsContextKey, opts)
		request = request.WithContext(optCtx)

		handlerFunc(writer, request)
	}
}
