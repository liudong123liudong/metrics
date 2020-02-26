package collect

import (
	"futong_server_agent_go/utils/agent_id"
	"futong_server_agent_go/utils/logger"
)

type HostMetrics struct {
	AgentId      string      `json:"agentId"`
	AgentVersion string      `json:"agentVersion"`
	Cpu          interface{} `json:"cpu"`
	Memory       interface{} `json:"memory"`
	Load         interface{} `json:"load"`
	Process      interface{} `json:"process"`
	Net          interface{} `json:"net"`
	Disk         interface{} `json:"disk"`
}

// 采集监控信息
func CollectHostMetrics() *HostMetrics {
	loadAvg, err := CollectAvgLoadMetrics()
	if err != nil {
		logger.Sugar.Errorf("collect loadAvg err:", err)
	}

	cpuInfo, err := CollectCpuMetrics()
	if err != nil {
		logger.Sugar.Errorf("collect cpu err:", err)
	}

	diskUsage, err := CollectDiskPartitionMetrics()
	if err != nil {
		logger.Sugar.Errorf("collect disk err:", err)
	}

	netStat, err := CollectNetMetrics()
	if err != nil {
		logger.Sugar.Errorf("collect net err:", err)
	}

	memStat, err := CollectMemoryMetrics()
	if err != nil {
		logger.Sugar.Errorf("collect memory err:", err)
	}

	processMetrics, err := CollectProcessMetrics()
	if err != nil {
		logger.Sugar.Errorf("collect process err:", err)
	}

	return &HostMetrics{
		AgentId:      agent_id.AgentId,
		//AgentId:      agent_id.GetAgentId(),
		AgentVersion: "1.0.0",
		Cpu:          cpuInfo,
		Memory:       memStat,
		Load:         loadAvg,
		Process:      processMetrics,
		Net:          netStat,
		Disk:         diskUsage,
	}
}
