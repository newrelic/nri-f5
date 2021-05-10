package entities

import (
	"errors"
	"regexp"
	"strings"

	"github.com/newrelic/infra-integrations-sdk/data/attribute"
	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/definition"
)

// CollectPoolMembers collects metrics and inventory for every member of a pool given its path
func CollectPoolMembers(fullPath string, i *integration.Integration, client *client.F5Client, hostPort string, args arguments.ArgumentList) {
	tildePath := strings.Replace(fullPath, "/", "~", -1) // f5 uses tildes in requests rather than slashes

	var memberStats definition.LtmPoolMemberStats
	if err := client.Request("/mgmt/tm/ltm/pool/"+tildePath+"/members/stats", &memberStats); err != nil {
		log.Error("Failed to collect inventory: %s", err)
	}

	if args.HasInventory() {
		populatePoolMembersInventory(memberStats, i, hostPort)
	}

	if args.HasMetrics() {
		populatePoolMembersMetrics(memberStats, i, client.BaseURL, hostPort)
	}
}

func populatePoolMembersInventory(memberStats definition.LtmPoolMemberStats, i *integration.Integration, hostPort string) {
	for poolURL, poolMember := range memberStats.Entries {
		entries := poolMember.NestedStats.Entries
		memberName, err := buildPoolMemberPath(poolURL)
		if err != nil {
			log.Error("Failed to parse pool name from url %s: %s", err)
			continue
		}
		poolmemberIDAttr := integration.NewIDAttribute("poolmember", memberName)
		entity, err := i.Entity(hostPort, "f5-poolmember", poolmemberIDAttr)
		if err != nil {
			log.Error("Failed to get entity for pool %s: %s", memberName, err.Error())
			continue
		}

		for k, v := range map[string]interface{}{
			"maxConnections": entries.MaximumConnections.Value,
			"monitorRule":    entries.MonitorRule.Description,
			"nodeName":       memberName,
			"poolName":       entries.PoolName.Description,
			"port":           entries.Port.Value,
			"kind":           poolMember.NestedStats.Kind,
		} {
			err := entity.SetInventoryItem(k, "value", v)
			if err != nil {
				log.Error("Failed to set inventory item %s: %s", k, err.Error())
			}
		}
	}
}

func populatePoolMembersMetrics(memberStats definition.LtmPoolMemberStats, i *integration.Integration, url string, hostPort string) {
	for poolURL, poolMember := range memberStats.Entries {
		entries := poolMember.NestedStats.Entries
		memberName, err := buildPoolMemberPath(poolURL)
		if err != nil {
			log.Error("Failed to parse pool name from url %s: %s", err)
			continue
		}
		poolmemberIDAttr := integration.NewIDAttribute("poolmember", memberName)
		entity, err := i.Entity(hostPort, "f5-poolmember", poolmemberIDAttr)
		if err != nil {
			log.Error("Failed to get entity for pool %s: %s", memberName, err.Error())
			continue
		}

		entries.AvailabilityState.ProcessedDescription = convertAvailabilityState(entries.AvailabilityState.Description)
		entries.EnabledState.ProcessedDescription = convertEnabledState(entries.EnabledState.Description)
		entries.SessionStatus.ProcessedDescription = convertSessionStatus(entries.SessionStatus.Description)
		entries.MonitorStatus.ProcessedDescription = convertMonitorStatus(entries.MonitorStatus.Description)
		dataIn := entries.DataIn.Value / 8
		dataOut := entries.DataIn.Value / 8
		entries.DataIn.ProcessedValue = &dataIn
		entries.DataOut.ProcessedValue = &dataOut

		ms := entity.NewMetricSet("F5BigIpPoolMemberSample",
			attribute.Attribute{Key: "displayName", Value: memberName},
			attribute.Attribute{Key: "entityName", Value: "poolmember:" + memberName},
			attribute.Attribute{Key: "poolName", Value: entries.PoolName.Description},
			attribute.Attribute{Key: "url", Value: url},
		)

		err = ms.MarshalMetrics(entries)
		if err != nil {
			log.Error("Failed to marshal metrics for pool %s: %s", memberName, err.Error())
			continue
		}
	}

}

func buildPoolMemberPath(url string) (string, error) {
	re := regexp.MustCompile("members/([^/]*)/stats")
	match := re.FindStringSubmatch(url)
	if len(match) != 2 {
		err := errors.New("failed to find a match for pool member path in the url string")
		return "", err
	}
	poolMemberPath := match[1]
	return replaceTildes(poolMemberPath), nil
}

func replaceTildes(match string) string {
	return strings.Replace(match, "~", "/", -1)
}
