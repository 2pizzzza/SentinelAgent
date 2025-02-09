package metrics

type Metrics struct {
	CPUUsage    float64
	MemoryTotal uint64
	MemoryUsed  uint64
	MemoryFree  uint64
	SwapTotal   uint64
	SwapFree    uint64
	Disk        DiskStatus
}

type DiskStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
}
