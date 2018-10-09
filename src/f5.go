package main

import (
	"os"
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/entities"
)

const (
	integrationName    = "com.newrelic.f5"
	integrationVersion = "0.1.0"
)

var (
	args arguments.ArgumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	exitOnErr(err)

	client, err := client.NewClient(&args)
	exitOnErr(err)

	collectEntities(i, client)

	exitOnErr(i.Publish())
}

func collectEntities(i *integration.Integration, client *client.F5Client) {
	// set up and run goroutines for each entity
	var wg sync.WaitGroup
	wg.Add(5)
	go entities.CollectSystem(i, client, &wg)
	go entities.CollectApplications(i, client, &wg)
	go entities.CollectVirtualServers(i, client, &wg)
	go entities.CollectNodes(i, client, &wg, args.NodeFilter)
	go entities.CollectPools(i, client, &wg, args.PoolMemberFilter)
	wg.Wait()
}

func exitOnErr(err error) {
	if err != nil {
		log.Error("Encountered fatal error: %v", err)
		os.Exit(1)
	}
}
