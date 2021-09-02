protoc -I=proto/rate --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative  --go_out proto/rate --go-grpc_out proto/rate $(find proto/rate -iname "*.proto")

# *** Make sure "github.com/grpc-ecosystem/grpc-gateway/v2" is installed 
# and on go.mod file ***
protoc -I=proto/rate --grpc-gateway_out=paths=source_relative,logtostderr=true:proto/rate $(find proto/rate -iname "*.proto")
