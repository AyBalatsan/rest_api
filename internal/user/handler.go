package user

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"rest_api/internal/hendlers"
	"rest_api/pkg/logging"
)

// Подсказка
var _ hendlers.Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/:uuid"
)

type handler struct {
	logger *logging.Logger
}

// отдаем ссылку на структуру в интерфейсе
func NewHandler(logger *logging.Logger) hendlers.Handler {
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(usersURL, h.GetList)
	router.GET(userURL, h.GetUserByUUID)
	router.POST(usersURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)

}

func (h *handler) GetList(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	h.logger.Info("GetList")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is list of users"))
}
func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("New user"))
}

func (h *handler) GetUserByUUID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is user"))
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("UpdateUser"))
}
func (h *handler) PartiallyUpdateUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("PartiallyUpdateUser"))
}
func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("delete user by id"))
}
