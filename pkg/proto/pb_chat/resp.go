package pb_chat

func (r *CreateGroupChatResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *EditGroupChatResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *QuitGroupChatResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *RemoveGroupChatMemberResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *DeleteContactResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *UploadAvatarResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *GetChatInfoResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}

func (r *GroupChatDetailsResp) Set(code int32, msg string) {
	r.Code = code
	r.Msg = msg
}
