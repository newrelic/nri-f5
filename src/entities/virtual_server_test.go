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

func TestCollectVirtualServers(t *testing.T) {
	i, _ := integration.New("test", "test")

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)

		if req.URL.String() == "/mgmt/tm/ltm/virtual" {
			res.Write([]byte(`{
				"kind": "tm:ltm:virtual:virtualcollectionstate",
				"selfLink": "https://localhost/mgmt/tm/ltm/virtual?ver=12.1.1",
				"items": [{
					"kind": "tm:ltm:virtual:virtualstate",
					"name": "StevesListener",
					"partition": "Common",
					"fullPath": "/Common/StevesListener",
					"generation": 1589,
					"selfLink": "https://localhost/mgmt/tm/ltm/virtual/~Common~StevesListener?ver=12.1.1",
					"addressStatus": "yes",
					"autoLasthop": "default",
					"cmpEnabled": "yes",
					"connectionLimit": 7,
					"destination": "/Common/0.0.0.3:53",
					"enabled": true,
					"gtmScore": 0,
					"ipProtocol": "udp",
					"mask": "255.255.255.255",
					"mirror": "disabled",
					"mobileAppTunnel": "disabled",
					"nat64": "disabled",
					"pool": "/Common/StevesPool",
					"poolReference": {
						"link": "https://localhost/mgmt/tm/ltm/pool/~Common~StevesPool?ver=12.1.1"
					},
					"rateLimit": "disabled",
					"rateLimitDstMask": 0,
					"rateLimitMode": "object",
					"rateLimitSrcMask": 0,
					"serviceDownImmediateAction": "none",
					"source": "0.0.0.0/0",
					"sourceAddressTranslation": {
						"type": "none"
					},
					"sourcePort": "preserve",
					"synCookieStatus": "not-activated",
					"translateAddress": "disabled",
					"translatePort": "disabled",
					"vlansDisabled": true,
					"vsIndex": 12,
					"policiesReference": {
						"link": "https://localhost/mgmt/tm/ltm/virtual/~Common~StevesListener/policies?ver=12.1.1",
						"isSubcollection": true
					},
					"profilesReference": {
						"link": "https://localhost/mgmt/tm/ltm/virtual/~Common~StevesListener/profiles?ver=12.1.1",
						"isSubcollection": true
					}
				}]
			}`))
		} else if req.URL.String() == "/mgmt/tm/ltm/virtual/stats" {
			res.Write([]byte(`{
				"kind": "tm:ltm:virtual:virtualcollectionstats",
				"selfLink": "https://localhost/mgmt/tm/ltm/virtual/stats?ver=12.1.1",
				"entries": {
					"https://localhost/mgmt/tm/ltm/virtual/~Common~StevesListener/stats": {
					"nestedStats": {
						"kind": "tm:ltm:virtual:virtualstats",
						"selfLink": "https://localhost/mgmt/tm/ltm/virtual/~Common~StevesListener/stats?ver=12.1.1",
						"entries": {
							"actualPvaAccel": { "description": "none" },
							"clientside.bitsIn": { "value": 0 },
							"clientside.bitsOut": { "value": 0 },
							"clientside.curConns": { "value": 4 },
							"clientside.evictedConns": { "value": 0 },
							"clientside.maxConns": { "value": 0 },
							"clientside.pktsIn": { "value": 0 },
							"clientside.pktsOut": { "value": 0 },
							"clientside.slowKilled": { "value": 0 },
							"clientside.totConns": { "value": 0 },
							"cmpEnableMode": { "description": "all-cpus" },
							"cmpEnabled": { "description": "enabled" },
							"csMaxConnDur": { "value": 0 },
							"csMeanConnDur": { "value": 0 },
							"csMinConnDur": { "value": 0 },
							"destination": { "description": "0.0.0.3:53" },
							"ephemeral.bitsIn": { "value": 0 },
							"ephemeral.bitsOut": { "value": 0 },
							"ephemeral.curConns": { "value": 0 },
							"ephemeral.evictedConns": { "value": 0 },
							"ephemeral.maxConns": { "value": 0 },
							"ephemeral.pktsIn": { "value": 0 },
							"ephemeral.pktsOut": { "value": 0 },
							"ephemeral.slowKilled": { "value": 0 },
							"ephemeral.totConns": { "value": 0 },
							"fiveMinAvgUsageRatio": { "value": 0 },
							"fiveSecAvgUsageRatio": { "value": 0 },
							"tmName": { "description": "/Common/StevesListener" },
							"oneMinAvgUsageRatio": { "value": 0 },
							"status.availabilityState": { "description": "offline" },
							"status.enabledState": { "description": "enabled" },
							"status.statusReason": { "description": "The children pool member(s) are down" },
							"syncookieStatus": { "description": "not-activated" },
							"syncookie.accepts": { "value": 0 },
							"syncookie.hwAccepts": { "value": 0 },
							"syncookie.hwSyncookies": { "value": 0 },
							"syncookie.hwsyncookieInstance": { "value": 0 },
							"syncookie.rejects": { "value": 0 },
							"syncookie.swsyncookieInstance": { "value": 0 },
							"syncookie.syncacheCurr": { "value": 0 },
							"syncookie.syncacheOver": { "value": 0 },
							"syncookie.syncookies": { "value": 0 },
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
	CollectVirtualServers(i, client, &wg, partitionFilter)
	wg.Wait()

	assert.Equal(t, 1, len(i.Entities))
	virtualServerEntity, _ := i.Entity("/Common/StevesListener", "virtualServer")
	metrics := virtualServerEntity.Metrics[0].Metrics
	assert.Equal(t, float64(4), metrics["virtualserver.connections"])
	assert.Equal(t, float64(0), metrics["virtualserver.availabilityState"])
	assert.Equal(t, float64(1), metrics["virtualserver.enabled"])

	inventory := virtualServerEntity.Inventory.Items()
	assert.Equal(t, 7, inventory["maxConnections"]["value"])
}
