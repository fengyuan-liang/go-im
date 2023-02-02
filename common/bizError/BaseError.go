package bizError

import (
	"fmt"
	"time"
)

type BizErrorer interface {
	BizError() string
	ErrorMsg() string
	ErrorTime() time.Time
}

// BizError 自定义错误
type BizError struct {
	when time.Time
	what string
}

const FORMAT_DATA = "2006/01/02 15:04.000" // 标准格式化时间

// 绑定一个方法
func (bizError *BizError) Error() string {
	return fmt.Sprintf("异常时间:【%v】，异常提示:【%v】", bizError.when.Format(FORMAT_DATA), bizError.what)
}

// ErrorMsg 只返回错误信息本身
func (bizError *BizError) ErrorMsg() string {
	return bizError.what
}

func (bizError *BizError) ErrorTime() time.Time {
	return bizError.when
}

// BizError 绑定自定义异常进行区分
func (bizError *BizError) BizError() string {
	return bizError.Error()
}

// NewBizError 产生一个异常
func NewBizError(errMsg ...string) BizErrorer {
	return &BizError{when: time.Now(), what: func() string {
		var totalMsg = ""
		if len(errMsg) <= 1 {
			return errMsg[0]
		}
		for _, msg := range errMsg {
			totalMsg += msg
		}
		return totalMsg
	}()}
}
