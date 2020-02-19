package constant

// TODO: 服务类型
const (
	// 服务类型
	SERVER_TYPE = 9
)

// TODO: 注册命令
const (
	// 注册服务主命令
	MDM_SERVICE = 0x0001
	// 注册码服务子命令
	SUB_SERVICE_REGISTER = 0x0001
)

// TODO: 心跳命令
const (
	// 心跳主命令
	MDM_HEARTBEAT = 0x0000
	// 心跳子命令
	SUB_HEARTBEAT = 0x0000
)

// TODO: TWALLET服务业务命令
const (
	// TWALLET服务主命令
	MDM_TWALLET = 0x0301
	// 重载信息
	SUB_TWALLET_RELOAD = 0x0002
)
