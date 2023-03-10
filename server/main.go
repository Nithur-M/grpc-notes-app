package main

import (
	pb "grpc-notes/server/notes/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedStorageServer
}

func (s *server) GetNoteList(ctx context.Context, in *pb.GetNoteListRequest) (*pb.GetNoteListResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	return &pb.GetNoteListResponse{
		Notes: getSampleNotes(),
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterStorageServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func getSampleNotes() []*pb.Note {
	sampleNotes := []*pb.Note{
		{
			Title:     "My first Note",
			Body:      "This is my first note",
		},
		{
			Title:     "My second note",
			Body:      "This is my second note",
		},
	}
	return sampleNotes
}