package entities

import (
	"strings"

	"github.com/newrelic/infra-integrations-sdk/data/metric"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectPoolMembers collects metrics and inventory for every member of a pool given its path
func CollectPoolMembers(fullPath string, i *integration.Integration, client *client.F5Client) {
	tildePath := strings.Replace(fullPath, "/", "~", -1) // f5 uses tildes in requests rather than slashes

	var memberStats definition.LtmPoolMemberStats
	if err := client.Request("/mgmt/tm/ltm/pool/"+tildePath+"/members/stats", &memberStats); err != nil {
		log.Error("Failed to collect inventory: %s", err)
	}

	populatePoolMembersInventory(memberStats, i)
	populatePoolMembersMetrics(memberStats, i)
}

func populatePoolMembersInventory(memberStats definition.LtmPoolMemberStats, i *integration.Integration) {
	for _, poolMember := range memberStats.Entries {
		entries := poolMember.NestedStats.Entries
		memberName := entries.NodeName.Description
		entity, err := i.Entity(memberName, "poolmember") // TODO verify that this is the correct name (and htat it's unique)
		if err != nil {
			log.Error("Failed to get entity for pool %s: %s", memberName, err.Error())
			continue
		}

		err = entity.SetInventoryItem("Maximum Connections", "value", entries.MaximumConnections.Value)
		err = entity.SetInventoryItem("Monitor Rule", "value", entries.MonitorRule.Description)
		err = entity.SetInventoryItem("Node Name", "value", memberName)
		err = entity.SetInventoryItem("Pool Name", "value", entries.PoolName.Description)
		err = entity.SetInventoryItem("Port", "value", entries.Port.Value)
		err = entity.SetInventoryItem("Kind", "value", poolMember.NestedStats.Kind)
	}
}

func populatePoolMembersMetrics(memberStats definition.LtmPoolMemberStats, i *integration.Integration) {
	for _, poolMember := range memberStats.Entries {
		entries := poolMember.NestedStats.Entries
		memberName := entries.NodeName.Description
		entity, err := i.Entity(memberName, "poolmember") // TODO verify that this is the correct name (and that it's unique)
		if err != nil {
			log.Error("Failed to get entity for pool %s: %s", memberName, err.Error())
			continue
		}

		entries.AvailabilityState.ProcessedDescription = convertAvailabilityState(entries.AvailabilityState.Description)
		entries.EnabledState.ProcessedDescription = convertEnabledState(entries.EnabledState.Description)
		entries.SessionStatus.ProcessedDescription = convertSessionStatus(entries.SessionStatus.Description)

		ms := entity.NewMetricSet("F5BigIpPoolMemberSample",
			metric.Attribute{Key: "displayName", Value: memberName},
			metric.Attribute{Key: "entityType", Value: "poolmember"},
			metric.Attribute{Key: "poolName", Value: entries.PoolName.Description},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to marshal metrics for pool %s: %s", memberName, err.Error())
			continue
		}
	}

}
