module github.com/RamzassH/LeadIt/gateway

go 1.23.3

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3
	google.golang.org/genproto/googleapis/api v0.0.0-20250124145028-65684f501c47
	google.golang.org/grpc v1.70.0
	google.golang.org/protobuf v1.36.4
)

require (
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto v0.0.0-20220822174746-9e6da59bd2fc // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
)

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20250207221924-e9438ea467c6
