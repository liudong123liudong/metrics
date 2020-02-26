package collect

import (
	"futong_server_agent_go/utils/agent_id"
	"futong_server_agent_go/utils/logger"
	"github.com/shirou/gopsutil/host"
)

type HostInfo struct {
	*host.InfoStat
	AgentId string      `json:"agentId"`
	Ips     []string    `json:"ips"`
	CpuBaseInfo
	Memory  uint64      `json:"memory"`
}

// 采集主机信息
// BootTime returns the system boot time expressed in seconds since the epoch.
func CollectHostInfo() *HostInfo {
	hostStat, err := host.Info()
	if err != nil {
		logger.Sugar.Errorf("collect hostStat err:", err)
	}

	ips, err := CollectIPList()
	if err != nil {
		logger.Sugar.Errorf("collect ips err:", err)
	}

	threads, err := CollectCpuThreads()
	if err != nil {
		logger.Sugar.Errorf("collect cpu threads err:", err)
	}
	cores, err := CollectCpuCors()
	if err != nil {
		logger.Sugar.Errorf("collect cpu cores err:", err)
	}

	cpuInfoStats, err := CollectCpuBaseInfo()
	if err != nil {
		logger.Sugar.Errorf("collect cpu infostats err:", err)
	}

	memoryStat, err := CollectMemoryMetrics()
	if err != nil {
		logger.Sugar.Errorf("collect memory info err:", err)
	}

	return &HostInfo{
		AgentId:  agent_id.AgentId,
		InfoStat: hostStat,
		Ips:      ips,
		Memory: memoryStat.Total,
		CpuBaseInfo: CpuBaseInfo{
			ModelName: cpuInfoStats[0].ModelName,
			Cores:     cores,
			Threads:   threads,
			GHz:       cpuInfoStats[0].Mhz / 1000,
		},
	}
}
