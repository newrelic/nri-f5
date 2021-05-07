package entities

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"sync"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/stretchr/testify/assert"
)

func TestCollectPools(t *testing.T) {
	i, _ := integration.New("test", "test")

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)

		if req.URL.String() == "/mgmt/tm/ltm/pool" {
			_, err := res.Write([]byte(`{
				"kind": "tm:ltm:pool:poolcollectionstate",
				"selfLink": "https://localhost/mgmt/tm/ltm/pool?ver=12.1.1",
				"items": [{
					"kind": "tm:ltm:pool:poolstate",
					"name": "CitrixPool",
					"partition": "Common",
					"fullPath": "/Common/CitrixPool",
					"generation": 1,
					"selfLink": "https://localhost/mgmt/tm/ltm/pool/~Common~CitrixPool?ver=12.1.1",
					"allowNat": "yes",
					"allowSnat": "yes",
					"ignorePersistedWeight": "disabled",
					"ipTosToClient": "pass-through",
					"ipTosToServer": "pass-through",
					"linkQosToClient": "pass-through",
					"linkQosToServer": "pass-through",
					"loadBalancingMode": "fastest-node",
					"minActiveMembers": 0,
					"minUpMembers": 0,
					"minUpMembersAction": "failover",
					"minUpMembersChecking": "disabled",
					"queueDepthLimit": 0,
					"queueOnConnectionLimit": "disabled",
					"queueTimeLimit": 0,
					"reselectTries": 0,
					"serviceDownAction": "none",
					"slowRampTime": 10,
					"membersReference": {
						"link": "https://localhost/mgmt/tm/ltm/pool/~Common~CitrixPool/members?ver=12.1.1",
						"isSubcollection": true
					}
				},{
					"kind": "tm:ltm:pool:poolstate",
					"name": "CitrixPosol",
					"partition": "Test",
					"fullPath": "/Test/CitrixPosol",
					"generation": 1,
					"selfLink": "https://localhost/mgmt/tm/ltm/pool/~Test~CitrixPosol?ver=12.1.1",
					"allowNat": "yes",
					"allowSnat": "yes",
					"ignorePersistedWeight": "disabled",
					"ipTosToClient": "pass-through",
					"ipTosToServer": "pass-through",
					"linkQosToClient": "pass-through",
					"linkQosToServer": "pass-through",
					"loadBalancingMode": "fastest-node",
					"minActiveMembers": 0,
					"minUpMembers": 0,
					"minUpMembersAction": "failover",
					"minUpMembersChecking": "disabled",
					"queueDepthLimit": 0,
					"queueOnConnectionLimit": "disabled",
					"queueTimeLimit": 0,
					"reselectTries": 0,
					"serviceDownAction": "none",
					"slowRampTime": 10,
					"membersReference": {
						"link": "https://localhost/mgmt/tm/ltm/pool/~Test~CitrixPosol/members?ver=12.1.1",
						"isSubcollection": true
					}
				}]
			}`))
			assert.NoError(t, err)
		} else if req.URL.String() == "/mgmt/tm/ltm/pool/stats" {
			_, err := res.Write([]byte(`{
				"kind": "tm:ltm:pool:poolcollectionstats",
				"selfLink": "https://localhost/mgmt/tm/ltm/pool/stats?ver=12.1.1",
				"entries": {
					"https://localhost/mgmt/tm/ltm/pool/~Common~CitrixPool/stats": {
						"nestedStats": {
							"kind": "tm:ltm:pool:poolstats",
							"selfLink": "https://localhost/mgmt/tm/ltm/pool/~Common~CitrixPool/stats?ver=12.1.1",
							"entries": {
								"activeMemberCnt": { "value": 3 },
								"connqAll.ageEdm": { "value": 0 },
								"connqAll.ageEma": { "value": 0 },
								"connqAll.ageHead": { "value": 0 },
								"connqAll.ageMax": { "value": 0 },
								"connqAll.depth": { "value": 0 },
								"connqAll.serviced": { "value": 0 },
								"connq.ageEdm": { "value": 0 },
								"connq.ageEma": { "value": 0 },
								"connq.ageHead": { "value": 0 },
								"connq.ageMax": { "value": 0 },
								"connq.depth": { "value": 0 },
								"connq.serviced": { "value": 0 },
								"curSessions": { "value": 0 },
								"minActiveMembers": { "value": 0 },
								"monitorRule": { "description": "none" },
								"tmName": { "description": "/Common/CitrixPool" },
								"serverside.bitsIn": { "value": 0 },
								"serverside.bitsOut": { "value": 0 },
								"serverside.curConns": { "value": 0 },
								"serverside.maxConns": { "value": 0 },
								"serverside.pktsIn": { "value": 0 },
								"serverside.pktsOut": { "value": 0 },
								"serverside.totConns": { "value": 0 },
								"status.availabilityState": { "description": "offline" },
								"status.enabledState": { "description": "enabled" },
								"status.statusReason": { "description": " " },
								"totRequests": { "value": 0 }
							}
						}
					}
				}
			}`))
			assert.NoError(t, err)
		} else if pattern := regexp.MustCompile(".*"); pattern.Match([]byte(req.URL.String())) {
			_, err := res.Write([]byte(`{
				"kind": "tm:ltm:pool:members:memberscollectionstats",
				"selfLink": "https://localhost/mgmt/tm/ltm/pool/~Common~CreatePoolNew/members/stats?ver=12.1.1",
				"entries": {
					"https://localhost/mgmt/tm/ltm/pool/~Common~CreatePoolNew/members/~Common~Pool123:80/stats": {
						"nestedStats": {
							"kind": "tm:ltm:pool:members:membersstats",
							"selfLink": "https://localhost/mgmt/tm/ltm/pool/~Common~CreatePoolNew/members/~Common~Pool123:80/stats?ver=12.1.1",
							"entries": {
								"addr": { "description": "0.0.0.227" },
								"connq.ageEdm": { "value": 0 },
								"connq.ageEma": { "value": 0 },
								"connq.ageHead": { "value": 0 },
								"connq.ageMax": { "value": 0 },
								"connq.depth": { "value": 0 },
								"connq.serviced": { "value": 0 },
								"curSessions": { "value": 0 },
								"monitorRule": { "description": "none" },
								"monitorStatus": { "description": "unchecked" },
								"nodeName": { "description": "/Common/Pool123" },
								"poolName": { "description": "/Common/CreatePoolNew" },
								"port": { "value": 80 },
								"serverside.bitsIn": { "value": 0 },
								"serverside.bitsOut": { "value": 0 },
								"serverside.curConns": { "value": 2 },
								"serverside.maxConns": { "value": 0 },
								"serverside.pktsIn": { "value": 0 },
								"serverside.pktsOut": { "value": 0 },
								"serverside.totConns": { "value": 0 },
								"sessionStatus": { "description": "enabled" },
								"status.availabilityState": { "description": "unknown" },
								"status.enabledState": { "description": "enabled" },
								"status.statusReason": { "description": "Pool member does not have service checking enabled" },
								"totRequests": { "value": 0 }
							}
						}
					}
				}
			}`))
			assert.NoError(t, err)
		}
	}))

	defer func() { testServer.Close() }()

	client := &client.F5Client{
		BaseURL:          testServer.URL,
		Username:         "testUser",
		Password:         "testPass",
		HTTPClient:       http.DefaultClient,
		AuthToken:        "asdfd",
		RequestSemaphore: make(chan struct{}, 1),
	}

	var wg sync.WaitGroup
	partitionFilter := &arguments.PathMatcher{Partitions: []string{"Common"}}

	wg.Add(1)
	CollectPools(i, client, &wg, partitionFilter, testServer.URL, arguments.ArgumentList{})
	wg.Wait()

	assert.Equal(t, 3, len(i.Entities))
	idattr := integration.NewIDAttribute("pool", "/Common/CitrixPool")
	poolEntity, _ := i.EntityReportedVia(testServer.URL, testServer.URL, "f5-pool", idattr)
	poolMetrics := poolEntity.Metrics[0].Metrics
	assert.Equal(t, "/Common/CitrixPool", poolMetrics["displayName"])
	assert.Equal(t, float64(3), poolMetrics["pool.activeMembers"])
	assert.Equal(t, float64(0), poolMetrics["pool.availabilityState"])
	assert.Equal(t, float64(1), poolMetrics["pool.enabled"])

	idattr = integration.NewIDAttribute("poolmember", "/Common/Pool123:80")
	memberEntity, _ := i.Entity(testServer.URL, "f5-poolmember", idattr)
	assert.Equal(t, 2, len(memberEntity.Metrics))
	memberMetrics := memberEntity.Metrics[0].Metrics
	assert.Equal(t, "/Common/Pool123:80", memberMetrics["displayName"])
	assert.Equal(t, float64(2), memberMetrics["member.connections"])
	assert.Equal(t, float64(1), memberMetrics["member.enabled"])
}

