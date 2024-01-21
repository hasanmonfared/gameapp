package protobufencoder

import (
	"encoding/base64"
	"gameapp/contract/golang/matching"
	"gameapp/entity"
	"gameapp/pkg/slice"
	"google.golang.org/protobuf/proto"
)

func EncodeEvent(event entity.Event, data any) string {
	var payload []byte
	switch event {
	case entity.MatchingUserMatchedEvent:
		mu, ok := data.(entity.MatchUsers)
		if !ok {
			return ""
		}
		pbMu := matching.MatchUsers{
			Category: string(mu.Category),
			UserIds:  slice.MapFromUintToUint64(mu.UserIDs),
		}
		var err error
		payload, err = proto.Marshal(&pbMu)
		if err != nil {
			panic(err)
		}
	}
	return base64.StdEncoding.EncodeToString(payload)
}
func DecodeEvent(event entity.Event, data string) any {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil

	}
	switch event {
	case entity.MatchingUserMatchedEvent:
		pbMu := matching.MatchUsers{}
		if err := proto.Unmarshal(payload, &pbMu); err != nil {
			return nil
		}

		return entity.MatchUsers{
			Category: entity.Category(pbMu.Category),
			UserIDs:  slice.MapFromUint64ToUint(pbMu.UserIds),
		}
	}
	return nil
}
