package protobuf

import (
	"gameapp/contract/golang/presence"
	"gameapp/param"
)

func MapGetPresenceResponseToProtobuf(g param.GetPresenceResponse) *presence.GetPresenceResponse {
	r := &presence.GetPresenceResponse{}
	for _, item := range g.Items {
		r.Items = append(r.Items, &presence.GetPresenceItem{
			UserId:    uint64(item.UserID),
			Timestamp: item.Timestamp,
		})
	}
	return r
}
