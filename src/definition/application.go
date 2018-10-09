package definition

type SysApplicationService struct {
  Kind string `json:"kind"`
  Items []SysApplicationServiceItem `json:"items"`
}

type SysApplicationServiceItem struct {
  DeviceGroup string `json:"deviceGroup"`  
  Kind string `json:"kind"`
  Name string `json:"name"`
  Template string `json:"template"`
  TemplateModified string `json:"templateModified"`
  TrafficGroup string `json:"trafficGroup"`
}
