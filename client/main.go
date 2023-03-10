package main

import (
	pb "grpc-notes/client/notes/pb"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStorageClient(conn)
	noteList, err := client.GetNoteList(context.Background(), &pb.GetNoteListRequest{})
	if err != nil {
		log.Fatalf("failed to get notes list: %v", err)
	}
	log.Printf("note list: %v", noteList)
}