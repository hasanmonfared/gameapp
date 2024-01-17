package matchingservice

import (
	"context"
	"fmt"
	"gameapp/entity"
	"gameapp/param"
	"gameapp/pkg/richerror"
	timestamp "gameapp/pkg/timsestamp"
	funk "github.com/thoas/go-funk"
	"sync"
	"time"
)

type Repo interface {
	AddToWaitingList(userID uint, category entity.Category) error
	GetWaitingListByCategory(ctx context.Context, category entity.Category) ([]entity.WaitingMember, error)
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
}

func New(config Config, repo Repo, presenceClient PresenceClient) Service {
	return Service{
		config:         config,
		repo:           repo,
		presenceClient: presenceClient,
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
	var wg *sync.WaitGroup
	for _, category := range entity.CategoryList() {
		wg.Add(1)
		go s.match(ctx, category, wg)
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
	userIDs := make([]uint, len(list))
	for _, l := range list {
		userIDs = append(userIDs, l.UserID)
	}
	presenceList, err := s.presenceClient.GetPresence(ctx, param.GetPresenceRequest{UserID: userIDs})
	if err != nil {
		return
	}
	presenceUserIDs := make([]uint, len(list))
	for _, l := range presenceList.Items {
		presenceUserIDs = append(presenceUserIDs, l.UserID)
	}
	var finalList = make([]entity.WaitingMember, 0)
	for _, l := range list {
		if funk.ContainsUInt(presenceUserIDs, l.UserID) && l.Timestamp < timestamp.Add(-20*time.Second) {
			finalList = append(finalList, l)
		} else {
			// remove from list
		}
	}

	for i := 0; i < len(list)-1; i = i + 2 {
		mu := entity.MatchUsers{
			Category: category,
			UserID:   []uint{list[i].UserID, list[i+1].UserID},
		}
		fmt.Println(mu)
	}
}
