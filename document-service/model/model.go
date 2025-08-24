package model

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DocumentModel struct {
	DynamoDB  *dynamodb.Client
	TableName string
}

type Document struct {
	DocumentID string   `dynamodbav:"document_id"`
	Title      string   `dynamodbav:"title"`
	Content    string   `dynamodbav:"content"`
	Author     string   `dynamodbav:"author"`
	Versions   []string `dynamodbav:"versions"`
	Timestamp  string   `dynamodbav:"timestamp"`
}

func (d *DocumentModel) CreateDocument(ctx context.Context, document *Document) error {
	av, err := attributevalue.MarshalMap(document)

	if err != nil {
		return fmt.Errorf("failed to create new document %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &d.TableName,
	}

	_, err = d.DynamoDB.PutItem(ctx, input)

	if err != nil {
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}

func (d *DocumentModel) GetDocumentByID(ctx context.Context, documentId string) (*Document, error) {
	input := &dynamodb.GetItemInput{
		TableName: &d.TableName,
		Key: map[string]types.AttributeValue{
			"document_id": &types.AttributeValueMemberS{Value: documentId},
		},
	}

	result, err := d.DynamoDB.GetItem(ctx, input)

	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	// Check if item exists
	if result.Item == nil {
		return nil, fmt.Errorf("document not found")
	}

	var document Document

	err = attributevalue.UnmarshalMap(result.Item, &document)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal document: %w", err)
	}

	return &document, nil
}

func (d *DocumentModel) DeleteDocumentByID(ctx context.Context, documentId string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: &d.TableName,
		Key: map[string]types.AttributeValue{
			"document_id": &types.AttributeValueMemberS{Value: documentId},
		},
	}

	_, err := d.DynamoDB.DeleteItem(ctx, input)

	if err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}

	return nil
}

func (d *DocumentModel) UpdateDocumentByID(ctx context.Context, documentId string, content string) error {
	timestamp := time.Now().Format(time.RFC3339)

	// Use UpdateExpression to avoid overwriting other fields
	input := &dynamodb.UpdateItemInput{
		TableName: &d.TableName,
		Key: map[string]types.AttributeValue{
			"document_id": &types.AttributeValueMemberS{Value: documentId},
		},
		UpdateExpression: "SET content = :content, #ts = :timestamp ADD versions :version",
		ExpressionAttributeNames: map[string]string{
			"#ts": "timestamp",
		},
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":content":   &types.AttributeValueMemberS{Value: content},
			":timestamp": &types.AttributeValueMemberS{Value: timestamp},
			":version":   &types.AttributeValueMemberL{Value: []types.AttributeValue{&types.AttributeValueMemberS{Value: timestamp}}},
		},
	}

	_, err := d.DynamoDB.UpdateItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}
	return nil
}

func (d *DocumentModel) ListDocumentVersions(ctx context.Context, documentId string) ([]string, error) {
	doc, err := d.GetDocumentByID(ctx, documentId)

	if err != nil {
		return nil, err
	}

	return doc.Versions, nil
}
