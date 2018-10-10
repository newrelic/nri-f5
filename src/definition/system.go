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
      } `json:"cpuId" metric_name:"system.cpuID" source_type:"attribute"`
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
      } `json:"oneMinAvgNiced" metric_name:"system.cpuNiceLevelUtilization" source_type:"gauge"`
			AverageCPUSoftirq struct {
				Value int
      } `json:"oneMinSoftirq" metric_name:"system.cpuSoftInterruptRequestUtilization" source_type:"gauge"`
			AverageCPUStolen struct {
				Value int
      } `json:"oneMinStolen" metric_name:"system.cpuStolenUtilization" source_type:"gauge"`
			AverageCPUSystem struct {
				Value int
      } `json:"oneMinSystem" metric_name:"system.cpuSystemUtilization" source_type:"gauge"`
			AverageCPUUser struct {
				Value int
      } `json:"oneMinUser" metric_name:"system.cpuUserUtilization" source_type:"gauge"`
			CPUIdleTicks struct {
				Value int
      } `json:"idle" metric_name:"system.cpuIdleTicksPerSecond" source_type:"rate"`
			CPUSystemTicks struct {
				Value int
      } `json:"system" metric_name:"system.cpuSystemTicksPerSecond" source_type:"rate"`
			CPUUserTicks struct {
				Value int
      } `json:"user" metric_name:"system.cpuUserTicksPerSecond" source_type:"rate"`
		}
	}
}
