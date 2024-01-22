package redispresence

import (
	"context"
	"gameapp/pkg/richerror"
	timestamp "gameapp/pkg/timsestamp"
	"time"
)

func (d DB) Upsert(ctx context.Context, key string, timestamp int64, exp time.Duration) error {
	const op = richerror.Op("redispresence.Upsert")
	_, err := d.adapter.Client().Set(ctx, key, timestamp, exp).Result()
	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return nil
}
func (d DB) GetPresence(ctx context.Context, prefixKey string, userIDs []uint) (map[uint]int64, error) {
	const op = richerror.Op("redispresence.GetPresence")
	m := make(map[uint]int64)
	for _, u := range userIDs {
		m[u] = timestamp.Add(time.Millisecond * -100)
	}
	return m, nil
}
