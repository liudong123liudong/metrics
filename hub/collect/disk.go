package collect

import (
	"github.com/shirou/gopsutil/disk"
)

// 采集disk分区使用情况指标
func CollectDiskPartitionMetrics() ([]*disk.UsageStat, error) {
	partitionStats, err := collectDiskPartitions()
	if err != nil {
		return nil, err
	}

	var diskUsageStatList []*disk.UsageStat

	for _, partition := range partitionStats {
		// 采集各分区的使用情况
		us, err :=  disk.Usage(partition.Mountpoint)
		if err != nil {
			return nil, err
		}

		diskUsageStatList = append(diskUsageStatList, us)
	}
	return diskUsageStatList, nil
}

// 采集disk分区 挂载点
func collectDiskPartitions() ([]disk.PartitionStat, error) {
	// Partitions returns disk partitions. If all is false, returns
	// physical devices only (e.g. hard disks, cd-rom drives, USB keys)
	// and ignore all others (e.g. memory partitions such as /dev/shm)
	return  disk.Partitions(false)
}
