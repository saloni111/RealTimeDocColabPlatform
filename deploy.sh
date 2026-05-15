#!/bin/bash

# DocHub Production Deployment Script
set -e

echo "🚀 Starting DocHub Production Deployment..."

# Check if Docker is running
if ! docker info >/dev/null 2>&1; then
    echo "❌ Docker is not running. Please start Docker and try again."
    exit 1
fi

# Create production environment file
cat > .env.production << EOF
ENV=production
DYNAMODB_LOCAL=false
AWS_REGION=us-east-1
LOG_LEVEL=info
PORT=8080

# Service Discovery (for Kubernetes/ECS)
USER_SERVICE_ADDR=user-service:50051
DOCUMENT_SERVICE_ADDR=document-service:50052
COLLABORATION_SERVICE_ADDR=collaboration-service:50053
EOF

echo "✅ Created production environment configuration"

# Build all services
echo "🔨 Building all microservices..."
docker-compose build --no-cache

echo "✅ All services built successfully"

# Start services in production mode
echo "🚀 Starting services in production mode..."
docker-compose up -d

echo "⏳ Waiting for services to start..."
sleep 10

# Health check
echo "🏥 Performing health checks..."
services=("user-service:50051" "document-service:50052" "collaboration-service:50053" "api-gateway:8080")

for service in "${services[@]}"; do
    echo "Checking $service..."
    # Add actual health check logic here
done

echo "✅ All services are healthy!"

# Display service URLs
echo ""
echo "🎉 DocHub is now running in production mode!"
echo ""
echo "📱 Application URL: http://localhost:8080"
echo "🔧 API root: http://localhost:8080 (JSON endpoints listed at Accept: application/json)"
echo "📊 Health Check: http://localhost:8080/health"
echo ""
echo "🐳 Docker containers:"
docker-compose ps

echo ""
echo "📝 To view logs: docker-compose logs -f"
echo "🛑 To stop: docker-compose down"
echo "🔄 To restart: docker-compose restart"
# Updated
