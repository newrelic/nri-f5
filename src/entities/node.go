package entities

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/v3/data/attribute"
	"github.com/newrelic/infra-integrations-sdk/v3/integration"
	"github.com/newrelic/infra-integrations-sdk/v3/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectNodes collects node entities from F5 and adds them to the integration
func CollectNodes(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup, pathFilter *arguments.PathMatcher, hostPort string, args arguments.ArgumentList) {
	defer wg.Done()

	if args.HasInventory() {
		var ltmNode definition.LtmNode
		if err := client.Request("/mgmt/tm/ltm/node", &ltmNode); err != nil {
			log.Error("Failed to collect inventory for nodes: %s", err.Error())
		}
		populateNodesInventory(i, ltmNode, pathFilter, hostPort)
	}

	if args.HasMetrics() {
		var ltmNodeStats definition.LtmNodeStats
		if err := client.Request("/mgmt/tm/ltm/node/stats", &ltmNodeStats); err != nil {
			log.Error("Failed to collect metrics for nodes: %s", err.Error())
		}
		populateNodesMetrics(i, ltmNodeStats, pathFilter, hostPort)
	}
}

func populateNodesInventory(i *integration.Integration, ltmNode definition.LtmNode, pathFilter *arguments.PathMatcher, hostPort string) {
	for _, node := range ltmNode.Items {
		if !pathFilter.Matches(node.FullPath) {
			continue
		}

		nodeIDAttr := integration.NewIDAttribute("node", node.FullPath)
		nodeEntity, err := i.EntityReportedVia(hostPort, hostPort, "f5-node", nodeIDAttr)
		if err != nil {
			log.Error("Failed to get entity object for node %s: %s", node.Name, err.Error())
		}

		for k, v := range map[string]interface{}{
			"fqdn":           node.FQDN.TMName,
			"kind":           node.Kind,
			"address":        node.Address,
			"maxConnections": node.MaxConnections,
			"monitorRule":    node.MonitorRule,
		} {
			err := nodeEntity.SetInventoryItem(k, "value", v)
			if err != nil {
				log.Error("Failed to set inventory item %s: %s", k, err.Error())
			}
		}
	}
}

func populateNodesMetrics(i *integration.Integration, ltmNodeStats definition.LtmNodeStats, pathFilter *arguments.PathMatcher, hostPort string) {
	for _, node := range ltmNodeStats.Entries {
		if !pathFilter.Matches(node.NestedStats.Entries.TmName.Description) {
			continue
		}

		entries := node.NestedStats.Entries
		nodeName := entries.TmName.Description
		nodeIDAttr := integration.NewIDAttribute("node", nodeName)
		nodeEntity, err := i.EntityReportedVia(hostPort, hostPort, "f5-node", nodeIDAttr)
		if err != nil {
			log.Error("Failed to get entity object for node %s: %s", nodeName, err.Error())
		}

		entries.MonitorStatus.ProcessedDescription = convertMonitorStatus(entries.MonitorStatus.Description)
		entries.EnabledState.ProcessedDescription = convertEnabledState(entries.EnabledState.Description)
		entries.SessionStatus.ProcessedDescription = convertSessionStatus(entries.SessionStatus.Description)
		entries.AvailabilityState.ProcessedDescription = convertAvailabilityState(entries.AvailabilityState.Description)
		dataIn := entries.DataIn.Value / 8
		dataOut := entries.DataIn.Value / 8
		entries.DataIn.ProcessedValue = &dataIn
		entries.DataOut.ProcessedValue = &dataOut

		ms := nodeEntity.NewMetricSet("F5BigIpNodeSample",
			attribute.Attribute{Key: "displayName", Value: nodeName},
			attribute.Attribute{Key: "entityName", Value: "node:" + nodeName},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to populate metrics for node %s: %s", nodeName, err)
		}
	}
}
