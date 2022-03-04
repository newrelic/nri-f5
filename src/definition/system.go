package definition

type DeviceSystemInfo interface {
	ChassisSerialNumber() string
	Platform() string
	Product() string
	NumItems() int
}

// CloudNetSystemInformation is an unmarshalling struct
type CloudNetSystemInformation struct {
	Kind  string                          `json:"kind"`
	Items []CloudNetSystemInformationItem `json:"items"`
}

// CloudNetSystemInformationItem is an unmarshalling struct
type CloudNetSystemInformationItem struct {
	ChassisSerialNumber string `json:"chassisSerialNumber"`
	Platform            string `json:"platform"`
	Product             string `json:"product"`
}

func (c CloudNetSystemInformation) ChassisSerialNumber() string {
	return c.Items[0].ChassisSerialNumber
}
func (c CloudNetSystemInformation) Platform() string {
	return c.Items[0].ChassisSerialNumber
}
func (c CloudNetSystemInformation) Product() string {
	return c.Items[0].ChassisSerialNumber
}
func (c CloudNetSystemInformation) NumItems() int {
	return len(c.Items)
}

type CMDevice struct {
	Kind  string         `json:"kind"`
	Items []CMDeviceItem `json:"items"`
}

// CMDeviceItem is an unmarshalling struct
type CMDeviceItem struct {
	ChassisSerialNumber string `json:"chassisId"`
	Platform            string `json:"platform"`
	Product             string `json:"product"`
}

func (c CMDevice) ChassisSerialNumber() string {
	return c.Items[0].ChassisSerialNumber
}
func (c CMDevice) Platform() string {
	return c.Items[0].ChassisSerialNumber
}
func (c CMDevice) Product() string {
	return c.Items[0].ChassisSerialNumber
}
func (c CMDevice) NumItems() int {
	return len(c.Items)
}

// =====================

// SysCPU is an unmarshalling struct
type SysCPU struct {
	// CPUs
	Entries map[string]SysCPUEntryValue
}

// SysCPUEntryValue is an unmarshalling struct
type SysCPUEntryValue struct {
	NestedStats struct {
		// CPU Infos
		Entries map[string]SysCPUNestedStatsEntryValue
	}
}

// SysCPUNestedStatsEntryValue is an unmarshalling struct
type SysCPUNestedStatsEntryValue struct {
	NestedStats struct {
		// CPU Cores
		Entries map[string]SysCPUSecondNestedStatsEntryValue
	}
}

// SysCPUSecondNestedStatsEntryValue is an unmarshalling struct
type SysCPUSecondNestedStatsEntryValue struct {
	NestedStats struct {
		Entries struct {
			CPUID struct {
				Value int
			} `json:"cpuId" metric_name:"system.cpuID" source_type:"attribute"`
			AverageCPUIdle struct {
				Value int `json:"value"`
			} `json:"oneMinAvgIdle"`
			AverageCPUInterruptRequest struct {
				Value int `json:"value"`
			} `json:"oneMinAvgIrq"`
			AverageCPUIoWait struct {
				Value int `json:"value"`
			} `json:"oneMinAvgIowait"`
			AverageCPUNice struct {
				Value int `json:"value"`
			} `json:"oneMinAvgNiced"`
			AverageCPUSoftirq struct {
				Value int `json:"value"`
			} `json:"oneMinSoftirq"`
			AverageCPUStolen struct {
				Value int `json:"value"`
			} `json:"oneMinStolen"`
			AverageCPUSystem struct {
				Value int `json:"value"`
			} `json:"oneMinSystem"`
			AverageCPUUser struct {
				Value int `json:"value"`
			} `json:"oneMinUser"`
			CPUIdleTicks struct {
				Value int `json:"value"`
			} `json:"idle"`
			CPUSystemTicks struct {
				Value int `json:"value"`
			} `json:"system"`
			CPUUserTicks struct {
				Value int `json:"value"`
			} `json:"user"`
		}
	}
}

