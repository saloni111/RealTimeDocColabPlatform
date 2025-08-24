package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/model"
	pb "github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/proto"
	"github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/utils"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Server struct {
	pb.UnimplementedCollaborationServiceServer
	DocumentStore *model.DocumentStore
}

func (s *Server) JoinDocument(ctx context.Context, req *pb.JoinDocumentRequest) (*pb.JoinDocumentResponse, error) {
	doc, err := s.DocumentStore.GetDocument(req.DocumentId)
	if err != nil {
		return nil, err
	}

	sessionID := uuid.New().String()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := utils.Upgrade(w, r)
		if err != nil {
			http.Error(w, "Failed to upgrade to websocket", http.StatusInternalServerError)
			return
		}
		user := &model.User{
			UserID:     req.UserId,
			Connection: conn,
		}
		doc.Mutex.Lock()
		doc.Users[req.UserId] = user
		doc.Mutex.Unlock()

		go s.handleUserConnection(doc, user)
	})

	log.Printf("User %s joined document %s", req.UserId, req.DocumentId)
	return &pb.JoinDocumentResponse{SessionId: sessionID}, nil
}

func (s *Server) SyncChanges(ctx context.Context, req *pb.SyncChangesRequest) (*pb.SyncChangesResponse, error) {
	doc, err := s.DocumentStore.GetDocument(req.DocumentId)
	if err != nil {
		return nil, err
	}

	// Accept plain text changes - no need to parse as JSON
	log.Printf("User %s syncing changes to document %s: %s", req.UserId, req.DocumentId, req.Changes)

	doc.Mutex.Lock()
	defer doc.Mutex.Unlock()

	// Update document content with the changes
	doc.Content = req.Changes
	go s.DocumentStore.UpdateDocument(req.DocumentId, req.Changes)

	// Broadcast changes to other users if they have active connections
	for _, user := range doc.Users {
		if user.UserID != req.UserId && user.Connection != nil {
			if err := user.Connection.WriteMessage(websocket.TextMessage, []byte(req.Changes)); err != nil {
				log.Printf("Failed to send message to user %s: %v", user.UserID, err)
			}
		}
	}

	log.Printf("Changes synced successfully for document %s", req.DocumentId)
	return &pb.SyncChangesResponse{Success: true}, nil
}

func (s *Server) LeaveDocument(ctx context.Context, req *pb.LeaveDocumentRequest) (*pb.LeaveDocumentResponse, error) {
	doc, err := s.DocumentStore.GetDocument(req.DocumentId)
	if err != nil {
		return nil, err
	}

	doc.Mutex.Lock()
	defer doc.Mutex.Unlock()

	if user, exists := doc.Users[req.UserId]; exists {
		user.Connection.Close()
		delete(doc.Users, req.UserId)
		log.Printf("User %s left document %s", req.UserId, req.DocumentId)
	}

	return &pb.LeaveDocumentResponse{Success: true}, nil
}

func (s *Server) handleUserConnection(doc *model.Document, user *model.User) {
	defer func() {
		doc.Mutex.Lock()
		delete(doc.Users, user.UserID)
		doc.Mutex.Unlock()
		user.Connection.Close()
		log.Printf("Connection closed for user %s", user.UserID)
	}()

	for {
		_, message, err := user.Connection.ReadMessage()
		if err != nil {
			log.Printf("Read error for user %s: %v", user.UserID, err)
			break
		}
		log.Printf("Received message from user %s: %s", user.UserID, message)
	}
}
# Updated
