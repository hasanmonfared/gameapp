package matchingservice

import (
	"context"
	"fmt"
	"gameapp/entity"
	"gameapp/param"
	"gameapp/pkg/protobufencoder"
	"gameapp/pkg/richerror"
	timestamp "gameapp/pkg/timsestamp"
	"sync"
	"time"
)

type Publisher interface {
	Publish(event entity.Event, payload string)
}
type Repo interface {
	AddToWaitingList(userID uint, category entity.Category) error
	GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error)
	removeUsersFromWaitingList(category entity.Category, userIDs []uint)
}
type PresenceClient interface {
	GetPresence(ctx context.Context, request param.GetPresenceRequest) (param.GetPresenceResponse, error)
}
type Config struct {
	WaitingTimeout time.Duration `koanf:"waiting_timeout"`
}
type Service struct {
	repo           Repo
	config         Config
	presenceClient PresenceClient
	pub            Publisher
}

func New(config Config, repo Repo, presenceClient PresenceClient, pub Publisher) Service {
	return Service{
		config:         config,
		repo:           repo,
		presenceClient: presenceClient,
		pub:            pub,
	}
}
func (s Service) AddToWaitingList(req param.AddToWaitingListRequest) (param.AddToWaitingListResponse, error) {
	const op = richerror.Op("matchingservice.AddToWaitingList")
	err := s.repo.AddToWaitingList(req.UserID, req.Category)
	if err != nil {
		return param.AddToWaitingListResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return param.AddToWaitingListResponse{Timeout: s.config.WaitingTimeout}, nil
}

func (s Service) MatchWaitedUsers(ctx context.Context, _ param.MatchWaitedUsersRequest) (param.MatchWaitedUsersResponse, error) {
	const op = richerror.Op("matchingservice.MatchWaitedUsers")
	var wg sync.WaitGroup
	for _, category := range entity.CategoryList() {
		wg.Add(1)
		go s.match(ctx, category, &wg)
	}
	wg.Wait()
	return param.MatchWaitedUsersResponse{}, nil
}

func (s Service) match(ctx context.Context, category entity.Category, wg *sync.WaitGroup) {
	const op = richerror.Op("matchingservice.match")

	defer wg.Done()
	list, err := s.repo.GetWaitingListByCategory(ctx, category)

	if err != nil {
		return
	}
	userIDs := make([]uint, 0)
	for _, l := range list {
		userIDs = append(userIDs, l.UserID)
	}
	if len(userIDs) < 2 {
		return
	}
	presenceList, err := s.presenceClient.GetPresence(ctx, param.GetPresenceRequest{UserIDs: userIDs})
	if err != nil {
		return
	}
	presenceUserIDs := make([]uint, len(list))
	for _, l := range presenceList.Items {
		presenceUserIDs = append(presenceUserIDs, l.UserID)
	}
	var toBeRemoveUsers = make([]uint, 0)
	var finalList = make([]entity.WaitingMember, 0)
	for _, l := range list {
		lastOnlineTimestamp, ok := getPresenceItem(presenceList, l.UserID)
		if ok &&
			lastOnlineTimestamp > timestamp.Add(-20*time.Second) &&
			l.Timestamp > timestamp.Add(-300*time.Second) {
			finalList = append(finalList, l)
		} else {
			toBeRemoveUsers = append(toBeRemoveUsers, l.UserID)
		}
	}
	go s.repo.removeUsersFromWaitingList(category, toBeRemoveUsers)
	matchedUsersToBeRemoved := make([]uint, 0)
	for i := 0; i < len(list)-1; i = i + 2 {
		mu := entity.MatchUsers{
			Category: category,
			UserIDs:  []uint{list[i].UserID, list[i+1].UserID},
		}
		fmt.Println(mu)
		go s.pub.Publish(entity.MatchingUserMatchedEvent, protobufencoder.EncodeEventMatchingUserMatchedEvent(mu))

		matchedUsersToBeRemoved = append(matchedUsersToBeRemoved, mu.UserIDs...)
	}
	go s.repo.removeUsersFromWaitingList(category, matchedUsersToBeRemoved)

}

func getPresenceItem(presenceList param.GetPresenceResponse, userID uint) (int64, bool) {
	for _, item := range presenceList.Items {
		if item.UserID == userID {
			return item.Timestamp, true
		}
	}
	return 0, false
}
