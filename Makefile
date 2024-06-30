get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	mkdir -p pkg/note_v1
	protoc --proto_path api/note_v1 \
    	--go_out=pkg/note_v1 --go_opt=paths=source_relative \
    	--go-grpc_out=pkg/note_v1 --go-grpc_opt=paths=source_relative \
    	api/note_v1/note.proto 