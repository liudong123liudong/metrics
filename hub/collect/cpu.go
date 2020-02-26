package collect

import (
	"github.com/shirou/gopsutil/cpu"
)

type CpuBaseInfo struct {
	ModelName string  `json:"modelName"`
	Cores     int     `json:"cores"`
	Threads   int     `json:"threads"`
	GHz       float64 `json:"GHz"`
}

// 采集cpu线程数
func CollectCpuCors() (int, error) {
	return cpu.Counts(false)
}

// 采集cpu核数
func CollectCpuThreads() (int, error) {
	return cpu.Counts(true)
}

func CollectCpuBaseInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}

type CpuInfo struct {
	TimesStat   *cpu.TimesStat `json:"timesStat"`
	UsedPercent float64        `json:"usedPercent"`
}

// 采集cpu指标
func CollectCpuMetrics() (*CpuInfo, error) {
	var cpuInfo *CpuInfo

	cpuTimesStat, err := collectCpuTimes()
	if err != nil {
		return cpuInfo, err
	}

	cpuPercent, err := collectCpuPercent()
	if err != nil {
		return cpuInfo, err
	}

	cpuInfo = &CpuInfo{&cpuTimesStat[0], cpuPercent[0]}

	return cpuInfo, nil
}

// 采集cpu使用率
func collectCpuPercent() ([]float64, error) {
	return cpu.Percent(0, false)
}

// 采集cpu时间
func collectCpuTimes() ([]cpu.TimesStat, error) {
	return cpu.Times(false)
}
