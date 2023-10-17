package apperror

import (
	"advanced-rest-yt/pkg/logging"
	"errors"
	"net/http"
)

var logger *logging.Logger

type appHandler func(w http.ResponseWriter, r *http.Request) error

func Middleware(h appHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var appErr *AppError

		err := h(w, r)
		if err != nil {
			if errors.As(err, &appErr) {
				if errors.Is(err, ErrNotFound) {
					logger.Error(ErrNotFound.Msg)

					w.WriteHeader(http.StatusNotFound)
					w.Write(ErrNotFound.Marshal())

					return
				}
				err = err.(*AppError)

				logger.Error(appErr.Error())
				w.WriteHeader(http.StatusBadRequest)
				w.Write((appErr.Marshal()))
				return
			}

			e := systemError(err)
			logger.Error(e.Unwrap())
			w.WriteHeader(http.StatusTeapot)
			w.Write(e.Marshal())
			return
		}

	}
}

func SetMiddlewareLogger(log *logging.Logger) {
	logger = log
}
