package methods_interface

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"testing"
	"time"
)

// 测试定时任务
func TestCron(t *testing.T) {
	location, _ := time.LoadLocation("Asia/Shanghai")
	cron := cron.New(
		cron.WithSeconds(),          // 精确到秒
		cron.WithLocation(location)) // 设置时区为北京时间
	t.Log(cron.Location())

	cron.AddFunc("*/2 * * * * ?", func() {
		fmt.Println("zhenshi heheheh  ")
	})

	cron.Start()
	select {} // 保证不退出进程

}
