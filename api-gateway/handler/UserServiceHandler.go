package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/saloni111/RealTimeDocColabPlatform/api-gateway/utils"

	pb "github.com/saloni111/RealTimeDocColabPlatform/user-service/proto"
)

func RegisterUserHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50051")
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	var req pb.RegisterUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := client.RegisterUser(context.Background(), &req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50051")
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	var req pb.LoginUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := client.LoginUser(context.Background(), &req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {
	conn := utils.GetGRPCConnection("localhost:50051")
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	var req pb.GetUserProfileRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := client.GetUserProfile(context.Background(), &req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}