// ProcessedCPUMetrics is an unmarshalling struct
type ProcessedCPUMetrics struct {
	AverageCPUIdle             *float64 `metric_name:"system.cpuIdleUtilization" source_type:"gauge"`
	AverageCPUInterruptRequest *float64 `metric_name:"system.cpuInterruptRequestUtilization" source_type:"gauge"`
	AverageCPUIoWait           *float64 `metric_name:"system.cpuIOWaitUtilization" source_type:"gauge"`
	AverageCPUNice             *float64 `metric_name:"system.cpuNiceLevelUtilization" source_type:"gauge"`
	AverageCPUSoftirq          *float64 `metric_name:"system.cpuSoftInterruptRequestUtilization" source_type:"gauge"`
	AverageCPUStolen           *float64 `metric_name:"system.cpuStolenUtilization" source_type:"gauge"`
	AverageCPUSystem           *float64 `metric_name:"system.cpuSystemUtilization" source_type:"gauge"`
	AverageCPUUser             *float64 `metric_name:"system.cpuUserUtilization" source_type:"gauge"`
	CPUIdleTicks               *float64 `metric_name:"system.cpuIdleTicksPerSecond" source_type:"rate"`
	CPUSystemTicks             *float64 `metric_name:"system.cpuSystemTicksPerSecond" source_type:"rate"`
	CPUUserTicks               *float64 `metric_name:"system.cpuUserTicksPerSecond" source_type:"rate"`
}

// ======================

// MemoryStatsList is an unmarshalling struct
type MemoryStatsList struct {
	Entries  map[string]MemoryTopLevelEntries `json:"entries,omitempty"`
	Kind     string                           `json:"kind,omitempty" pretty:"expanded"`
	SelfLink string                           `json:"selfLink,omitempty" pretty:"expanded"`
}

// MemoryTopLevelEntries is an unmarshalling struct
type MemoryTopLevelEntries struct {
	NestedStats MemoryInnerStatsList `json:"nestedStats,omitempty"`
}

// MemoryInnerStatsList is an unmarshalling struct
type MemoryInnerStatsList struct {
	Entries map[string]MemoryStatsEntries `json:"entries,omitempty"`
}

// MemoryStatsEntries is an unmarshalling struct
type MemoryStatsEntries struct {
	NestedStats MemoryStats `json:"nestedStats,omitempty"`
}

// MemoryStats is an unmarshalling struct
type MemoryStats struct {
	Entries struct {
		HostID struct {
			Description string `metric_name:"system.hostId" source_type:"attribute"`
		} `json:"hostId"`
		MemoryFree struct {
			Value int `json:"value" metric_name:"system.memoryFreeInBytes" source_type:"gauge"`
		} `json:"memoryFree,omitempty"`
		MemoryTotal struct {
			Value int `json:"value" metric_name:"system.memoryTotalInBytes" source_type:"gauge"`
		} `json:"memoryTotal,omitempty"`
		MemoryUsed struct {
			Value int `json:"value" metric_name:"system.memoryUsedInBytes" source_type:"gauge"`
		} `json:"memoryUsed,omitempty"`
		OtherMemoryFree struct {
			Value int `json:"value" metric_name:"system.otherMemoryFreeInBytes" source_type:"gauge"`
		} `json:"otherMemoryFree,omitempty"`
		OtherMemoryTotal struct {
			Value int `json:"value" metric_name:"system.otherMemoryTotalInBytes" source_type:"gauge"`
		} `json:"otherMemoryTotal,omitempty"`
		OtherMemoryUsed struct {
			Value int `json:"value" metric_name:"system.otherMemoryUsedInBytes" source_type:"gauge"`
		} `json:"otherMemoryUsed,omitempty"`
		SwapFree struct {
			Value int `json:"value" metric_name:"system.swapFreeInBytes" source_type:"gauge"`
		} `json:"swapFree,omitempty"`
		SwapTotal struct {
			Value int `json:"value" metric_name:"system.swapTotalInBytes" source_type:"gauge"`
		} `json:"swapTotal,omitempty"`
		SwapUsed struct {
			Value int `json:"value" metric_name:"system.swapUsedInBytes" source_type:"gauge"`
		} `json:"swapUsed,omitempty"`
		TmmMemoryFree struct {
			Value int `json:"value" metric_name:"system.tmmMemoryFreeInBytes" source_type:"gauge"`
		} `json:"tmmMemoryFree,omitempty"`
		TmmMemoryTotal struct {
			Value int `json:"value" metric_name:"system.tmmMemoryTotalInBytes" source_type:"gauge"`
		} `json:"tmmMemoryTotal,omitempty"`
		TmmMemoryUsed struct {
			Value int `json:"value" metric_name:"system.tmmMemoryUsedInBytes" source_type:"gauge"`
		} `json:"tmmMemoryUsed,omitempty"`
	} `json:"entries,omitempty"`
}
