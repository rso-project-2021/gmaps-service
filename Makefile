test: 
	go test -cover ./...

protoc:
	protoc --proto_path=proto --go_out=proto --go_opt=paths=source_relative proto/location.proto

.PHONY: test, protoc