package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/handler"
	"github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/model"
	pb "github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/proto"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	documentStore := model.NewDocumentStore()

	server := &handler.Server{
		DocumentStore: documentStore,
	}

	pb.RegisterCollaborationServiceServer(s, server)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
# Updated
