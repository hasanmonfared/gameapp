package presenceservice

import (
	"fmt"
	"gameapp/param"
	"gameapp/pkg/richerror"
	"golang.org/x/net/context"
)

type Config struct {
	PresencePrefix string `koanf:"presence_prefix"`
}
type Repo interface {
	Upsert(ctx context.Context, key string, timestamp int64) error
}
type Service struct {
	config Config
	repo   Repo
}

func New(config Config, repo Repo) Service {
	return Service{repo: repo, config: config}
}

func (s Service) UpsertPresence(ctx context.Context, req param.UpsertPresenceRequest) (param.UpsertPresenceResponse, error) {
	const op = richerror.Op("presenceservice.UpsertPresence")
	err := s.repo.Upsert(ctx, fmt.Sprintf("%s:%d", s.config.PresencePrefix, req.UserID), req.Timestamp)
	if err != nil {
		return param.UpsertPresenceResponse{}, richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}
	return param.UpsertPresenceResponse{}, nil
}
