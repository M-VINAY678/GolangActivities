package main

import (
	"io"
	"log"
	"net"

	db "chatApplication/pkg/db"
	pb "chatApplication/pkg/protoFiles"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
	chatMgr db.ChatManager
}

func NewServer(chatMgr db.ChatManager) *server {
	return &server{
		chatMgr: chatMgr,
	}
}

// ChatRoom handles group and private messages
/*func (s *server) ChatRoom(stream pb.ChatService_ChatRoomServer) error {
	return s.chatMgr.ChatRoom(stream)
}*/
func (s *server) ChatRoom(stream pb.ChatService_ChatRoomServer) error {
	var userName string
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			s.chatMgr.DeleteClient(userName)
			return nil
		}
		if err != nil {
			return err
		}
		userName = msg.Sender
		s.chatMgr.AddStream(userName, stream)
		s.chatMgr.SendChat(msg)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	chatMgr := db.NewChatManager()
	pb.RegisterChatServiceServer(grpcServer, NewServer(chatMgr))
	log.Println("Server listening on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
