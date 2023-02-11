package do

type ChatInvite struct {
	InviteId      int64          `gorm:"column:invite_id" json:"invite_id"`                             // invite ID
	ChatId        int64          `gorm:"column:chat_id" json:"chat_id"`                                 // Chat ID
	ChatType      int            `gorm:"column:chat_type" json:"chat_type"`                             // 1:私聊/2:群聊
	InitiatorUid  int64          `gorm:"column:initiator_uid" json:"initiator_uid"`                     // 发起人 UID
	InviteeUid    int64          `gorm:"column:invitee_uid" json:"invitee_uid"`                         // 被邀请人 UID
	InvitationMsg string         `gorm:"column:invitation_msg" json:"invitation_msg"`                   // 邀请消息
	HandlerUid    int64          `gorm:"column:handler_uid" json:"handler_uid"`                         // 处理人 UID
	HandleResult  int            `gorm:"column:handle_result" json:"handle_result"`                     // 结果
	HandleMsg     string         `gorm:"column:handle_msg" json:"handle_msg"`                           // 处理消息
	HandledTs     int64          `gorm:"column:handled_ts" json:"handled_ts"`                           // 处理时间
	CreatedTs     int64          `gorm:"column:created_ts" json:"created_ts"`                           // 邀请时间
	InitiatorInfo *InitiatorInfo `json:"initiator_info" gorm:"foreignKey:uid;references:initiator_uid"` // 发起人信息
}

func (ChatInvite) TableName() string {
	return "chat_invites"
}

type InitiatorInfo struct {
	Uid       int64  `gorm:"column:uid" json:"uid"`               // 用户ID
	LarkId    string `gorm:"column:lark_id" json:"lark_id"`       // 账户ID
	Nickname  string `gorm:"column:nickname" json:"nickname"`     // 昵称
	Gender    int    `gorm:"column:gender" json:"gender"`         // 性别
	BirthTs   int64  `gorm:"column:birth_ts" json:"birth_ts"`     // 生日
	CityId    int    `gorm:"column:city_id" json:"city_id"`       // 城市ID
	AvatarKey string `gorm:"column:avatar_key" json:"avatar_key"` // 小图 72*72
}

func (InitiatorInfo) TableName() string {
	return "users"
}
