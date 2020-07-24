//go:generate goversioninfo
package main

import (
	"os"
	"strconv"
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/entities"
)

const (
	integrationName    = "com.newrelic.f5"
	integrationVersion = "2.2.0"
)

var (
	args arguments.ArgumentList
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	exitOnErr(err)

	if args.AuthHost == "" {
		args.AuthHost = args.Hostname
	}

	if args.AuthPort == -1 {
		args.AuthPort = args.Port
	}

	log.SetupLogging(args.Verbose)

	pathFilter, err := args.Parse()
	exitOnErr(err)

	client, err := client.NewClient(&args)
	exitOnErr(err)

	err = client.LogIn()
	exitOnErr(err)

	collectEntities(i, client, pathFilter)

	exitOnErr(i.Publish())
}

func collectEntities(i *integration.Integration, client *client.F5Client, pathFilter *arguments.PathMatcher) {
	hostPort := args.Hostname + ":" + strconv.Itoa(args.Port)
	// set up and run goroutines for each entity
	var wg sync.WaitGroup
	wg.Add(5)
	go entities.CollectSystem(i, client, &wg, hostPort)
	go entities.CollectApplications(i, client, &wg, pathFilter, hostPort)
	go entities.CollectVirtualServers(i, client, &wg, pathFilter, hostPort)
	go entities.CollectPools(i, client, &wg, pathFilter, hostPort)
	go entities.CollectNodes(i, client, &wg, pathFilter, hostPort)
	wg.Wait()
}

func exitOnErr(err error) {
	if err != nil {
		log.Error("Encountered fatal error: %v", err)
		os.Exit(1)
	}
}
