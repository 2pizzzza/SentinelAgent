package metrics

import (
	"fmt"
	"time"

	"github.com/prometheus/procfs"
)

type LinuxMetrics struct {
	procfs *procfs.FS
}

func NewLinuxMetrics() (*LinuxMetrics, error) {
	proc, err := procfs.NewFS("/proc")
	if err != nil {
		return nil, fmt.Errorf("failed to create procfs: %w", err)
	}
	return &LinuxMetrics{procfs: &proc}, nil
}

func (m *LinuxMetrics) Collect() (*Metrics, error) {
	metrics := &Metrics{}

	cpuUsage, err := m.CollectCPUUsage()
	if err != nil {
		return nil, err
	}
	metrics.CPUUsage = cpuUsage

	memStats, err := m.CollectMemoryStats()
	if err != nil {
		return nil, err
	}
	metrics.MemoryTotal = *memStats.MemTotal
	metrics.MemoryUsed = *memStats.MemTotal - *memStats.MemFree - *memStats.Buffers - *memStats.Cached
	metrics.MemoryFree = *memStats.MemFree
	metrics.SwapTotal = *memStats.SwapTotal
	metrics.SwapFree = *memStats.SwapFree

	return metrics, nil
}

func (m *LinuxMetrics) CollectCPUUsage() (float64, error) {

	stat1, err := m.procfs.Stat()
	if err != nil {
		return 0, err
	}

	time.Sleep(time.Second)

	stat2, err := m.procfs.Stat()
	if err != nil {
		return 0, err
	}

	cpu1 := stat1.CPUTotal
	cpu2 := stat2.CPUTotal

	total := (cpu2.User + cpu2.Nice + cpu2.System + cpu2.Idle) - (cpu1.User + cpu1.Nice + cpu1.System + cpu1.Idle)
	idle := cpu2.Idle - cpu1.Idle

	usage := 100 * (1 - float64(idle)/float64(total))
	return usage, nil
}

func (m *LinuxMetrics) CollectMemoryStats() (*procfs.Meminfo, error) {

	meminfo, err := m.procfs.Meminfo()
	if err != nil {
		return nil, err
	}

	return &meminfo, nil
}

func (m *LinuxMetrics) StartCollecting(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		metrics, err := m.Collect()
		if err != nil {
			fmt.Println("Error collecting metrics:", err)
			continue
		}

		fmt.Printf("Collected Metrics: %+v\n", metrics)
	}
}
