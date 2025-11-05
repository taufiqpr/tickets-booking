module ticket-booking/user-service

go 1.24.0

replace ticket-booking/proto => ../proto

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/jackc/pgx/v5 v5.6.0
	github.com/joho/godotenv v1.5.1
	golang.org/x/crypto v0.40.0
	google.golang.org/grpc v1.76.0
	ticket-booking/proto v0.0.0-00010101000000-000000000000
)

require (
	github.com/jackc/pgpassfile v1.0.0 
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a 
	github.com/jackc/puddle/v2 v2.2.1 
	golang.org/x/net v0.42.0 
	golang.org/x/sync v0.16.0 
	golang.org/x/sys v0.34.0 
	golang.org/x/text v0.27.0 
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250804133106-a7a43d27e69b 
	google.golang.org/protobuf v1.36.10 
)
