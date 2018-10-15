package entities

import (
	"sync"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectNodes collects node entities from F5 and adds them to the integration
func CollectNodes(i *integration.Integration, client *client.F5Client, wg *sync.WaitGroup) {
	defer wg.Done()

	var ltmNode definition.LtmNode
	if err := client.Request("/mgmt/tm/ltm/node", &ltmNode); err != nil {
		log.Error("Failed to collect inventory for nodes: %s", err.Error())
	}

	var ltmNodeStats definition.LtmNodeStats
	if err := client.Request("/mgmt/tm/ltm/node/stats", &ltmNodeStats); err != nil {
		log.Error("Failed to collect metrics for nodes: %s", err.Error())
	}

	populateNodesInventory(i, ltmNode)
	populateNodesMetrics(i, ltmNodeStats)
}

func populateNodesInventory(i *integration.Integration, ltmNode definition.LtmNode) {
	for _, node := range ltmNode.Items {
		nodeEntity, err := i.Entity(node.FullPath, "node") // TODO ensure everywhere is using FullPath as node name
		if err != nil {
			log.Error("Failed to get entity object for node %s: %s", node.Name, err.Error())
		}

		// TODO handle errors
		err = nodeEntity.SetInventoryItem("FQDN", "value", node.FQDN.TMName)
		if err != nil {
			log.Error("Failed to set inventory item: %s", err.Error())
		}
		err = nodeEntity.SetInventoryItem("Kind", "value", node.Kind)
		err = nodeEntity.SetInventoryItem("IP Address", "value", node.Address)
		err = nodeEntity.SetInventoryItem("Maximum Connections", "value", node.MaxConnections)
		err = nodeEntity.SetInventoryItem("Monitor Rule", "value", node.MonitorRule)
	}
}

func populateNodesMetrics(i *integration.Integration, ltmNodeStats definition.LtmNodeStats) {
	for _, node := range ltmNodeStats.Entries {
		entries := node.NestedStats.Entries
		nodeName := entries.TmName.Description
		nodeEntity, err := i.Entity(nodeName, "node")
		if err != nil {
			log.Error("Failed to get entity object for node %s: %s", nodeName, err.Error())
		}

		ms := nodeEntity.NewMetricSet("F5BigIpNodeSample",
			metric.Attribute{Key: "displayName", Value: nodeName},
			metric.Attribute{Key: "entityType", Value: "node"},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to populate metrics for node %s: %s", nodeName, err)
		}
	}
}
