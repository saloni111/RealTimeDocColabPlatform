#!/bin/bash

# Start user-service
cd user-service
nohup go run main.go > ../user-service.log 2>&1 &
echo "Started user-service"

# Start document-service
cd ../document-service
nohup go run main.go > ../document-service.log 2>&1 &
echo "Started document-service"

# Start collaboration-service
cd ../collaboration-service
nohup go run main.go > ../collaboration-service.log 2>&1 &
echo "Started collaboration-service"

# Start api-gateway
cd ../api-gateway
nohup go run main.go > ../api-gateway.log 2>&1 &
echo "Started api-gateway"

cd ..
echo "All services started. Check *.log files for output." 