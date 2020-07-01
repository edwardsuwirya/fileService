package main

import (
	"fmt"
	"github.com/edwardsuwirya/fileService/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()
	arg := os.Args
	fileRoute := r.PathPrefix("/upload").Subrouter()
	fileRoute.HandleFunc("", handler.NewFileUploadHandler().Handler).Methods(http.MethodPost)

	log.Printf("Server is listening on %s port %s", arg[1], arg[2])
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", arg[1], arg[2]), r); err != nil {
		log.Panic(err)
	}
}
