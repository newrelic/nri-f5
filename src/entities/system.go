package entities

import (
	"strings"
	"sync"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectSystem collects the system entity from F5 and adds it to the integration
func CollectSystem(integration *integration.Integration, client *client.F5Client, hostPort string, wg *sync.WaitGroup) {
	defer wg.Done()

	systemEntity, err := integration.Entity(hostPort, "system")
	if err != nil {
		log.Error("Couldn't create system entity: %v", err)
		return
	}
	systemMetrics := systemEntity.NewMetricSet("F5SystemSample",
		metric.Attribute{Key: "displayName", Value: systemEntity.Metadata.Name},
		metric.Attribute{Key: "entityName", Value: systemEntity.Metadata.Namespace + ":" + systemEntity.Metadata.Name},
	)

	var systemWg sync.WaitGroup
	systemWg.Add(3)
	go marshalSystemInfo(systemEntity, client, &systemWg)
	go marshalHostInfo(systemMetrics, client, &systemWg)
	go marshalCPUStats(systemMetrics, client, &systemWg)
	systemWg.Wait()
}

func marshalSystemInfo(systemEntity *integration.Entity, client *client.F5Client, wg *sync.WaitGroup) {
	defer wg.Done()

	var sysInfo definition.CloudNetSystemInformation
	endpoint := "/mgmt/tm/cloud/net/system-information"
	if err := client.Request(endpoint, &sysInfo); err != nil {
		log.Error("Couldn't get response from API for endpoint '%s': %v", endpoint, err)
		return
	}

	if len(sysInfo.Items) == 0 {
		log.Error("Couldn't get system information: no items returned from system endpoint")
		return
	}

	sysInfoItem := sysInfo.Items[0]

	for k, v := range map[string]interface{}{
		"chassisSerialNumber": sysInfoItem.ChassisSerialNumber,
		"platform":            sysInfoItem.Platform,
		"product":             sysInfoItem.Product,
	} {
		if err := systemEntity.SetInventoryItem(k, "value", v); err != nil {
			log.Error("Couldn't set inventory item '%s' on system entity: %v", k, err)
		}
	}
}

func marshalHostInfo(systemMetrics *metric.Set, client *client.F5Client, wg *sync.WaitGroup) {
	defer wg.Done()

	var hostInfo definition.CloudSysHostInfoStat
	endpoint := "/mgmt/tm/cloud/sys/host-info-stat"
	if err := client.Request(endpoint, &hostInfo); err != nil {
		log.Error("Couldn't get response from API for endpoint '%s': %v", endpoint, err)
		return
	}

	if len(hostInfo.Items) == 0 {
		log.Error("Couldn't get host info stats: no items returned from host info stat endpoint")
		return
	}

	if err := systemMetrics.MarshalMetrics(&hostInfo.Items[0]); err != nil {
		log.Error("Couldn't marshal system metrics from host info stat: %v", err)
	}
}

func marshalCPUStats(systemMetrics *metric.Set, client *client.F5Client, wg *sync.WaitGroup) {
	defer wg.Done()

	var cpuInfo definition.SysCPU
	endpoint := "/mgmt/tm/sys/cpu"
	if err := client.Request(endpoint, &cpuInfo); err != nil {
		log.Error("Couldn't get response from API for endpoint '%s': %v", endpoint, err)
		return
	}

	if len(cpuInfo.Entries) == 0 {
		log.Error("Couldn't get CPU stats: no entries returned in CPU stat response")
		return
	}

	processedCPU := definition.ProcessedCPUMetrics{
		AverageCPUIdle:             new(float64),
		AverageCPUInterruptRequest: new(float64),
		AverageCPUIoWait:           new(float64),
		AverageCPUNice:             new(float64),
		AverageCPUSoftirq:          new(float64),
		AverageCPUStolen:           new(float64),
		AverageCPUSystem:           new(float64),
		AverageCPUUser:             new(float64),
		CPUIdleTicks:               new(float64),
		CPUSystemTicks:             new(float64),
		CPUUserTicks:               new(float64),
	}
	cpuCounter := 0.0
	for cpuKey, cpu := range cpuInfo.Entries {
		log.Info("Looping %s", cpuKey)
		for cpuInfoKey, cpuInfo := range cpu.NestedStats.Entries {
			if !strings.HasSuffix(cpuInfoKey, "cpuInfo") {
				continue
			}
			for cpuCoreKey, cpuCore := range cpuInfo.NestedStats.Entries {
				log.Info("\tLooping %s", cpuCoreKey)
				// core stats, add to counters
				coreStats := cpuCore.NestedStats.Entries
				cpuCounter++
				*processedCPU.AverageCPUIdle += float64(coreStats.AverageCPUIdle.Value)
				*processedCPU.AverageCPUInterruptRequest += float64(coreStats.AverageCPUInterruptRequest.Value)
				*processedCPU.AverageCPUIoWait += float64(coreStats.AverageCPUIoWait.Value)
				*processedCPU.AverageCPUNice += float64(coreStats.AverageCPUNice.Value)
				*processedCPU.AverageCPUSoftirq += float64(coreStats.AverageCPUSoftirq.Value)
				*processedCPU.AverageCPUStolen += float64(coreStats.AverageCPUStolen.Value)
				*processedCPU.AverageCPUSystem += float64(coreStats.AverageCPUSystem.Value)
				*processedCPU.AverageCPUUser += float64(coreStats.AverageCPUUser.Value)
				*processedCPU.CPUIdleTicks += float64(coreStats.CPUIdleTicks.Value)
				*processedCPU.CPUSystemTicks += float64(coreStats.CPUSystemTicks.Value)
				*processedCPU.CPUUserTicks += float64(coreStats.CPUUserTicks.Value)
			}
		}
	}
	// take averages of utilization metrics
	*processedCPU.AverageCPUIdle /= cpuCounter
	*processedCPU.AverageCPUInterruptRequest /= cpuCounter
	*processedCPU.AverageCPUIoWait /= cpuCounter
	*processedCPU.AverageCPUNice /= cpuCounter
	*processedCPU.AverageCPUSoftirq /= cpuCounter
	*processedCPU.AverageCPUStolen /= cpuCounter
	*processedCPU.AverageCPUSystem /= cpuCounter
	*processedCPU.AverageCPUUser /= cpuCounter

	// marshal
	if err := systemMetrics.MarshalMetrics(&processedCPU); err != nil {
		log.Error("Couldn't marshal system CPU stats: %v", err)
	}
}
