package redis

import (
	"context"
	"gameapp/entity"
	"github.com/labstack/gommon/log"
	"time"
)

func (a Adapter) Publish(event entity.Event, payload string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	if err := a.client.Publish(ctx, string(event), payload).Err(); err != nil {
		log.Errorf("publish err %v\n", err)
	}
}
