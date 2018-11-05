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

func TestCollectApplications(t *testing.T) {
	i, _ := integration.New("test", "test")

	testServer := httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)

		if req.URL.String() == "/mgmt/tm/sys/application/service" {
			res.Write([]byte(`{
				"kind": "tm:sys:application:service:servicecollectionstate",
				"selfLink": "https://localhost/mgmt/tm/sys/application/service?ver=12.1.1",
				"items": [{
					"kind": "tm:sys:application:service:servicestate",
					"name": "apache-0",
					"partition": "Common",
					"subPath": "apache-0.app",
					"fullPath": "/Common/apache-0.app/apache-0",
					"generation": 1,
					"selfLink": "https://localhost/mgmt/tm/sys/application/service/~Common~apache-0.app~apache-0?ver=12.1.1",
					"deviceGroup": "/Common/TestDeviceGroup",
					"deviceGroupReference": { "link": "https://localhost/mgmt/tm/cm/device-group/~Common~TestDeviceGroup?ver=12.1.1" },
					"inheritedDevicegroup": "true",
					"inheritedTrafficGroup": "true",
					"strictUpdates": "disabled",
					"template": "/Common/f5.http",
					"templateReference": { "link": "https://localhost/mgmt/tm/sys/application/template/~Common~f5.http?ver=12.1.1" },
					"templateModified": "no",
					"trafficGroup": "/Common/traffic-group-1",
					"trafficGroupReference": { "link": "https://localhost/mgmt/tm/cm/traffic-group/~Common~traffic-group-1?ver=12.1.1" },
					"tables": [
						{ "name": "basic__snatpool_members" },
						{ "name": "net__snatpool_members" },
						{ "name": "optimizations__hosts" },
						{ "name": "pool__hosts", "columnNames": [ "name" ],
							"rows": [{ "row": [ "bluemedora.localnet" ] }]
						},
						{ "name": "pool__members" },
						{ "name": "server_pools__servers" }
					],
					"variables": [
						{
							"name": "client__http_compression",
							"encrypted": "no",
							"value": "/#create_new#"
						},
						{
							"name": "net__client_mode",
							"encrypted": "no",
							"value": "lan"
						},
						{
							"name": "net__server_mode",
							"encrypted": "no",
							"value": "lan"
						},
						{
							"name": "pool__addr",
							"encrypted": "no",
							"value": "10.66.8.240"
						},
						{
							"name": "pool__pool_to_use",
							"encrypted": "no",
							"value": "/Common/QaTest"
						},
						{
							"name": "pool__port",
							"encrypted": "no",
							"value": "80"
						},
						{
							"name": "ssl__mode",
							"encrypted": "no",
							"value": "no_ssl"
						},
						{
							"name": "ssl_encryption_questions__advanced",
							"encrypted": "no",
							"value": "no"
						},
						{
							"name": "ssl_encryption_questions__help",
							"encrypted": "no",
							"value": "hide"
						}
					]
				}]
			}`))
		}
		//else if pattern := regexp.MustCompile(".*"); pattern.Match([]byte(req.URL.String())) {
		//res.Write([]byte())
		//}
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
	CollectApplications(i, client, &wg, partitionFilter)
	wg.Wait()

	assert.Equal(t, 1, len(i.Entities))
	applicationEntity, _ := i.Entity("/Common/apache-0.app/apache-0", "application")
	inventory := applicationEntity.Inventory.Items()
	assert.Equal(t, "/Common/QaTest", inventory["poolToUse"]["value"])
}
