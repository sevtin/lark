package constant

type ACTION uint8

const (
	ACTION_CREATE ACTION = 1 // 创建
	ACTION_UPDATE ACTION = 2 // 更新
	ACTION_DELETE ACTION = 3 // 删除
	ACTION_QUERY  ACTION = 4 // 查询
)
