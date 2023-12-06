build:
	rm -f shipper
	go build -o shipper main.go
run:
	./shipper
proto:
	protoc --go_out=plugins=grpc:. server/handlers/shipper/proto/shipper.proto
