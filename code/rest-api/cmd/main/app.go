package main

import (
	//"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.mod/internal/user"
	"go.mod/pkg/logging"
)

// func IndexHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params){
// 	name := params.ByName("name")
// 	w.Write([]byte(fmt.Sprintf("Hello %s", name)))
// }

func main(){
	logger := logging.GetLogger() 
	logger.Info("Create router")
	router := httprouter.New() 
	logger.Info("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(router)
	start(router)

}
func start(router *httprouter.Router){
	log.Println("start application")

	listener, err := net.Listen("tcp", "0.0.0.0:1234")
	if err != nil{
		panic(err)
	}

	server := &http.Server{
		Handler: router,
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}
	log.Println("server is listening port 0.0.0.0:1234")
	log.Fatal(server.Serve(listener))
	
}