package main

import (
	"context"
	"io"
	"log"
	"net"
	"time"

	pb "module_12_gRPC/route_guide"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedRouteGuideServer
}

// Simple RPC
func (s *server) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	log.Printf("GetFeature request: %v,%v", point.Latitude, point.Longitude)
	return &pb.Feature{
		Name:     "Example Feature",
		Location: point,
	}, nil
}

// Server-side streaming RPC
func (s *server) ListFeatures(rect *pb.Rectangle, stream grpc.ServerStreamingServer[pb.Feature]) error {
	for i := 0; i < 5; i++ { // send 5 features
		feature := &pb.Feature{
			Name:     "Feature " + string('A'+i),
			Location: &pb.Point{Latitude: rect.Lo.Latitude + int32(i), Longitude: rect.Lo.Longitude + int32(i)},
		}
		if err := stream.Send(feature); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 2000)
	}
	return nil
}

// Client-side streaming RPC
func (s *server) RecordRoute(stream pb.RouteGuide_RecordRouteServer) error {
	var count int32
	for {
		point, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.RouteSummary{
				PointCount:   count,
				FeatureCount: count / 2,
				Distance:     100,
				ElapsedTime:  1,
			})
		}
		if err != nil {
			return err
		}
		log.Printf("RecordRoute received: %v,%v", point.Latitude, point.Longitude)
		count++
	}
}

// Bidirectional streaming RPC
func (s *server) RouteChat(stream pb.RouteGuide_RouteChatServer) error {
	for {
		note, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("RouteChat received: %s at %v,%v", note.Message, note.Location.Latitude, note.Location.Longitude)
		reply := &pb.RouteNote{
			Message:  "Ack: " + note.Message,
			Location: note.Location,
		}
		if err := stream.Send(reply); err != nil {
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, &server{})
	log.Println("Server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
