# RealTimeDocColabPlatform

This is a simple collaborative document platform built with Go microservices. It lets users create, edit, and collaborate on documents in real time.

## Features
- User registration and login
- Create, update, and delete documents
- Real-time collaboration on documents
- API Gateway for easy access to all services

## How to Run
1. **Start all services:**
   - You can use the provided script:
     ```sh
     ./start_all.sh
     ```
   - Or run each service manually in its folder:
     ```sh
     go run main.go
     ```

2. **API Gateway:**
   - The API Gateway runs on `http://localhost:8080`

## Example API Usage
- **Create a document:**
  ```sh
  curl -X POST http://localhost:8080/document/create \
    -H "Content-Type: application/json" \
    -d '{"title":"Test Doc","content":"Hello world!","author":"yourname"}'
  ```

- **Register a user:**
  ```sh
  curl -X POST http://localhost:8080/user/register \
    -H "Content-Type: application/json" \
    -d '{"email":"test@example.com","password":"yourpassword","name":"Your Name"}'
  ```

## Project Structure
- `api-gateway/` - Entry point for all API requests
- `user-service/` - Handles user accounts
- `document-service/` - Manages documents
- `collaboration-service/` - Real-time collaboration logic
- `proto/` - Protobuf definitions for gRPC

## Requirements
- Go 1.22+
- (Optional) AWS account for DynamoDB, or use DynamoDB Local for development

## License
This project is for learning and demo purposes. 