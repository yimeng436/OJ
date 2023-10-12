package enum

const (
	SubmitWaiting = 0
	SubmitRunning = 1
	SubmitSucceed = 2
	SubmitFail    = 3
)

var StatusMap = make(map[int]string)

func init() {
	StatusMap[SubmitWaiting] = "排队中"
	StatusMap[SubmitRunning] = "提交中"
	StatusMap[SubmitSucceed] = "提交成功"
	StatusMap[SubmitFail] = "提交失败"
}

func GetValue(code int) string {
	status, e := StatusMap[code]
	if !e {
		return "错误状态"
	}
	return status
}
