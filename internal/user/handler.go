package user

import (
	"advanced-rest-yt/internal/http/handlers"
	"advanced-rest-yt/pkg/logging"
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
	r.HandlerFunc("GET", usersUrl, h.GetList)
	r.GET(userUrl, h.UserByUuid)

	r.POST(userCreateUrl, h.Create)

	r.PUT(userUrl, h.Update)
	r.PATCH(userUrl, h.ParticiallUpdate)

	r.DELETE(userUrl, h.Delete)
}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("nothing"))
	w.WriteHeader(200)
}

func (h *handler) UserByUuid(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(200)
	w.Write([]byte("this is UserByUuid"))

}

func (h *handler) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(201)
	w.Write([]byte("this is Create"))
}

func (h *handler) Update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is Update"))
}

func (h *handler) ParticiallUpdate(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is ParticiallUpdate"))
}

func (h *handler) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(204)
	w.Write([]byte("this is Delete"))
}
