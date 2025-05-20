module github.com/mikhailvzhzhv/grpc-ping-pong/pong_service

go 1.23

toolchain go1.23.9

require google.golang.org/grpc v1.72.1

require (
	github.com/joho/godotenv v1.5.1
	google.golang.org/protobuf v1.36.6 // indirect
)

require (
	github.com/mikhailvzhzhv/grpc-ping-pong/proto v0.0.0-20250520064117-a69e5a94d075
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)
