build:
	rm -f service/shipper
	go build -o service/shipper service/main.go
run:
	./service/shipper
proto:
	protoc --go_out=plugins=grpc:. service/server/handlers/shipper/proto/shipper.proto
