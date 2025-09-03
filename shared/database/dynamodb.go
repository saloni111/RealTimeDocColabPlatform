package database

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// GetDynamoDBClient returns a configured DynamoDB client
// Automatically detects local vs production environment
func GetDynamoDBClient() (*dynamodb.Client, error) {
	var cfg aws.Config
	var err error

	// Check if we're in local development mode
	if os.Getenv("DYNAMODB_LOCAL") == "true" || os.Getenv("ENV") == "development" {
		cfg, err = getLocalDynamoDBConfig()
	} else {
		cfg, err = getProductionDynamoDBConfig()
	}

	if err != nil {
		log.Printf("Failed to load DynamoDB config: %v", err)
		return nil, err
	}

	return dynamodb.NewFromConfig(cfg), nil
}

// Local DynamoDB configuration for development
func getLocalDynamoDBConfig() (aws.Config, error) {
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			PartitionID:   "aws",
			URL:           "http://localhost:9000",
			SigningRegion: "us-east-1",
		}, nil
	})

	return config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     "local",
				SecretAccessKey: "local",
			}, nil
		})),
	)
}

// Production DynamoDB configuration
func getProductionDynamoDBConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(getRegion()),
	)
}

// Get AWS region from environment or default
func getRegion() string {
	if region := os.Getenv("AWS_REGION"); region != "" {
		return region
	}
	return "us-east-1" // Default region
}
