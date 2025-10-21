package constant

type MultiKeyMode string

const (
	MultiKeyModeRandom       MultiKeyMode = "random"        // 随机
	MultiKeyModePolling      MultiKeyMode = "polling"       // 轮询
	MultiKeyModeSmartPolling MultiKeyMode = "smart_polling" // 智能轮询
)
