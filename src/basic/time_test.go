package basic

import (
	"strconv"
	"testing"
	"time"
)

func TestTime(t *testing.T) {

	t.Log(time.Now().Unix())
	// 格式化数字，  两个参数， i:格式化的值  ,  base: 10进制
	t.Log(strconv.FormatInt(1e6,10))

	loadLocation, _ := time.LoadLocation("Asia/Shanghai")
	location, _ := time.ParseInLocation("2006-01-02 15:04:05", "2020-10-26 00:00:00", loadLocation)

	t.Log(location)
	t.Log(location.UnixNano() / int64(time.Millisecond))

}

