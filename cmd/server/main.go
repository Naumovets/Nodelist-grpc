package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/Naumovets/Nodelist-grpc/pkg/note_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const grpsPort = 50051

type server struct {
	desc.UnimplementedNoteV1Server
}

func (s *server) Create(context.Context, *desc.CreateRequest) (*desc.CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s *server) Get(context.Context, *desc.GetRequest) (*desc.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (s *server) List(context.Context, *desc.ListRequest) (*desc.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (s *server) Update(context.Context, *desc.UpdateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (s *server) Delete(context.Context, *desc.DeleteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func main() {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpsPort))

	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	s := grpc.NewServer()

	reflection.Register(s)

	desc.RegisterNoteV1Server(s, &server{})

	log.Printf("server listening at: %s", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
