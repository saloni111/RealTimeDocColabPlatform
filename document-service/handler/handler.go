package handler

import (
	"context"
	"time"

	"github.com/saloni111/RealTimeDocColabPlatform/document-service/model"
	pb "github.com/saloni111/RealTimeDocColabPlatform/document-service/proto"

	"github.com/google/uuid"
)

type Server struct {
	pb.UnimplementedDocumentServiceServer
	DocumentModel *model.DocumentModel
}

func (s *Server) CreateDocument(ctx context.Context, req *pb.CreateDocumentRequest) (*pb.CreateDocumentResponse, error) {
	documentID := uuid.New().String()
	creationTimestamp := time.Now().Format(time.RFC3339)

	document := model.Document{
		Title:      req.Title,
		Author:     req.Author,
		Content:    req.Content, // Fixed: Use actual content instead of " "
		DocumentID: documentID,
		Timestamp:  creationTimestamp,
		Versions:   []string{creationTimestamp},
	}

	err := s.DocumentModel.CreateDocument(ctx, &document)

	if err != nil {
		return nil, err
	}

	return &pb.CreateDocumentResponse{
		DocumentId: documentID,
	}, nil
}

func (s *Server) GetDocument(ctx context.Context, req *pb.GetDocumentRequest) (*pb.GetDocumentResponse, error) {
	document, err := s.DocumentModel.GetDocumentByID(ctx, req.DocumentId)

	if err != nil {
		return nil, err
	}

	return &pb.GetDocumentResponse{
		DocumentId: document.DocumentID,
		Title:      document.Title,
		Content:    document.Content,
		Author:     document.Author,
		Versions:   document.Versions,
	}, nil
}

func (s *Server) DeleteDocument(ctx context.Context, req *pb.DeleteDocumentRequest) (*pb.DeleteDocumentResponse, error) {
	err := s.DocumentModel.DeleteDocumentByID(ctx, req.DocumentId)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteDocumentResponse{
		DocumentId: req.DocumentId,
	}, nil
}

func (s *Server) UpdateDocument(ctx context.Context, req *pb.UpdateDocumentRequest) (*pb.UpdateDocumentResponse, error) {
	err := s.DocumentModel.UpdateDocumentByID(ctx, req.DocumentId, req.Content)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateDocumentResponse{
		DocumentId: req.DocumentId,
	}, nil
}

func (s *Server) ListDocumentVersions(ctx context.Context, req *pb.ListDocumentVersionsRequest) (*pb.ListDocumentVersionsResponse, error) {
	versions, err := s.DocumentModel.ListDocumentVersions(ctx, req.DocumentId)

	if err != nil {
		return nil, err
	}

	return &pb.ListDocumentVersionsResponse{
		Versions: versions,
	}, nil
}
