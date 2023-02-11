package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"lark/pkg/common/xes"
	"lark/pkg/proto/pb_user"
	"lark/pkg/utils"
)

func (s *userService) SearchUser(ctx context.Context, req *pb_user.SearchUserReq) (resp *pb_user.SearchUserResp, _ error) {
	resp = &pb_user.SearchUserResp{List: make([]*pb_user.UserSummary, 0)}
	var (
		buf    bytes.Buffer
		err    error
		res    *esapi.Response
		client *elasticsearch.Client
		r      map[string]interface{}
		i      int
		hit    interface{}
		total  interface{}
	)
	query := map[string]interface{}{
		"size": req.Size,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": map[string]interface{}{
					"term": map[string]interface{}{
						"lark_name": req.Query,
					},
				},
				"filter": map[string]interface{}{
					"range": map[string]interface{}{
						"uid": map[string]interface{}{
							"gt": req.LastUid,
						},
					},
				},
			},
		},
		"sort": []map[string]interface{}{
			{"uid": "asc"},
		},
	}

	if err = json.NewEncoder(&buf).Encode(query); err != nil {
		return
	}
	client = xes.GetClient()
	res, err = client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("es_users"),
		client.Search.WithBody(&buf),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.IsError() {
		return
	}
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return
	}
	resp.List = make([]*pb_user.UserSummary, len(r["hits"].(map[string]interface{})["hits"].([]interface{})))
	total = r["hits"].(map[string]interface{})["total"]
	resp.Total, _ = utils.ToInt64(total.(map[string]interface{})["value"])

	for i, hit = range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		var msg = &pb_user.UserSummary{}
		utils.Copy(hit.(map[string]interface{})["_source"], msg)
		resp.List[i] = msg
	}
	return
}
