package entities

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/stretchr/testify/assert"
)

func TestCollectNodes(t *testing.T) {
	i, _ := integration.New("test", "test")

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)

		if req.URL.String() == "/mgmt/tm/ltm/node" {
			res.Write([]byte(`{
				"kind": "tm:ltm:node:nodecollectionstate",
				"selfLink": "https://localhost/mgmt/tm/ltm/node?ver=12.1.1",
				"items": [{
					"kind": "tm:ltm:node:nodestate",
					"name": "0.0.0.1",
					"partition": "Common",
					"fullPath": "/Common/0.0.0.1",
					"generation": 1,
					"selfLink": "https://localhost/mgmt/tm/ltm/node/~Common~0.0.0.1?ver=12.1.1",
					"address": "0.0.0.1",
					"connectionLimit": 7,
					"dynamicRatio": 1,
					"ephemeral": "false",
					"fqdn": {
						"addressFamily": "ipv4",
						"autopopulate": "disabled",
						"downInterval": 5,
						"interval": "3600"
					},
					"logging": "disabled",
					"monitor": "default",
					"rateLimit": "disabled",
					"ratio": 1,
					"session": "user-enabled",
					"state": "unchecked"
				}]
			}`))
		} else if req.URL.String() == "/mgmt/tm/ltm/node/stats" {
			res.Write([]byte(`{
					"kind": "tm:ltm:node:nodecollectionstats",
					"selfLink": "https://localhost/mgmt/tm/ltm/node/stats?ver=12.1.1",
					"entries": {
						"https://localhost/mgmt/tm/ltm/node/~Common~0.0.0.1/stats": {
						"nestedStats": {
							"kind": "tm:ltm:node:nodestats",
							"selfLink": "https://localhost/mgmt/tm/ltm/node/~Common~0.0.0.1/stats?ver=12.1.1",
							"entries": {
								"addr": { "description": "0.0.0.1" },
								"curSessions": { "value": 0 },
								"monitorRule": { "description": "none" },
								"monitorStatus": { "description": "unchecked" }, "tmName": { "description": "/Common/0.0.0.1" },
								"serverside.bitsIn": { "value": 0 },
								"serverside.bitsOut": { "value": 0 },
								"serverside.curConns": { "value": 3 },
								"serverside.maxConns": { "value": 4 },
								"serverside.pktsIn": { "value": 0 },
								"serverside.pktsOut": { "value": 0 },
								"serverside.totConns": { "value": 0 },
								"sessionStatus": { "description": "enabled" },
								"status.availabilityState": { "description": "unknown" },
								"status.enabledState": { "description": "enabled" },
								"status.statusReason": { "description": "Node address does not have service checking enabled" },
								"totRequests": { "value": 0 }
							}
						}
					}
				}
			}`))
		}
	}))

	defer func() { testServer.Close() }()

	client := &client.F5Client{
		BaseURL:    testServer.URL,
		Username:   "testUser",
		Password:   "testPass",
		HTTPClient: http.DefaultClient,
		AuthToken:  "asdfd",
	}

	var wg sync.WaitGroup
	partitionFilter := &arguments.PathMatcher{[]string{"Common"}}

	wg.Add(1)
	CollectNodes(i, client, &wg, partitionFilter, testServer.URL)
	wg.Wait()

	assert.Equal(t, 1, len(i.Entities))

	idattr := integration.NewIDAttribute("node", "/Common/0.0.0.1")
	nodeEntity, _ := i.EntityReportedVia(testServer.URL, testServer.URL, "f5-node", idattr)
	metrics := nodeEntity.Metrics[0].Metrics
	assert.Equal(t, float64(3), metrics["node.connections"])
	assert.Equal(t, float64(1), metrics["node.monitorStatus"])
	assert.Equal(t, float64(1), metrics["node.enabled"])
	assert.Equal(t, float64(1), metrics["node.sessionStatus"])

	inventory := nodeEntity.Inventory.Items()
	assert.Equal(t, int(7), inventory["maxConnections"]["value"])
}
