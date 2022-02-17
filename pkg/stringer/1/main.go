// https://zhuanlan.zhihu.com/p/358758940
package main

import "fmt"

// 定义错误码
const (
	ERR_CODE_OK             = 0 // OK
	ERR_CODE_INVALID_PARAMS = 1 // 无效参数
	ERR_CODE_TIMEOUT        = 2 // 超时
)

// 定义错误码与描述信息的映射
var mapErrDesc = map[int]string{
	ERR_CODE_OK:             "OK",
	ERR_CODE_INVALID_PARAMS: "无效参数",
	ERR_CODE_TIMEOUT:        "超时",
}

// 根据错误码返回描述信息
func GetDescription(errCode int) string {
	if desc, exist := mapErrDesc[errCode]; exist {
		return desc
	}
	return fmt.Sprintf("error code: %d", errCode)
}

func main()  {
	fmt.Println(GetDescription(ERR_CODE_OK))
}