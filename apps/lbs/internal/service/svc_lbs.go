package service

import (
	"context"
	"lark/apps/lbs/internal/config"
	"lark/domain/cache"
	"lark/domain/mrepo"
	"lark/domain/repo"
	"lark/pkg/proto/pb_lbs"
)

type LbsService interface {
	ReportLngLat(ctx context.Context, req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp, err error)
	PeopleNearby(ctx context.Context, req *pb_lbs.PeopleNearbyReq) (resp *pb_lbs.PeopleNearbyResp, err error)
}

type lbsService struct {
	cfg       *config.Config
	lbsRepo   mrepo.LbsRepository
	userRepo  repo.UserRepository
	userCache cache.UserCache
}

func NewLbsService(cfg *config.Config, lbsRepo mrepo.LbsRepository, userRepo repo.UserRepository, userCache cache.UserCache) LbsService {
	return &lbsService{cfg: cfg, lbsRepo: lbsRepo, userRepo: userRepo, userCache: userCache}
}
