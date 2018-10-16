package definition

// SysApplicationService is an unmarshalling struct
type SysApplicationService struct {
	Kind  string                      `json:"kind"`
	Items []SysApplicationServiceItem `json:"items"`
}

// SysApplicationServiceItem is an unmarshalling struct
type SysApplicationServiceItem struct {
	DeviceGroup      string                              `json:"deviceGroup"`
	Kind             string                              `json:"kind"`
	Name             string                              `json:"name"`
	Template         string                              `json:"template"`
	TemplateModified string                              `json:"templateModified"`
	TrafficGroup     string                              `json:"trafficGroup"`
	Variables        []SysApplicationServiceItemVariable `json:"variables"`
}

// SysApplicationServiceItemVariable is an unmarshalling struct
type SysApplicationServiceItemVariable struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
