module github.com/RamzassH/LeadIt/notificationService

go 1.24.1

require (
	github.com/RamzassH/LeadIt/libs/kafka v0.0.0-20250323104901-2e098016fd82
	github.com/go-playground/validator/v10 v10.25.0
	github.com/ilyakaznacheev/cleanenv v1.5.0
	github.com/rs/zerolog v1.34.0
	google.golang.org/grpc v1.71.0
)

require (
	github.com/BurntSushi/toml v1.2.1 // indirect
	github.com/gabriel-vasile/mimetype v1.4.8 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pierrec/lz4/v4 v4.1.16 // indirect
	github.com/segmentio/kafka-go v0.4.47 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/protobuf v1.36.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace (
	github.com/RamzassH/LeadIt/libs/kafka => ../libs/kafka
	github.com/RamzassH/LeadIt/libs/contracts => ../libs/contracts
	github.com/RamzassH/LeadIt/libs/redis => ../libs/redis
)