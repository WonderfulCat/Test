package test_constant

const (
	STORE_INCREASE_CAPACITY_COUNT = 1  //可扩容次数
	STORE_INCREASE_CAPACITY       = 10 //一次扩容增加容量
	STORE_ITEM_STACK_MAX_NUM      = 5  //最大堆叠数量
	STORE_INIT_CAPACITY           = 30 //仓库初始大小

	ALLIANCE_MAX_MEMBERS = 50 //最大成员上限

	RES_REGISTER   = 2
	RES_OK         = 0
	RES_ERR        = -1
	RES_ERR_MSG_1  = "己扩容过仓库"
	RES_ERR_MSG_2  = "指定位置没有物品"
	RES_ERR_MSG_3  = "指定位置不存在"
	RES_ERR_MSG_4  = "未加入任何公会"
	RES_ERR_MSG_5  = "成员己达上限"
	RES_ERR_MSG_6  = "仓库无法放入此物品"
	RES_ERR_MSG_7  = "道具己放入仓库,放入数量%d,未放入数量%d"
	RES_ERR_MSG_8  = "密码错误"
	RES_ERR_MSG_9  = "公会信息读取错误"
	RES_ERR_MSG_10 = "以加入公会,不能重复加入"
	RES_ERR_MSG_11 = "以有公会,不能创建公会"
	RES_ERR_MSG_12 = "公会名称己被使用"
	RES_ERR_MSG_13 = "公会不存在"
	RES_ERR_MSG_14 = "未加入公会"
	RES_ERR_MSG_15 = "权限不足"
	RES_ERR_MSG_16 = "Json解析失败: %s"
	RES_ERR_MSG_17 = "未登陆,请先登陆"
	RES_ERR_MSG_18 = "%s 登陆成功"
	RES_ERR_MSG_19 = "成员列表:  %s"
	RES_ERR_MSG_20 = "公会 : %s 创建成功"
	RES_ERR_MSG_21 = "公会 : %s 解散成功"
	RES_ERR_MSG_22 = "仓库扩容成功,当前容量为 : %d"
	RES_ERR_MSG_23 = "物品列表 : %s"
	RES_ERR_MSG_24 = "指定格子 : %d 物品己销毁"
	RES_ERR_MSG_25 = "创建角色失败"
)

const (
	MEMBER_PERMISSION_ADMIN = 1 << iota
	MEMBER_PERMISSION_NORMAL
)

const (
	REGISTER_NAME_CHARACTER = "character"
	REGISTER_NAME_ALLIANCE  = "alliance"
	REGISTER_NAME_CACHE     = "cache"
)
