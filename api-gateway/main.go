package main

import (
	"log"
	"net/http"

	handler "github.com/saloni111/RealTimeDocColabPlatform/api-gateway/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/user/register", handler.RegisterUserHandler).Methods("POST")
	r.HandleFunc("/login", handler.LoginUserHandler).Methods("POST")
	r.HandleFunc("/user", handler.GetUserProfileHandler).Methods("GET")

	r.HandleFunc("/document/create", handler.CreateDocumentHandler).Methods("POST")
	r.HandleFunc("/document/{document_id}", handler.GetDocumentHandler).Methods("GET")
	r.HandleFunc("/document/{document_id}", handler.DeleteDocumentHandler).Methods("DELETE")
	r.HandleFunc("/document/{document_id}", handler.UpdateDocumentHandler).Methods("PUT")
	r.HandleFunc("/document/{document_id}/version", handler.ListDocumentVersionHandler).Methods("GET")

	r.HandleFunc("/document/join/{document_id}", handler.CreateDocumentHandler).Methods("POST")
	r.HandleFunc("/document/sync/{document_id}", handler.GetDocumentHandler).Methods("POST")
	r.HandleFunc("/document/leave/{document_id}", handler.DeleteDocumentHandler).Methods("POST")

	log.Printf("API Gateway listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
