package collect

import (
	"github.com/shirou/gopsutil/load"
)

// 采集系统平均负载 win不支持此参数
func CollectAvgLoadMetrics() (*load.AvgStat, error) {
	return load.Avg()
}
