package model

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/saloni111/RealTimeDocColabPlatform/collaboration-service/utils"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gorilla/websocket"
)

type Document struct {
	DocumentID string
	Content    string
	Users      map[string]*User
	Mutex      sync.Mutex
}

type User struct {
	UserID     string
	Connection *websocket.Conn
}

type DocumentStore struct {
	Documents map[string]*Document
	Mutex     sync.Mutex
	DynamoDB  *dynamodb.Client
	TableName string
}

func NewDocumentStore() *DocumentStore {
	dynamodb, err := utils.DynamoDBInstance()

	if err != nil {
		log.Fatalf("failed to establish connection to db: %v", err)
	}

	return &DocumentStore{
		Documents: make(map[string]*Document),
		DynamoDB:  dynamodb,
		TableName: "docs",
	}
}

func (store *DocumentStore) GetDocument(documentID string) (*Document, error) {
	store.Mutex.Lock()
	defer store.Mutex.Unlock()

	if doc, exists := store.Documents[documentID]; exists {
		return doc, nil
	}

	input := &dynamodb.GetItemInput{
		TableName: &store.TableName,
		Key: map[string]types.AttributeValue{
			"document_id": &types.AttributeValueMemberS{Value: documentID},
		},
	}

	result, err := store.DynamoDB.GetItem(context.Background(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to get item from DynamoDB: %w", err)
	}

	var doc Document
	err = attributevalue.UnmarshalMap(result.Item, &doc)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal document: %w", err)
	}
	doc.Users = make(map[string]*User)

	store.Documents[documentID] = &doc

	return &doc, nil
}

func (store *DocumentStore) UpdateDocument(documentID string, changes string) error {
	timestamp := time.Now().Format(time.RFC3339)

	update := map[string]types.AttributeValueUpdate{
		"content": {
			Action: types.AttributeActionPut,
			Value:  &types.AttributeValueMemberS{Value: changes},
		},
		"timestamp": {
			Action: types.AttributeActionPut,
			Value:  &types.AttributeValueMemberS{Value: timestamp},
		},
		"versions": {
			Action: types.AttributeActionAdd,
			Value:  &types.AttributeValueMemberL{Value: []types.AttributeValue{&types.AttributeValueMemberS{Value: timestamp}}},
		},
	}

	input := &dynamodb.UpdateItemInput{
		TableName: &store.TableName,
		Key: map[string]types.AttributeValue{
			"document_id": &types.AttributeValueMemberS{Value: documentID},
		},
		AttributeUpdates: update,
		ReturnValues:     types.ReturnValueAllNew,
	}

	_, err := store.DynamoDB.UpdateItem(context.Background(), input)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}
