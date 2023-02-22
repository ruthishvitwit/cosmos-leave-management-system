proto-gen:
	protoc --go_out=.  --go-grpc_out=. ./proto/cosmos/std/v1beta1/*.proto