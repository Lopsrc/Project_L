
package user

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.mod/internal/handlers"
	"go.mod/pkg/logging"
)
const(
	usersURL = "/users"
	userURL = "/user/:uuid"
)
type handler struct{
	logger logging.Logger
}

func NewHandler(logger logging.Logger) handlers.Handler{
	return &handler{
		logger: logger,
	}
}

func (h *handler) Register(router *httprouter.Router){
	router.GET(usersURL, h.Getlist)
	router.POST(usersURL, h.GetUserByUID)
	router.GET(userURL, h.CreateUser)
	router.PUT(userURL, h.UpdateUser)
	router.PATCH(userURL, h.PartiallyUpdateUser)
	router.DELETE(userURL, h.DeleteUser)
}
func (h *handler) Getlist(w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("this is list of users"))
}

func (h *handler)GetUserByUID (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("this is user by uid"))
}
func (h *handler)CreateUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("this is create  user"))
}
func (h *handler)UpdateUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("this is update user"))
}
func (h *handler)PartiallyUpdateUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("this is partially update user"))
}
func (h *handler)DeleteUser (w http.ResponseWriter, r *http.Request, params httprouter.Params){
	w.Write([]byte("this is delete user"))
}