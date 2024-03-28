package functions

import (
	"math/rand"
	"time"
)

func Short() string {
	StringArray := "0123456789abcdefghijklmnopqrstuvwxyz" //定义一个字符串字集
	length := 4                                           //  short的长度固定为4(注意是切片已经经过了初始化)
	result := make([]byte, 4)                             // 声明一个字符串切片用来存储shor
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		result[i] = StringArray[rand.Intn(len(StringArray))]
	}
	return string(result)
}
