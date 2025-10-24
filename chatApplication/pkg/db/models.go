package db

import (
	"sync"

	pb "chatApplication/pkg/protoFiles"

	"google.golang.org/grpc"
)

type ChatManager interface {
	DeleteClient(userName string)
	AddStream(userName string, stream grpc.BidiStreamingServer[pb.ChatRequest, pb.ChatResponse])
	SendChat(msg *pb.ChatRequest)
}
type chatManager struct {
	clients map[string]pb.ChatService_ChatRoomServer
	mu      sync.Mutex
}
