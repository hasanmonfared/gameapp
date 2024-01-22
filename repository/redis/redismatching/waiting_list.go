package redismatching

import (
	"context"
	"fmt"
	"gameapp/entity"
	"gameapp/pkg/richerror"
	timestamp "gameapp/pkg/timsestamp"
	"github.com/labstack/gommon/log"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
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
func (d DB) GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error) {
	const op = richerror.Op("redismatching.GetWaitingListByCategory")
	minimum := strconv.Itoa(int(timestamp.Add(-2 * time.Hour)))
	maximum := strconv.Itoa(int(timestamp.Add(-2 * time.Hour)))
	list, err := d.adapter.Client().ZRevRangeByScoreWithScores(ctx, getCategoryKey(category), &redis.ZRangeBy{
		Min:    minimum,
		Max:    maximum,
		Offset: 0,
		Count:  0,
	}).Result()
	if err != nil {
		return nil, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	var result = make([]entity.WaitingMember, 0)
	for _, l := range list {

		userID, _ := strconv.Atoi(l.Member.(string))

		result = append(result, entity.WaitingMember{
			UserID:    uint(userID),
			Timestamp: int64(l.Score),
			Category:  category,
		})
	}
	return result, nil
}
func (d DB) RemoveUsersFromWaitingList(category entity.Category, userIDs []uint) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	members := make([]any, 0)
	for _, u := range userIDs {
		members = append(members, strconv.Itoa(int(u)))
	}

	numberOfRemovedMembers, err := d.adapter.Client().ZRem(ctx, getCategoryKey(category), members...).Result()
	if err != nil {
		log.Errorf("remove from waiting list: %v", err)
	}
	log.Printf("%d items remove from %s", numberOfRemovedMembers, getCategoryKey(category))
}
func getCategoryKey(category entity.Category) string {
	return fmt.Sprintf("%s:%s", WaitingListPrefix, category)
}
