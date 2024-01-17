package presenceservice

import (
	"fmt"
	"gameapp/param"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
	"time"
)

type Config struct {
	ExpirationTime time.Duration `koanf:"expiration_time"`
	Prefix         string        `koanf:"prefix"`
}
type Repo interface {
	Upsert(ctx context.Context, key string, timestamp int64, expTime time.Duration) error
}
type Service struct {
	config Config
	repo   Repo
}

func New(config Config, repo Repo) Service {
	return Service{repo: repo, config: config}
}

func (s Service) Upsert(ctx context.Context, req param.UpsertPresenceRequest) (param.UpsertPresenceResponse, error) {
	const op = richerror.Op("presenceservice.Upsert")
	err := s.repo.Upsert(ctx, fmt.Sprintf("%s:%d", s.config.Prefix, req.UserID), req.Timestamp, s.config.ExpirationTime)
	if err != nil {
		return param.UpsertPresenceResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return param.UpsertPresenceResponse{}, nil
}
func (s Service) GetPresence(ctx context.Context, request param.GetPresenceRequest) (param.GetPresenceResponse, error) {
	fmt.Println("req", request)
	return param.GetPresenceResponse{Items: []param.GetPresenceItem{
		{UserID: 1, Timestamp: 12345678},
		{UserID: 2, Timestamp: 12345673},
	}}, nil
}
