package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/saloni111/RealTimeDocColabPlatform/api-gateway/utils"

	pb "github.com/saloni111/RealTimeDocColabPlatform/document-service/proto"

	"github.com/gorilla/mux"
)

func CreateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewDocumentServiceClient(conn)

	var req pb.CreateDocumentRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := client.CreateDocument(context.Background(), &req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func GetDocumentHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewDocumentServiceClient(conn)

	vars := mux.Vars(r)
	req := &pb.GetDocumentRequest{
		DocumentId: vars["document_id"],
	}

	// GET requests don't have a request body, so no need to decode JSON
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	resp, err := client.GetDocument(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func DeleteDocumentHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewDocumentServiceClient(conn)

	vars := mux.Vars(r)
	req := &pb.DeleteDocumentRequest{
		DocumentId: vars["document_id"],
	}

	// DELETE requests don't need JSON body decoding
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	resp, err := client.DeleteDocument(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func UpdateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewDocumentServiceClient(conn)

	vars := mux.Vars(r)

	// Parse JSON body for content update
	var updateRequest struct {
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.UpdateDocumentRequest{
		DocumentId: vars["document_id"],
		Content:    updateRequest.Content,
	}

	resp, err := client.UpdateDocument(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func ListDocumentVersionHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewDocumentServiceClient(conn)

	vars := mux.Vars(r)
	req := &pb.ListDocumentVersionsRequest{
		DocumentId: vars["document_id"],
	}

	// GET requests don't need JSON body decoding
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	resp, err := client.ListDocumentVersions(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
