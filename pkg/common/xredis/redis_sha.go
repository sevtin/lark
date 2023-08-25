package xredis

// SCRIPT EXISTS e3f182128e86a9e10353874f247be4aaaf54f5aa
// EVALSHA e3f182128e86a9e10353874f247be4aaaf54f5aa 2 {1110018} 19
// EVALSHA eb98162fa13aa1ead4d609d64d8cb86e744ae765 2 {1110018} 19 20
const (
	// set_message_id.lua
	SHA_SET_MESSAGE_ID = "a88869032333fa66526b7d8dd7a7d0ad3a829002"
	// red_envelope_receive_2.1.lua
	SHA_DISTRIBUTION_RED_ENVELOPE = "7e4103d6569a822a646dda3f8819dc0d6eedc6d9" // 红包发放
	// red_envelope_receive_rollback_2.1.lua
	SHA_ROLLBACK_RED_ENVELOPE = "8c288ba1d7816636b06e2c72e7fce7029ce843b0" // 发放失败回滚

	/*
		// hdel_chat_member.lua
		SHA_HDEL_CHAT_MEMBER = "fffb39ae0256756cf33e4ed2b314099ec421eb34"

		// hmset_chat_member.lua 弃用
		SHA_HMSET_CHAT_MEMBER = "db83841ecc4948c9d2d3abab3f534035b8c881d9"

		// hmset_dist_chat_member.lua
		SHA_HMSET_DIST_CHAT_MEMBER = "713a888cfbadeee7f046ac4be05c4103e7c3b9b1"

		// mset_convo_message.lua
		SHA_MSET_CONVO_MESSAGE = "d7787bf09736f7d21e964892e183a90c1c3d2f7f"

		// mset_expire.lua
		//SHA_MSET_EXPIRE = "fbc4461da1b87a320220a2b9fdc5d043e232a851"

		// multiple_set_expire.lua {}
		SHA_MULTIPLE_SET_EXPIRE = "09567671e6667713ce33cd02371129e487a3893a"

		// zadd_conversation.lua
		SHA_ZADD_CONVERSATION = "162fc05f099fdbb0f1eb0b7350f27b81307581ef"
	*/
)
