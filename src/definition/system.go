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

type SysCPUEntryValue struct {
	NestedStats struct {
	}
}