func TestBuildPoolMemberPathValidUrl(t *testing.T) {
	actual, _ := buildPoolMemberPath("https://localhost/mgmt/tm/ltm/pool/~Common-members/~Common~Pool123:80/stats?ver=12.1.1")
	expected := "/Common/Pool123:80"
	assert.Equal(t, expected, actual)
}

func TestBuildPoolMemberPathMultipleMembers(t *testing.T) {
	actual, _ := buildPoolMemberPath("https://localhost/mgmt/tm/ltm/pool/~Common-members/members/~Common~Pool123:80/stats?ver=12.1.1")
	expected := "/Common/Pool123:80"
	assert.Equal(t, expected, actual)
}

func TestBuildPoolMemberPathNoMatches(t *testing.T) {
	_, err := buildPoolMemberPath("https://localhost/mgmt/tm/ltm/pool/~Common-members?ver=12.1.1")
	assert.Error(t, err)
}
func TestBuildPoolMemberPathEmptyUrl(t *testing.T) {
	_, err := buildPoolMemberPath("")
	assert.Error(t, err)
}

func TestReplaceTildes(t *testing.T) {
	test := "~Common~Pool123:80"
	actual := replaceTildes(test)
	expected := "/Common/Pool123:80"
	assert.Equal(t, expected, actual)
}

func TestReplaceTildesNoTildes(t *testing.T) {
	test := "CommonPool123:80"
	actual := replaceTildes(test)
	expected := "CommonPool123:80"
	assert.Equal(t, expected, actual)
}

func TestReplaceTildesEmptyString(t *testing.T) {
	test := ""
	actual := replaceTildes(test)
	expected := ""
	assert.Equal(t, expected, actual)
}
