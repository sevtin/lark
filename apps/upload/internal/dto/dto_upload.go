package dto

import (
	"lark/pkg/proto/pb_enum"
)

type UploadAvatarReq struct {
	OwnerType int32 `form:"owner_type" json:"owner_type" validate:"required,gte=1,lte=2"`
	OwnerId   int64 `form:"owner_id" json:"owner_id" validate:"omitempty,gt=0"`
}

type UploadAvatar struct {
	OwnerId      int64                `json:"owner_id"`
	OwnerType    pb_enum.AVATAR_OWNER `json:"owner_type"`
	AvatarSmall  string               `json:"avatar_small"`
	AvatarMedium string               `json:"avatar_medium"`
	AvatarLarge  string               `json:"avatar_large"`
}

type UploadPhotoResp struct {
	Small  string `json:"small"`  // 小图
	Medium string `json:"medium"` // 中图
	Large  string `json:"large"`  // 大图
	Origin string `json:"origin"` // 原始图
}

type PresignedReq struct {
	FileType string `form:"file_type" json:"file_type"`
}

type PresignedResp struct {
	Url string `form:"url" json:"url"`
}

type ObjectStorage struct {
	Bucket      string `json:"bucket"`
	Key         string `json:"key"`
	ETag        string `json:"e_tag"`
	Size        int64  `json:"size"`
	ContentType string `json:"content_type"`
	//Format      string `json:"format"`
	//UUID        string `json:"uuid"`
	FileName string `json:"file_name"`
	Tag      string `json:"tag"`
}
