package enum

const (
	Waiting = 0
	Running = 1
	Succeed = 2
	Fail    = 3
)

var StatusMap = make(map[int]string)

func init() {
	StatusMap[0] = "等待"
	StatusMap[1] = "运行中"
	StatusMap[2] = "成功"
	StatusMap[3] = "失败"
}

func GetValue(code int) string {
	status, e := StatusMap[code]
	if !e {
		return "错误状态"
	}
	return status
}
