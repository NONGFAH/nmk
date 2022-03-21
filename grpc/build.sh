protoc --go_out=plugins=grpc:./ ./nmk.proto
mv ./github.com/nongfah/nmk/grpc/nmk.pb.go .
rm -rf ./github.com