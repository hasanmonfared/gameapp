package presence

import (
	"context"
	"gameapp/contract/goproto/presence"
	"gameapp/param"
	"gameapp/pkg/protobufmapper"
	"gameapp/pkg/slice"
	"google.golang.org/grpc"
)

type Client struct {
	address string
}

func New(address string) Client {
	return Client{
		address: address,
	}
}
func (c Client) GetPresence(ctx context.Context, request param.GetPresenceRequest) (param.GetPresenceResponse, error) {
	conn, err := grpc.Dial(c.address, grpc.WithInsecure())
	if err != nil {
		return param.GetPresenceResponse{}, err
	}
	defer conn.Close()
	client := presence.NewPresenceServiceClient(conn)
	resp, err := client.GetPresence(ctx, &presence.GetPresenceRequest{UserIds: slice.MapFromUintToUint64(request.UserIDs)})
	if err != nil {
		return param.GetPresenceResponse{}, err
	}
	return protobufmapper.MapGetPresenceResponseFromProtobuf(resp), nil
}
