package protobufencoder

import (
	"encoding/base64"
	"gameapp/contract/goproto/notification"
	"gameapp/entity"
	"google.golang.org/protobuf/proto"
)

func EncodeNotification(mu entity.Notification) string {

	pbMu := notification.Notification{
		Type:    mu.EventType,
		Payload: mu.Payload,
	}
	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}
func DecodeNotification(data string) entity.Notification {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return entity.Notification{}

	}
	pbMu := notification.Notification{}
	if err := proto.Unmarshal(payload, &pbMu); err != nil {
		return entity.Notification{}
	}

	return entity.Notification{
		EventType: pbMu.Type,
		Payload:   pbMu.Payload,
	}
}
