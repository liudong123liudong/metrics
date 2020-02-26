package collect

import (
	"github.com/shirou/gopsutil/net"
	ipnet "github.com/toolkits/net"
	)

// 采集网络相关指标
func CollectNetMetrics() ([]net.IOCountersStat, error) {
	return net.IOCounters(false)
}

// 采集ipList
func CollectIPList() ([]string, error) {
	return ipnet.IntranetIP()
}
