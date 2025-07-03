package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/saloni111/RealTimeDocColabPlatform/api-gateway/utils"

	pb "github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/proto"

	"github.com/gorilla/mux"
)

func JoinDocumentHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewCollaborationServiceClient(conn)

	vars := mux.Vars(r)
	req := &pb.JoinDocumentRequest{
		DocumentId: vars["document_id"],
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := client.JoinDocument(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func SyncChangesHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewCollaborationServiceClient(conn)

	vars := mux.Vars(r)
	req := &pb.SyncChangesRequest{
		DocumentId: vars["document_id"],
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := client.SyncChanges(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func LeaveDocumentHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50052")
	defer conn.Close()

	client := pb.NewCollaborationServiceClient(conn)

	vars := mux.Vars(r)
	req := &pb.LeaveDocumentRequest{
		DocumentId: vars["document_id"],
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := client.LeaveDocument(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
