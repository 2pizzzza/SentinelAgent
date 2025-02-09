package metrics

type Metrics struct {
	CPUUsage    float64
	MemoryTotal uint64
	MemoryUsed  uint64
	MemoryFree  uint64
	SwapTotal   uint64
	SwapFree    uint64
}
