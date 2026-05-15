package utils

import "os"

func envOr(key, defaultVal string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultVal
}

// UserServiceGRPCAddr returns USER_SERVICE_ADDR or local default for dev.
func UserServiceGRPCAddr() string {
	return envOr("USER_SERVICE_ADDR", "localhost:50051")
}

// DocumentServiceGRPCAddr returns DOCUMENT_SERVICE_ADDR or local default for dev.
func DocumentServiceGRPCAddr() string {
	return envOr("DOCUMENT_SERVICE_ADDR", "localhost:50052")
}

// CollaborationServiceGRPCAddr returns COLLABORATION_SERVICE_ADDR or local default.
func CollaborationServiceGRPCAddr() string {
	return envOr("COLLABORATION_SERVICE_ADDR", "localhost:50053")
}

// DynamoDBEndpointURL returns HTTP URL for DynamoDB (local Docker or DynamoDB Local on host).
func DynamoDBEndpointURL() string {
	return envOr("DYNAMODB_ENDPOINT", "http://localhost:9000")
}
