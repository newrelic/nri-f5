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
	MemoryTotal int    `json:"memoryTotal"`
	MemoryUsed  int    `json:"memoryUsed"`
}

// ======================

type SysCPU struct {
	Entries []SysCPUEntryValue
}

type SysCPUEntryValue struct {
	NestedStats struct {
		Entries map[string]SysCPUNestedStatsEntryValue
	}
}

type SysCPUNestedStatsEntryValue struct {
	NestedStats struct {
		Entries map[string]SysCPUSecondNestedStatsEntryValue
	}
}

type SysCPUSecondNestedStatsEntryValue struct {
	NestedStats struct {
		Entries struct {
			CpuID struct {
				Value int
			} `json:"cpuId"`
			AverageCPUIdle struct {
				Value int
			} `json:"oneMinAvgIdle"`
			AverageCPUInterruptRequest struct {
				Value int
			} `json:"oneMinAvgIrq"`
			AverageCPUIoWait struct {
				Value int
			} `json:"oneMinAvgIowait"`
			AverageCPUNice struct {
				Value int
			} `json:"oneMinAvgNiced"`
			AverageCPUSoftirq struct {
				Value int
			} `json:"oneMinSoftirq"`
			AverageCPUStolen struct {
				Value int
			} `json:"oneMinStolen"`
			AverageCPUSystem struct {
				Value int
			} `json:"oneMinSystem"`
			AverageCPUUser struct {
				Value int
			} `json:"oneMinUser"`
			CPUIdleTicks struct {
				Value int
			} `json:"idle"`
			CPUSystemTicks struct {
				Value int
			} `json:"system"`
			CPUUserTicks struct {
				Value int
			} `json:"user"`
		}
	}
}
