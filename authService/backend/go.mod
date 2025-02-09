module github.com/RamzassH/LeadIt/authService/backend

go 1.23.3

require github.com/ilyakaznacheev/cleanenv v1.5.0

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-playground/validator/v10 v10.23.0
	github.com/golang-migrate/migrate/v4 v4.18.1
	github.com/lib/pq v1.10.9
	golang.org/x/crypto v0.32.0
	golang.org/x/net v0.34.0
	google.golang.org/grpc v1.70.0
)

require (
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.26.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/leodido/go-urn v1.4.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

require (
	github.com/BurntSushi/toml v1.4.0 // indirect
	github.com/RamzassH/LeadIt/libs/contracts v0.0.1
	github.com/joho/godotenv v1.5.1 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	olympos.io/encoding/edn v0.0.0-20201019073823-d3554ca0b0a3 // indirect
)

replace github.com/RamzassH/LeadIt/libs/contracts => ../../libs/contracts

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20250207221924-e9438ea467c6
