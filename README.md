# golang-grpc-example

간단한 grpc 테스트

사용법
1. proto파일 작성 (internal/greet/greet.proto)
2. pb파일 생성 (protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative internal/greet/greet.proto)
3. 서버실행 후 유지(go run cmd/server/*)
4. 클라이언트 실행(go run cmd/client/client.go)