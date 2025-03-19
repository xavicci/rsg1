github.com/xavicci/rsg1
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
winget install protobuf
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative studentpb/student.proto



docker build . -t server-grpc-db
docker run -p 54321:5432 server-grpc-db
go run server-student/main.go
