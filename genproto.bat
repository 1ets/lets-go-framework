protoc --go_out=. --go-grpc_out=. --proto_path=adapters/protobuf/service-account service.account.proto
protoc --go_out=. --go-grpc_out=. --proto_path=adapters/protobuf/service-transaction service.transaction.proto
protoc --go_out=. --go-grpc_out=. --proto_path=adapters/protobuf/service-transaction message.transaction.proto