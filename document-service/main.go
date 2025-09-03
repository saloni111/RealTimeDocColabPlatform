package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/saloni111/RealTimeDocColabPlatform/document-service/handler"
	"github.com/saloni111/RealTimeDocColabPlatform/document-service/model"
	pb "github.com/saloni111/RealTimeDocColabPlatform/document-service/proto"
	"github.com/saloni111/RealTimeDocColabPlatform/document-service/utils"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	dynamodb, err := utils.DynamoDBInstance()

	if err != nil {
		log.Fatalf("failed to establish connection to db: %v", err)
	}

	pb.RegisterDocumentServiceServer(s, &handler.Server{DocumentModel: &model.DocumentModel{
		DynamoDB:  dynamodb,
		TableName: "docs",
	}})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
