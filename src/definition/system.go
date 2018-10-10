package definition

type CloudNetSystemInformation struct {
	Kind  string                          `json:"kind"`
	Items []CloudNetSystemInformationItem `json:"items"`
}

type CloudNetSystemInformationItem struct {
	ChassisSerialNumber string `json:"chassisSerialNumber"`
	Platform            string `json:"platform"`
	Product             string `json:"product"`
}

// =====================

type CloudSysHostInfoStat struct {
	Kind  string `json:"kind"`
	Items []CloudSysHostInfoStatItem
}

type CloudSysHostInfoStatItem struct {
	HostID      string `json:"hostId"`
	MemoryTotal int    `json:"memoryTotal" metric_name:"system.memoryTotalInBytes" source_type:"gauge"`
	MemoryUsed  int    `json:"memoryUsed" metric_name:"system.memoryUsedInBytes" source_type:"gauge"`
}

// ======================

type SysCPU struct {
	// CPUs
	Entries map[string]SysCPUEntryValue
}

type SysCPUEntryValue struct {
	NestedStats struct {
		// CPU Infos
		Entries map[string]SysCPUNestedStatsEntryValue
	}
}

type SysCPUNestedStatsEntryValue struct {
	NestedStats struct {
		// CPU Cores
		Entries map[string]SysCPUSecondNestedStatsEntryValue
	}
}

type SysCPUSecondNestedStatsEntryValue struct {
	NestedStats struct {
		Entries struct {
			CpuID struct {
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
