package collect

import (
	"github.com/shirou/gopsutil/mem"
)

type MemoryStat struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
	Free        uint64  `json:"free"`
	Buffers     uint64  `json:"buffers"`
	Cached      uint64  `json:"cached"`
	Shared      uint64  `json:"shared"`
}

// 采集mem
func CollectMemoryMetrics() (*MemoryStat, error) {
	m, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &MemoryStat{
		Total:       m.Total,
		Available:   m.Available,
		Used:        m.Used,
		UsedPercent: m.UsedPercent,
		Free:        m.Free,
		Buffers:     m.Buffers,
		Cached:      m.Cached,
		Shared:      m.Shared,
	}, nil
}
