package dto_avatar

type AvatarInfo struct {
	OwnerId      int64  `json:"owner_id"`
	OwnerType    int32  `json:"owner_type"`
	AvatarSmall  string `json:"avatar_small"`
	AvatarMedium string `json:"avatar_medium"`
	AvatarLarge  string `json:"avatar_large"`
}
