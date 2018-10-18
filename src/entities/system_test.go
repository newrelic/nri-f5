package entities

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/stretchr/testify/assert"
)

func TestCollectSystem(t *testing.T) {
	i, _ := integration.New("test", "test")

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)

		if req.URL.String() == "/mgmt/tm/cloud/sys/host-info-stat" {
			res.Write([]byte(`{
				"items": [{
					"hostId": "0",
					"memoryTotal": 33767403520,
					"memoryUsed": 2487071784,
					"cpuInfo": [{
						"cpuId": 0,
						"slotId": 0,
						"user": 11048598,
						"system": 1756964,
						"idle": 132086519,
						"usageRatio": 8,
						"generation": 0,
						"lastUpdateMicros": 0
					},
					{
						"cpuId": 1,
						"slotId": 0,
						"user": 14391014,
						"system": 2574597,
						"idle": 130385560,
						"usageRatio": 1,
						"generation": 0,
						"lastUpdateMicros": 0
					},
					{
						"cpuId": 2,
						"slotId": 0,
						"user": 5809280,
						"system": 898209,
						"idle": 133824218,
						"usageRatio": 4,
						"generation": 0,
						"lastUpdateMicros": 0
					},
					{
						"cpuId": 3,
						"slotId": 0,
						"user": 12011635,
						"system": 1917778,
						"idle": 133376554,
						"usageRatio": 2,
						"generation": 0,
						"lastUpdateMicros": 0
					},
					{
						"cpuId": 4,
						"slotId": 0,
						"user": 5808599,
						"system": 844854,
						"idle": 134058167,
						"usageRatio": 5,
						"generation": 0,
						"lastUpdateMicros": 0
					},
					{
						"cpuId": 5,
						"slotId": 0,
						"user": 8327610,
						"system": 1440918,
						"idle": 137528558,
						"usageRatio": 0,
						"generation": 0,
						"lastUpdateMicros": 0
					},
					{
						"cpuId": 6,
						"slotId": 0,
						"user": 6023923,
						"system": 868216,
						"idle": 134004306,
						"usageRatio": 4,
						"generation": 0,
						"lastUpdateMicros": 0
					},
					{
						"cpuId": 7,
						"slotId": 0,
						"user": 8970290,
						"system": 1344764,
						"idle": 137294351,
						"usageRatio": 0,
						"generation": 0,
						"lastUpdateMicros": 0
					}
					],
					"cpuCount": 8,
					"activeCpuCount": 4,
					"generation": 0,
					"lastUpdateMicros": 0
				}],
				"generation": 0,
				"lastUpdateMicros": 0,
				"kind": "tm:cloud:sys:host-info-stat:ltmhostinfostatcollectionstate",
				"selfLink": "https://localhost/mgmt/tm/cloud/sys/host-info-stat"
			}`))
		} else if req.URL.String() == "/mgmt/tm/sys/cpu" {
			res.Write([]byte(`{
				"kind": "tm:sys:cpu:cpucollectionstats",
				"selfLink": "https://localhost/mgmt/tm/sys/cpu?ver=12.1.1",
				"entries": {
					"https://localhost/mgmt/tm/sys/cpu/0": {
					"nestedStats": {
						"kind": "tm:sys:cpu:cpustats",
						"selfLink": "https://localhost/mgmt/tm/sys/cpu/0?ver=12.1.1",
						"entries": {
							"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo": {
							"nestedStats": {
								"entries": {
									"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/0": {
									"nestedStats": {
										"entries": {
											"cpuId": { "value": 0 },
											"fiveMinAvgIdle": { "value": 90 },
											"fiveMinAvgIowait": { "value": 0 },
											"fiveMinAvgIrq": { "value": 0 },
											"fiveMinAvgNiced": { "value": 0 },
											"fiveMinAvgSoftirq": { "value": 0 },
											"fiveMinAvgStolen": { "value": 0 },
											"fiveMinAvgSystem": { "value": 1 },
											"fiveMinAvgUser": { "value": 7 },
											"fiveSecAvgIdle": { "value": 90 },
											"fiveSecAvgIowait": { "value": 0 },
											"fiveSecAvgIrq": { "value": 1 },
											"fiveSecAvgNiced": { "value": 0 },
											"fiveSecAvgSoftirq": { "value": 0 },
											"fiveSecAvgStolen": { "value": 0 },
											"fiveSecAvgSystem": { "value": 1 },
											"fiveSecAvgUser": { "value": 7 },
											"idle": { "value": 132106949 },
											"iowait": { "value": 68495 },
											"irq": { "value": 365828 },
											"niced": { "value": 0 },
											"oneMinAvgIdle": { "value": 90 },
											"oneMinAvgIowait": { "value": 0 },
											"oneMinAvgIrq": { "value": 0 },
											"oneMinAvgNiced": { "value": 0 },
											"oneMinAvgSoftirq": { "value": 0 },
											"oneMinAvgStolen": { "value": 0 },
											"oneMinAvgSystem": { "value": 1 },
											"oneMinAvgUser": { "value": 8 },
											"softirq": { "value": 382551 },
											"stolen": { "value": 0 },
											"system": { "value": 1757214 },
											"usageRatio": { "value": 9 },
											"user": { "value": 11050291 }
										}
									}
								},
								"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/1": {
								"nestedStats": {
									"entries": {
										"cpuId": { "value": 1 },
										"fiveMinAvgIdle": { "value": 91 },
										"fiveMinAvgIowait": { "value": 0 },
										"fiveMinAvgIrq": { "value": 0 },
										"fiveMinAvgNiced": { "value": 0 },
										"fiveMinAvgSoftirq": { "value": 0 },
										"fiveMinAvgStolen": { "value": 0 },
										"fiveMinAvgSystem": { "value": 1 },
										"fiveMinAvgUser": { "value": 8 },
										"fiveSecAvgIdle": { "value": 99 },
										"fiveSecAvgIowait": { "value": 0 },
										"fiveSecAvgIrq": { "value": 0 },
										"fiveSecAvgNiced": { "value": 0 },
										"fiveSecAvgSoftirq": { "value": 0 },
										"fiveSecAvgStolen": { "value": 0 },
										"fiveSecAvgSystem": { "value": 0 },
										"fiveSecAvgUser": { "value": 1 },
										"idle": { "value": 130406327 },
										"iowait": { "value": 15794 },
										"irq": { "value": 0 },
										"niced": { "value": 130693 },
										"oneMinAvgIdle": { "value": 90 },
										"oneMinAvgIowait": { "value": 0 },
										"oneMinAvgIrq": { "value": 0 },
										"oneMinAvgNiced": { "value": 0 },
										"oneMinAvgSoftirq": { "value": 0 },
										"oneMinAvgStolen": { "value": 0 },
										"oneMinAvgSystem": { "value": 1 },
										"oneMinAvgUser": { "value": 9 },
										"softirq": { "value": 64520 },
										"stolen": { "value": 0 },
										"system": { "value": 2574916 },
										"usageRatio": { "value": 0 },
										"user": { "value": 14392649 }
									}
								}
							},
							"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/2": {
							"nestedStats": {
								"entries": {
									"cpuId": { "value": 2 },
									"fiveMinAvgIdle": { "value": 90 },
									"fiveMinAvgIowait": { "value": 0 },
									"fiveMinAvgIrq": { "value": 0 },
									"fiveMinAvgNiced": { "value": 0 },
									"fiveMinAvgSoftirq": { "value": 0 },
									"fiveMinAvgStolen": { "value": 0 },
									"fiveMinAvgSystem": { "value": 1 },
									"fiveMinAvgUser": { "value": 4 },
									"fiveSecAvgIdle": { "value": 90 },
									"fiveSecAvgIowait": { "value": 0 },
									"fiveSecAvgIrq": { "value": 0 },
									"fiveSecAvgNiced": { "value": 0 },
									"fiveSecAvgSoftirq": { "value": 0 },
									"fiveSecAvgStolen": { "value": 0 },
									"fiveSecAvgSystem": { "value": 1 },
									"fiveSecAvgUser": { "value": 4 },
									"idle": { "value": 133844803 },
									"iowait": { "value": 1839 },
									"irq": { "value": 0 },
									"niced": { "value": 0 },
									"oneMinAvgIdle": { "value": 90 },
									"oneMinAvgIowait": { "value": 0 },
									"oneMinAvgIrq": { "value": 0 },
									"oneMinAvgNiced": { "value": 0 },
									"oneMinAvgSoftirq": { "value": 0 },
									"oneMinAvgStolen": { "value": 0 },
									"oneMinAvgSystem": { "value": 1 },
									"oneMinAvgUser": { "value": 4 },
									"softirq": { "value": 79 },
									"stolen": { "value": 0 },
									"system": { "value": 898346 },
									"usageRatio": { "value": 4 },
									"user": { "value": 5810191 }
								}
							}
						},
						"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/3": {
						"nestedStats": {
							"entries": {
								"cpuId": { "value": 3 },
								"fiveMinAvgIdle": { "value": 91 },
								"fiveMinAvgIowait": { "value": 0 },
								"fiveMinAvgIrq": { "value": 0 },
								"fiveMinAvgNiced": { "value": 0 },
								"fiveMinAvgSoftirq": { "value": 0 },
								"fiveMinAvgStolen": { "value": 0 },
								"fiveMinAvgSystem": { "value": 1 },
								"fiveMinAvgUser": { "value": 7 },
								"fiveSecAvgIdle": { "value": 90 },
								"fiveSecAvgIowait": { "value": 0 },
								"fiveSecAvgIrq": { "value": 0 },
								"fiveSecAvgNiced": { "value": 0 },
								"fiveSecAvgSoftirq": { "value": 0 },
								"fiveSecAvgStolen": { "value": 0 },
								"fiveSecAvgSystem": { "value": 1 },
								"fiveSecAvgUser": { "value": 9 },
								"idle": { "value": 133397472 },
								"iowait": { "value": 9748 },
								"irq": { "value": 0 },
								"niced": { "value": 147696 },
								"oneMinAvgIdle": { "value": 91 },
								"oneMinAvgIowait": { "value": 0 },
								"oneMinAvgIrq": { "value": 0 },
								"oneMinAvgNiced": { "value": 0 },
								"oneMinAvgSoftirq": { "value": 0 },
								"oneMinAvgStolen": { "value": 0 },
								"oneMinAvgSystem": { "value": 1 },
								"oneMinAvgUser": { "value": 8 },
								"softirq": { "value": 46325 },
								"stolen": { "value": 0 },
								"system": { "value": 1918014 },
								"usageRatio": { "value": 2 },
								"user": { "value": 12013135 }
							}
						}
					},
					"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/4": {
					"nestedStats": {
						"entries": {
							"cpuId": { "value": 4 },
							"fiveMinAvgIdle": { "value": 90 },
							"fiveMinAvgIowait": { "value": 0 },
							"fiveMinAvgIrq": { "value": 0 },
							"fiveMinAvgNiced": { "value": 0 },
							"fiveMinAvgSoftirq": { "value": 0 },
							"fiveMinAvgStolen": { "value": 0 },
							"fiveMinAvgSystem": { "value": 1 },
							"fiveMinAvgUser": { "value": 4 },
							"fiveSecAvgIdle": { "value": 90 },
							"fiveSecAvgIowait": { "value": 0 },
							"fiveSecAvgIrq": { "value": 0 },
							"fiveSecAvgNiced": { "value": 0 },
							"fiveSecAvgSoftirq": { "value": 0 },
							"fiveSecAvgStolen": { "value": 0 },
							"fiveSecAvgSystem": { "value": 0 },
							"fiveSecAvgUser": { "value": 4 },
							"idle": { "value": 134078770 },
							"iowait": { "value": 1041 },
							"irq": { "value": 0 },
							"niced": { "value": 0 },
							"oneMinAvgIdle": { "value": 91 },
							"oneMinAvgIowait": { "value": 0 },
							"oneMinAvgIrq": { "value": 0 },
							"oneMinAvgNiced": { "value": 0 },
							"oneMinAvgSoftirq": { "value": 0 },
							"oneMinAvgStolen": { "value": 0 },
							"oneMinAvgSystem": { "value": 1 },
							"oneMinAvgUser": { "value": 4 },
							"softirq": { "value": 0 },
							"stolen": { "value": 0 },
							"system": { "value": 844983 },
							"usageRatio": { "value": 4 },
							"user": { "value": 5809506 }
						}
					}
				},
				"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/5": {
				"nestedStats": {
					"entries": {
						"cpuId": { "value": 5 },
						"fiveMinAvgIdle": { "value": 95 },
						"fiveMinAvgIowait": { "value": 0 },
						"fiveMinAvgIrq": { "value": 0 },
						"fiveMinAvgNiced": { "value": 0 },
						"fiveMinAvgSoftirq": { "value": 0 },
						"fiveMinAvgStolen": { "value": 0 },
						"fiveMinAvgSystem": { "value": 1 },
						"fiveMinAvgUser": { "value": 4 },
						"fiveSecAvgIdle": { "value": 100 },
						"fiveSecAvgIowait": { "value": 0 },
						"fiveSecAvgIrq": { "value": 0 },
						"fiveSecAvgNiced": { "value": 0 },
						"fiveSecAvgSoftirq": { "value": 0 },
						"fiveSecAvgStolen": { "value": 0 },
						"fiveSecAvgSystem": { "value": 0 },
						"fiveSecAvgUser": { "value": 0 },
						"idle": { "value": 137550285 },
						"iowait": { "value": 12692 },
						"irq": { "value": 0 },
						"niced": { "value": 200322 },
						"oneMinAvgIdle": { "value": 94 },
						"oneMinAvgIowait": { "value": 0 },
						"oneMinAvgIrq": { "value": 0 },
						"oneMinAvgNiced": { "value": 0 },
						"oneMinAvgSoftirq": { "value": 0 },
						"oneMinAvgStolen": { "value": 0 },
						"oneMinAvgSystem": { "value": 1 },
						"oneMinAvgUser": { "value": 5 },
						"softirq": { "value": 51804 },
						"stolen": { "value": 0 },
						"system": { "value": 1441082 },
						"usageRatio": { "value": 1 },
						"user": { "value": 8328420 }
					}
				}
			},
			"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/6": {
			"nestedStats": {
				"entries": {
					"cpuId": { "value": 6 },
					"fiveMinAvgIdle": { "value": 91 },
					"fiveMinAvgIowait": { "value": 0 },
					"fiveMinAvgIrq": { "value": 0 },
					"fiveMinAvgNiced": { "value": 0 },
					"fiveMinAvgSoftirq": { "value": 0 },
					"fiveMinAvgStolen": { "value": 0 },
					"fiveMinAvgSystem": { "value": 1 },
					"fiveMinAvgUser": { "value": 4 },
					"fiveSecAvgIdle": { "value": 90 },
					"fiveSecAvgIowait": { "value": 0 },
					"fiveSecAvgIrq": { "value": 0 },
					"fiveSecAvgNiced": { "value": 0 },
					"fiveSecAvgSoftirq": { "value": 0 },
					"fiveSecAvgStolen": { "value": 0 },
					"fiveSecAvgSystem": { "value": 1 },
					"fiveSecAvgUser": { "value": 4 },
					"idle": { "value": 134024932 },
					"iowait": { "value": 327 },
					"irq": { "value": 246 },
					"niced": { "value": 0 },
					"oneMinAvgIdle": { "value": 91 },
					"oneMinAvgIowait": { "value": 0 },
					"oneMinAvgIrq": { "value": 0 },
					"oneMinAvgNiced": { "value": 0 },
					"oneMinAvgSoftirq": { "value": 0 },
					"oneMinAvgStolen": { "value": 0 },
					"oneMinAvgSystem": { "value": 1 },
					"oneMinAvgUser": { "value": 4 },
					"softirq": { "value": 3352 },
					"stolen": { "value": 0 },
					"system": { "value": 868353 },
					"usageRatio": { "value": 5 },
					"user": { "value": 6024837 }
				}
			}
		},
		"https://localhost/mgmt/tm/sys/cpu/0/cpuInfo/7": {
		"nestedStats": {
			"entries": {
				"cpuId": { "value": 7 },
				"fiveMinAvgIdle": { "value": 95 },
				"fiveMinAvgIowait": { "value": 0 },
				"fiveMinAvgIrq": { "value": 0 },
				"fiveMinAvgNiced": { "value": 0 },
				"fiveMinAvgSoftirq": { "value": 0 },
				"fiveMinAvgStolen": { "value": 0 },
				"fiveMinAvgSystem": { "value": 1 },
				"fiveMinAvgUser": { "value": 4 },
				"fiveSecAvgIdle": { "value": 99 },
				"fiveSecAvgIowait": { "value": 0 },
				"fiveSecAvgIrq": { "value": 0 },
				"fiveSecAvgNiced": { "value": 0 },
				"fiveSecAvgSoftirq": { "value": 0 },
				"fiveSecAvgStolen": { "value": 0 },
				"fiveSecAvgSystem": { "value": 0 },
				"fiveSecAvgUser": { "value": 1 },
				"idle": { "value": 137316082 },
				"iowait": { "value": 10369 },
				"irq": { "value": 5 },
				"niced": { "value": 168551 },
				"oneMinAvgIdle": { "value": 94 },
				"oneMinAvgIowait": { "value": 0 },
				"oneMinAvgIrq": { "value": 0 },
				"oneMinAvgNiced": { "value": 0 },
				"oneMinAvgSoftirq": { "value": 0 },
				"oneMinAvgStolen": { "value": 0 },
				"oneMinAvgSystem": { "value": 1 },
				"oneMinAvgUser": { "value": 5 },
				"softirq": { "value": 133029 },
				"stolen": { "value": 0 },
				"system": { "value": 1344898 },
				"usageRatio": { "value": 1 },
				"user": { "value": 8971205 }
			} } } } } }, "hostId": { "description": "0" } } } } } }`))
		} else if req.URL.String() == "/mgmt/tm/cloud/net/system-information" {
			res.Write([]byte(`{
				"items": [{
					"chassisSerialNumber": "f5-njcu-trbd",
					"product": "5000",
					"platform": "C109",
					"generation": 0,
					"lastUpdateMicros": 0
				}],
				"generation": 0,
				"lastUpdateMicros": 0,
				"kind": "tm:cloud:net:system-information:syssysteminfocollectionstate",
				"selfLink": "https://localhost/mgmt/tm/cloud/net/system-information"
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

	wg.Add(1)
	CollectSystem(i, client, "testhost", &wg)
	wg.Wait()

	assert.Equal(t, 1, len(i.Entities))
	systemEntity, _ := i.Entity("testhost", "system")
	metrics := systemEntity.Metrics[0].Metrics
	assert.Equal(t, float64(2487071784), metrics["system.memoryUsedInBytes"])

	inventory := systemEntity.Inventory.Items()
	assert.Equal(t, "f5-njcu-trbd", inventory["chassisSerialNumber"]["value"])
}
