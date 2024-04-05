protoc --go_out=. test.proto
protoc --go-grpc_out=. test.proto

go build -o ./bin/server ./server/main.go

go build -o ./bin/client ./client/main.go

./bin/server

./bin/client