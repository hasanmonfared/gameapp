package redismatching

import (
	"context"
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/richerror"
	timestamp "gameapp/pkg/timsestamp"
	"github.com/redis/go-redis/v9"
)

const WaitingListPrefix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = richerror.Op("redismatching.AddToWaitingList")

	_, err := d.adapter.Client().ZAdd(context.Background(), fmt.Sprintf("%s:%s", WaitingListPrefix, category), redis.Z{
		Score:  float64(timestamp.Now()),
		Member: fmt.Sprintf("%s", userID),
	}).Result()

	if err != nil {
		richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
