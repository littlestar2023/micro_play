protoc -I ../idl/protoc --go_out=../idl/pb --micro_out=../idl/pb userService.proto
protoc -I ../idl/protoc --go_out=../idl/pb --micro_out=../idl/pb taskService.proto

protoc-go-inject-tag -input="..\idl\pb\userService.pb.go"
protoc-go-inject-tag -input="..\idl\pb\userService.pb.micro.go"
protoc-go-inject-tag -input="..\idl\pb\taskService.pb.go"
protoc-go-inject-tag -input="..\idl\pb\taskService.pb.micro.go"