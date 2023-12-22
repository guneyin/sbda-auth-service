BINARY_NAME=auth

build:
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

output=../sbda-sdk/pb/

gen:
	protoc --go_out=$(output)  --go-grpc_out=$(output) proto/auth.proto