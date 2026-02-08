package user

import (
	"log"
	"net/http"
	"yoharsh14/krant-backend/internal/json"
)

type Handler interface {
	CreateUser (w http.ResponseWriter,r *http.Request)
	FindUserByNameAndEmail (w http.ResponseWriter,r *http.Request)
	UpdateUser (w http.ResponseWriter,r *http.Request)
	ListAllUser (w http.ResponseWriter,r *http.Request)
}

type h struct {
	service Service
}

func NewHandler(service Service) Handler{
	return &h{
		service: service,
	}
}

func (h *h)CreateUser(w http.ResponseWriter,r *http.Request){
	var user CreateUserInput
	log.Println(r.Body)

	if err := json.Read(r,&user);err!=nil{
		log.Println("Error occured in reading User date",err)
		json.Write(w,http.StatusNotAcceptable,err)
	}
	err :=h.service.CreateUser(r.Context(),user)

	if err!=nil{
			log.Println("Error has occured", err)
		json.Write(w, http.StatusInternalServerError, err)
	}else{
		json.Write(w,http.StatusCreated,nil)
	}
}
func (h *h)FindUserByNameAndEmail (w http.ResponseWriter,r *http.Request){
	json.Write(w,200,nil)
}
func (h *h)UpdateUser (w http.ResponseWriter,r *http.Request){
	json.Write(w,200,nil)
}
func (h *h)ListAllUser (w http.ResponseWriter,r *http.Request){
	json.Write(w,200,nil)
}