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
	github.com/aws/aws-sdk-go-v2 v1.38.1 // indirect
	github.com/aws/aws-sdk-go-v2/config v1.31.2 // indirect
	github.com/aws/aws-sdk-go-v2/credentials v1.18.6 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.4 // indirect
	github.com/aws/aws-sdk-go-v2/internal/ini v1.8.3 // indirect
	github.com/aws/aws-sdk-go-v2/service/dynamodb v1.49.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.0 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/endpoint-discovery v1.11.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.4 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.28.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.33.2 // indirect
	github.com/aws/aws-sdk-go-v2/service/sts v1.38.0 // indirect
	github.com/aws/smithy-go v1.22.5 // indirect
	golang.org/x/net v0.27.0 // indirect
	golang.org/x/sys v0.22.0 // indirect
	golang.org/x/text v0.16.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240711142825-46eb208f015d // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)

replace github.com/saloni111/RealTimeDocColabPlatform/user-service => ./user-service

replace github.com/saloni111/RealTimeDocColabPlatform/document-service => ./document-service

replace github.com/saloni111/RealTimeDocColabPlatform/collaboration-service => ./collaboration-service
