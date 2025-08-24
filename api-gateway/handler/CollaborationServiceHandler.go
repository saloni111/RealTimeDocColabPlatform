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
	conn := utils.GetGRPCConnection("localhost:50053")
	defer conn.Close()

	client := pb.NewCollaborationServiceClient(conn)

	vars := mux.Vars(r)

	var requestBody struct {
		UserID     string `json:"user_id"`
		DocumentID string `json:"document_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.JoinDocumentRequest{
		DocumentId: vars["document_id"],
		UserId:     requestBody.UserID,
	}

	resp, err := client.JoinDocument(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func SyncChangesHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50053")
	defer conn.Close()

	client := pb.NewCollaborationServiceClient(conn)

	vars := mux.Vars(r)

	var requestBody struct {
		SessionID  string `json:"session_id"`
		DocumentID string `json:"document_id"`
		UserID     string `json:"user_id"`
		Changes    string `json:"changes"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.SyncChangesRequest{
		SessionId:  requestBody.SessionID,
		DocumentId: vars["document_id"],
		UserId:     requestBody.UserID,
		Changes:    requestBody.Changes,
	}

	resp, err := client.SyncChanges(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func LeaveDocumentHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50053")
	defer conn.Close()

	client := pb.NewCollaborationServiceClient(conn)

	vars := mux.Vars(r)

	var requestBody struct {
		SessionID  string `json:"session_id"`
		DocumentID string `json:"document_id"`
		UserID     string `json:"user_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	req := &pb.LeaveDocumentRequest{
		SessionId:  requestBody.SessionID,
		DocumentId: vars["document_id"],
		UserId:     requestBody.UserID,
	}

	resp, err := client.LeaveDocument(context.Background(), req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
