package enum

const (
	Accept              = "Accept"
	WrongAnswer         = "Wrong Answer"
	CompileError        = "CompileError"
	MemoryLimitExceeded = "MemoryLimitExceeded"
	TimeLimitExceeded   = "TimeLimitExceeded"
)

var judegeinfoMap = make(map[string]string)

func init() {
	judegeinfoMap[Accept] = "通过"
	judegeinfoMap[WrongAnswer] = "答案错误"
	judegeinfoMap[CompileError] = "编译错误"
	judegeinfoMap[MemoryLimitExceeded] = "内存溢出"
	judegeinfoMap[TimeLimitExceeded] = "超时"
}

func GetJudegeInfo(s string) string {
	res, e := judegeinfoMap[s]
	if !e {
		return ""
	}
	return res
}
