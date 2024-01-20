package presenceserver

import (
	"context"
	"fmt"
	"gameapp/contract/golang/presence"
	"gameapp/param"
	"gameapp/pkg/protobufmapper"
	"gameapp/pkg/slice"
	"gameapp/service/presenceservice"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	presence.UnimplementedPresenceServiceServer
	svc presenceservice.Service
}

func New(svc presenceservice.Service) Server {
	return Server{
		UnimplementedPresenceServiceServer: presence.UnimplementedPresenceServiceServer{},
		svc:                                svc,
	}
}
func (s Server) GetPresence(ctx context.Context, req *presence.GetPresenceRequest) (*presence.GetPresenceResponse, error) {
	resp, err := s.svc.GetPresence(ctx, param.GetPresenceRequest{UserIDs: slice.MapFromUint64ToUint(req.GetUserIds())})
	if err != nil {
		return nil, err
	}
	return protobufmapper.MapGetPresenceResponseToProtobuf(resp), nil
}

func (s Server) Start() {
	address := fmt.Sprintf(":%d", 8086)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	presenceSvcServer := Server{}

	grpcServer := grpc.NewServer()

	log.Printf("presence grpc server started on %d", address)

	presence.RegisterPresenceServiceServer(grpcServer, &presenceSvcServer)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("couldn't server presence grpc server", err)
	}
	log.Printf("presence grpc server started on %d", address)
}
