module github.com/RamzassH/LeadIt/gateway

go 1.23.3

require (
	github.com/RamzassH/LeadIt/libs/contracts v0.0.0-20250209185459-c146a87bcf1d
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.1
	github.com/rs/cors v1.11.1
	google.golang.org/grpc v1.71.0
)

require (
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

replace github.com/RamzassH/LeadIt/libs/contracts => ../libs/contracts
