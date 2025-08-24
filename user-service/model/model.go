package model

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type User struct {
	UserID   string `dynamodbav:"user_id"`
	Email    string `dynamodbav:"email"`
	Password string `dynamodbav:"password"`
	UserName string `dynamodbav:"username"`
}

type UserModel struct {
	DynamoDB  *dynamodb.Client
	TableName string
}

func (m *UserModel) CreateUser(ctx context.Context, user *User) error {
	av, err := attributevalue.MarshalMap(user)

	if err != nil {
		return fmt.Errorf("failed to create new User %w", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &m.TableName,
	}

	_, err = m.DynamoDB.PutItem(ctx, input)

	if err != nil {
		return fmt.Errorf("failed to put item: %w", err)
	}

	return nil
}

func (m *UserModel) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	// Since email is not the primary key, we need to scan
	input := &dynamodb.ScanInput{
		TableName:        &m.TableName,
		FilterExpression: aws.String("email = :email"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{Value: email},
		},
	}

	result, err := m.DynamoDB.Scan(ctx, input)

	if err != nil {
		return nil, fmt.Errorf("failed to scan items: %w", err)
	}

	if len(result.Items) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	var user User

	err = attributevalue.UnmarshalMap(result.Items[0], &user)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &user, nil
}

func (m *UserModel) GetUserById(ctx context.Context, user_id string) (*User, error) {
	// Use GetItem since user_id is the primary key
	input := &dynamodb.GetItemInput{
		TableName: &m.TableName,
		Key: map[string]types.AttributeValue{
			"user_id": &types.AttributeValueMemberS{Value: user_id},
		},
		ProjectionExpression: aws.String("user_id, email, username"), // Specify the attributes you want to retrieve
	}

	result, err := m.DynamoDB.GetItem(ctx, input)

	if err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}

	if result.Item == nil {
		return nil, fmt.Errorf("user not found")
	}

	var user User

	err = attributevalue.UnmarshalMap(result.Item, &user)

	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user: %w", err)
	}

	return &user, nil
}
