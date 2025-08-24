package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Response structure for listing documents
type DocumentSummary struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Modified string `json:"modified"`
}

type ListDocumentsResponse struct {
	Documents []DocumentSummary `json:"documents"`
	Total     int               `json:"total"`
}

func GetDynamoDBClient() *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:9000"}, nil
			},
		)),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     "dummy",
				SecretAccessKey: "dummy",
				SessionToken:    "",
			},
		}),
	)
	if err != nil {
		panic(err)
	}
	return dynamodb.NewFromConfig(cfg)
}

func ListDocumentsHandler(w http.ResponseWriter, r *http.Request) {
	// Create a simple scan to get all documents from DynamoDB
	// For now, we'll make a direct DynamoDB call since the document service doesn't have ListDocuments gRPC method

	dynamoClient := GetDynamoDBClient()

	// Scan all documents from the table
	scanInput := &dynamodb.ScanInput{
		TableName: aws.String("docs"),
	}

	result, err := dynamoClient.Scan(context.Background(), scanInput)
	if err != nil {
		http.Error(w, "Failed to fetch documents: "+err.Error(), http.StatusInternalServerError)
		return
	}

	var documents []DocumentSummary

	for _, item := range result.Items {
		doc := DocumentSummary{}

		if id, ok := item["document_id"].(*types.AttributeValueMemberS); ok {
			doc.ID = id.Value
		}
		if title, ok := item["title"].(*types.AttributeValueMemberS); ok {
			doc.Title = title.Value
		}
		if author, ok := item["author"].(*types.AttributeValueMemberS); ok {
			doc.Author = author.Value
		}
		if content, ok := item["content"].(*types.AttributeValueMemberS); ok {
			doc.Content = content.Value
		}
		if timestamp, ok := item["timestamp"].(*types.AttributeValueMemberS); ok {
			doc.Modified = timestamp.Value
		}

		// Only add documents that have at least an ID
		if doc.ID != "" {
			documents = append(documents, doc)
		}
	}

	w.Header().Set("Content-Type", "application/json")

	response := ListDocumentsResponse{
		Documents: documents,
		Total:     len(documents),
	}

	json.NewEncoder(w).Encode(response)
}
