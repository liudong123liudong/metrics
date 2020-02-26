package collect

import (
	"github.com/shirou/gopsutil/process"
)

type ProcessMetrics struct {
	Total       int            `json:"total"`
	Pids        []int32        `json:"pids"`
	ProcessList []*ProcessInfo `json:"processList"`
}

type ProcessInfo struct {
	Pid           int32   `json:"pid"`
	Name          string  `json:"name"`
	Status        string  `json:"status"`
	CpuPercent    float64 `json:"cpuPercent"`
	MemoryPercent float32 `json:"memoryPercent"`
	CreateTime    int64   `json:"createTime"`
	OpenFiles     int     `json:"openFiles"`
	//IoCounters    *process.IOCountersStat `json:"ioCounters"`
}

// 采集进程相关指标
func CollectProcessMetrics() (*ProcessMetrics, error) {
	pids, err := process.Pids()
	if err != nil {
		return nil, err
	}

	var processMetrics = &ProcessMetrics{
		Total:       0,
		Pids:        make([]int32, 0),
		ProcessList: []*ProcessInfo{},
	}

	for _, pid := range pids {
		p, err := process.NewProcess(pid)
		if err != nil {
			continue
		}

		cpuPercent, _ := p.CPUPercent()
		memPercent, _ := p.MemoryPercent()
		createTime, _ := p.CreateTime()
		//ioCounters, _ := p.IOCounters()

		// OpenFiles returns a slice of OpenFilesStat opend by the process.
		// OpenFilesStat includes a file path and file descriptor.
		openFiles, _ := p.OpenFiles()
		name, _ := p.Name()

		// Status returns the process status.
		// Return value could be one of these.
		// R: Running S: Sleep T: Stop I: Idle
		// Z: Zombie W: Wait L: Lock
		// The character is same within all supported platforms.
		status, _ := p.Status()

		processMetrics.ProcessList = append(processMetrics.ProcessList, &ProcessInfo{
			Pid:           pid,
			Name:          name,
			Status:        status,
			CpuPercent:    cpuPercent,
			MemoryPercent: memPercent,
			//IoCounters:    ioCounters,
			OpenFiles:  len(openFiles),
			CreateTime: createTime / 1000,
		})
	}

	processMetrics.Total = len(pids)
	processMetrics.Pids = pids

	return processMetrics, nil
}
