package service

import (
	"context"
	"lark/domain/cr/cr_user"
	"lark/domain/po"
	"lark/pkg/proto/pb_lbs"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *lbsService) ReportLngLat(ctx context.Context, req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp, _ error) {
	resp = &pb_lbs.ReportLngLatResp{}
	var (
		user *pb_user.BasicUserInfo
		loc  *po.UserLocation
		err  error
	)
	user, err = cr_user.GetBasicUserInfo(s.userCache, s.userRepo, req.Uid)
	if err != nil {
		resp.Set(ERROR_CODE_LBS_QUERY_DB_FAILED, ERROR_LBS_QUERY_DB_FAILED)
		return
	}
	if user.Uid == 0 {
		resp.Set(ERROR_CODE_LBS_QUERY_DB_FAILED, ERROR_LBS_QUERY_DB_FAILED)
		return
	}
	loc = &po.UserLocation{
		Uid:       user.Uid,
		Longitude: req.Longitude,
		Latitude:  req.Latitude,
		OnlineTs:  utils.NowUnix(),
	}
	err = s.locRepo.Save(loc)
	if err != nil {
		resp.Set(ERROR_CODE_LBS_UPDATE_VALUE_FAILED, ERROR_LBS_UPDATE_VALUE_FAILED)
		return
	}
	return
}

/* 弃用 入库mysql通过flink同步到mongodb
func (s *lbsService) ReportLngLat(ctx context.Context, req *pb_lbs.ReportLngLatReq) (resp *pb_lbs.ReportLngLatResp, _ error) {
	resp = &pb_lbs.ReportLngLatResp{}
	var (
		user *pb_user.BasicUserInfo
		err  error
	)
	user, err = cr_user.GetBasicUserInfo(s.userCache, s.userRepo, req.Uid)
	if err != nil {
		resp.Set(ERROR_CODE_LBS_QUERY_DB_FAILED, ERROR_LBS_QUERY_DB_FAILED)
		return
	}
	if user.Uid == 0 {
		resp.Set(ERROR_CODE_LBS_QUERY_DB_FAILED, ERROR_LBS_QUERY_DB_FAILED)
		return
	}
	ul := &po.UserLocation{
		Uid:      req.Uid,
		Gender:   user.Gender,
		BirthTs:  user.BirthTs,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		OnlineTs: utils.NowUnix(),
		Location: &po.Location{
			Type:        "Point",
			Coordinates: []float64{req.Longitude, req.Latitude},
		},
	}
	err = s.lbsRepo.Upsert(ul)
	if err != nil {
		resp.Set(ERROR_CODE_LBS_UPDATE_VALUE_FAILED, ERROR_LBS_UPDATE_VALUE_FAILED)
		return
	}
	return
}
*/
