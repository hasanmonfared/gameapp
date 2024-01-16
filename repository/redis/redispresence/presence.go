package redispresence

import (
	"context"
	"gameapp/pkg/richerror"
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
