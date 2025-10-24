package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	pb "chatApplication/pkg/protoFiles"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Connect to the gRPC server
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Minute)
	defer cancel()

	// Open ChatRoom stream
	stream, err := client.ChatRoom(ctx)
	if err != nil {
		log.Fatalf("Error creating stream: %v", err)
	}

	// Goroutine: receive messages from server
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				log.Println("Server closed connection.")
				return
			}
			if err != nil {
				log.Printf("Receive error: %v", err)
				return
			}

			fmt.Printf("\n[%s]: %s\n", res.Sender, res.Text)
		}
	}()

	// Main goroutine: send messages
	for {
		fmt.Print("Enter Your Message: ")
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)
		err := stream.Send(&pb.ChatRequest{
			Sender:   username,
			Receiver: "",
			Text:     msg,
		})
		if err != nil {
			log.Printf("Send error: %v", err)
			break
		}
		time.Sleep(20 * time.Millisecond)
	}
}
