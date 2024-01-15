package redispresence

import "context"

func (d DB) Upsert(ctx context.Context, key string, timestamp int64) error {
	d.adapter.Client().Set(ctx,key,timestamp)
}
section 2
26 min