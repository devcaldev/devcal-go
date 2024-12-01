wget https://raw.githubusercontent.com/devcaldev/proto/refs/heads/main/devcal.proto

protoc --go_out=./rpc --go_opt=paths=source_relative \
       --go-grpc_out=./rpc --go-grpc_opt=paths=source_relative \
       devcal.proto

rm devcal.proto