package avatar

import (
	"context"
	"lark/pkg/proto/pb_avatar"
)

func (s *avatarServer) SetAvatar(ctx context.Context, req *pb_avatar.SetAvatarReq) (resp *pb_avatar.SetAvatarResp, err error) {
	return s.avatarService.SetAvatar(ctx, req)
}
