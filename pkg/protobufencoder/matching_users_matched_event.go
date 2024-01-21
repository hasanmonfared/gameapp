package protobufencoder

import (
	"encoding/base64"
	"gameapp/contract/golang/matching"
	"gameapp/entity"
	"gameapp/pkg/slice"
	"google.golang.org/protobuf/proto"
)

func EncodeEventMatchingUserMatchedEvent(mu entity.MatchUsers) string {

	pbMu := matching.MatchUsers{
		Category: string(mu.Category),
		UserIds:  slice.MapFromUintToUint64(mu.UserIDs),
	}
	payload, err := proto.Marshal(&pbMu)
	if err != nil {
		return ""
	}

	return base64.StdEncoding.EncodeToString(payload)
}
func DecodeEventEventMatchingUserMatchedEvent(data string) entity.MatchUsers {
	payload, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return entity.MatchUsers{}

	}

	pbMu := matching.MatchUsers{}
	if err := proto.Unmarshal(payload, &pbMu); err != nil {
		return entity.MatchUsers{}
	}

	return entity.MatchUsers{
		Category: entity.Category(pbMu.Category),
		UserIDs:  slice.MapFromUint64ToUint(pbMu.UserIds),
	}
}
