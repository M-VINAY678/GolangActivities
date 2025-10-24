package db

import (
	pb "chatApplication/pkg/protoFiles"

	"google.golang.org/grpc"
)

func NewChatManager() ChatManager {
	return &chatManager{
		clients: make(map[string]grpc.BidiStreamingServer[pb.ChatRequest, pb.ChatResponse]),
	}
}
func (c *chatManager) DeleteClient(userName string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.clients, userName)
}
func (c *chatManager) AddStream(userName string, stream grpc.BidiStreamingServer[pb.ChatRequest, pb.ChatResponse]) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.clients[userName] = stream
}
func (c *chatManager) SendChat(msg *pb.ChatRequest) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if msg.Receiver == "" {
		for _, cl := range c.clients {
			cl.Send(&pb.ChatResponse{Sender: msg.Sender, Text: msg.Text})
		}
	} else {
		if recevier, ok := c.clients[msg.Receiver]; ok {
			recevier.Send(&pb.ChatResponse{Sender: msg.Sender, Text: msg.Text})
		}
	}
}
