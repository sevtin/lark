package service

import (
	"context"
	"lark/domain/cr/cr_user"
	"lark/domain/po"
	"lark/pkg/proto/pb_lbs"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

/*
db.user_locations.ensureIndex({ location: "2dsphere"});
db.user_locations.createIndex( { "uid": 1 }, { unique: true } )
db.user_locations.createIndex({"gender":1})
db.user_locations.createIndex({"birth_ts":1})
db.user_locations.createIndex({"online_ts":1})
*/
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
		OnlineTs: utils.NowMilli(),
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
