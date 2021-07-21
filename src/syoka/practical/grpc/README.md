### install gen

go get -u github.com/golang/protobuf/protoc-gen-go

### install protoc

brew install protobuf

### generate proto go

protoc --go_out=plugins=grpc:../../practical calculator.proto

