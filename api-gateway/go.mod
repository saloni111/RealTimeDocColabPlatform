module github.com/saloni111/RealTimeDocColabPlatform/api-gateway

go 1.22.5

require github.com/gorilla/mux v1.8.1

require (
	github.com/saloni111/RealTimeDocColabPlatform/collaboration-service v0.0.0
	github.com/saloni111/RealTimeDocColabPlatform/document-service v0.0.0
	github.com/saloni111/RealTimeDocColabPlatform/user-service v0.0.0
	google.golang.org/grpc v1.65.0
)

require (
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240711142825-46eb208f015d // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/saloni111/RealTimeDocColabPlatform/user-service => ../user-service

replace github.com/saloni111/RealTimeDocColabPlatform/document-service => ../document-service

replace github.com/saloni111/RealTimeDocColabPlatform/collaboration-service => ../collaboration-service
