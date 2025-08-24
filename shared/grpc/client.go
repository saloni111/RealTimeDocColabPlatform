package grpc

import (
	"context"
	"fmt"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

// ServiceConfig holds gRPC service configuration
type ServiceConfig struct {
	Host    string
	Port    string
	Timeout time.Duration
}

// GetGRPCConnection creates a gRPC connection with production-ready settings
func GetGRPCConnection(address string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Production-ready gRPC options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()), // TODO: Use TLS in production
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             time.Second,
			PermitWithoutStream: true,
		}),
		grpc.WithBlock(),
	}

	conn, err := grpc.DialContext(ctx, address, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %s: %w", address, err)
	}

	return conn, nil
}

// GetServiceAddress returns the service address based on environment
func GetServiceAddress(serviceName string, defaultPort string) string {
	// Check for environment-specific service discovery
	if addr := os.Getenv(fmt.Sprintf("%s_SERVICE_ADDR", serviceName)); addr != "" {
		return addr
	}

	// Default to localhost for development
	return fmt.Sprintf("localhost:%s", defaultPort)
}

// ServicePorts defines default ports for each service
var ServicePorts = map[string]string{
	"USER_SERVICE":          "50051",
	"DOCUMENT_SERVICE":      "50052",
	"COLLABORATION_SERVICE": "50053",
}
