package service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"lark/domain/po"
	"lark/pkg/constant"
	"lark/pkg/entity"
	"lark/pkg/proto/pb_enum"
	"lark/pkg/proto/pb_lbs"
	"lark/pkg/utils"
	"time"
)

func (s *lbsService) PeopleNearby(ctx context.Context, req *pb_lbs.PeopleNearbyReq) (resp *pb_lbs.PeopleNearbyResp, _ error) {
	resp = &pb_lbs.PeopleNearbyResp{}
	var (
		locations []*po.UserLocation
		i         int
		loc       *po.UserLocation
		err       error
	)
	locations, err = s.getLocations(req, resp)
	if err != nil {
		resp.Set(ERROR_CODE_LBS_QUERY_DB_FAILED, ERROR_LBS_QUERY_DB_FAILED)
		return
	}
	resp.List = make([]*pb_lbs.UserLocation, len(locations))
	for i, loc = range locations {
		resp.List[i] = &pb_lbs.UserLocation{
			Uid:         loc.Uid,
			Gender:      pb_enum.GENDER(loc.Gender),
			Age:         int32(utils.CalculateAge(time.Unix(loc.BirthTs/1000, 0))),
			Nickname:    loc.Nickname,
			Avatar:      loc.Avatar,
			Distance:    int64(utils.GetDistance(loc.Location.Coordinates[0], loc.Location.Coordinates[1], req.Longitude, req.Latitude)),
			Coordinates: loc.Location.Coordinates,
		}
	}

	return
}

func (s *lbsService) getLocations(req *pb_lbs.PeopleNearbyReq, resp *pb_lbs.PeopleNearbyResp) (locations []*po.UserLocation, err error) {
	var (
		now         = time.Now()
		minTs       int64
		maxTs       int64
		minOnlineTs = now.UnixMilli() - constant.CONST_DURATION_LBS_QUERY_LAST_ONLINE_MILLISECOND
	)
	q := entity.NewMongoQuery()
	// q.SetSkip((req.Limit - 1) * req.Skip)
	q.SetLimit(req.Limit)
	q.SetFilter("online_ts", bson.M{"$gte": minOnlineTs})
	q.SetFilter("uid", bson.M{"$ne": req.Uid})
	if req.MinAge > 0 && req.MaxAge > 0 {
		minTs = now.AddDate(-int(req.MaxAge), 0, 0).UnixNano() / 1e6
		maxTs = now.AddDate(-int(req.MinAge), 0, 0).UnixNano() / 1e6
		q.SetFilter("birth_ts", bson.M{"$gte": minTs, "$lte": maxTs})
	}
	if req.Gender > 0 {
		q.SetFilter("gender", req.Gender)
	}
	// $near查询返回的数据，按距离由近到远排序。
	var location bson.M
	if req.Radius > 0 {
		location = bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{req.Longitude, req.Latitude},
				},
				"$maxDistance": req.Radius,              //最大距离
				"$minDistance": req.MinDistance + 0.001, //最小距离
			},
		}
	} else {
		location = bson.M{
			"$near": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{req.Longitude, req.Latitude},
				},
				"$minDistance": req.MinDistance + 0.001, //最小距离
			},
		}
	}
	q.SetFilter("location", location)
	//q.SetSort(bson.D{{"location", "2dsphere"}})
	locations, err = s.lbsRepo.UserLocations(q)
	if err != nil {
		resp.Set(ERROR_CODE_LBS_QUERY_DB_FAILED, ERROR_LBS_QUERY_DB_FAILED)
		return
	}
	return
}
