package main

import (
	"context"
	"io"
	"log"

	pb "module_12_gRPC/route_guide"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewRouteGuideClient(conn)

	// Simple RPC
	point := &pb.Point{Latitude: 409146138, Longitude: -746188906}
	feature, err := client.GetFeature(context.Background(), point)
	if err != nil {
		log.Fatalf("GetFeature failed: %v", err)
	}
	log.Printf("GetFeature: %v", feature)

	//Server streaming RPC
	rect := &pb.Rectangle{
		Lo: &pb.Point{Latitude: 0, Longitude: 0},
		Hi: &pb.Point{Latitude: 10, Longitude: 10},
	}
	stream1, _ := client.ListFeatures(context.Background(), rect)
	for {
		f, err := stream1.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListFeatures error: %v", err)
		}
		log.Printf("Feature: %v", f)
	}

	// Client streaming RPC
	stream2, _ := client.RecordRoute(context.Background())
	for i := 0; i < 5; i++ {
		stream2.Send(&pb.Point{Latitude: int32(i), Longitude: int32(i)})
	}
	summary, _ := stream2.CloseAndRecv()
	log.Printf("RouteSummary: %v", summary)

	// Bidirectional streaming RPC
	stream3, _ := client.RouteChat(context.Background())
	waitc := make(chan struct{})
	go func() {
		for {
			note, err := stream3.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("RouteChat error: %v", err)
			}
			log.Printf("Received note: %v", note)
		}
	}()
	for i := 0; i < 3; i++ {
		stream3.Send(&pb.RouteNote{
			Location: &pb.Point{Latitude: int32(i), Longitude: int32(i)},
			Message:  "Message " + string('A'+i),
		})
	}
	stream3.CloseSend()
	<-waitc
}
