protoc --go_out=. --go-grpc_out=. --proto_path=adapters/protobuf/service-account types.account.proto
protoc --go_out=. --go-grpc_out=. --proto_path=adapters/protobuf/service-account api.account.proto
protoc --go_out=. --go-grpc_out=. --proto_path=adapters/protobuf/service-transaction api.transaction.proto
protoc --go_out=. --go-grpc_out=. --proto_path=adapters/protobuf/service-transaction format.transaction.proto