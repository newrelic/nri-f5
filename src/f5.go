//go:generate goversioninfo
package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"github.com/newrelic/infra-integrations-sdk/integration"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/nri-f5/src/arguments"
	"github.com/newrelic/nri-f5/src/client"
	"github.com/newrelic/nri-f5/src/entities"
)

const (
	integrationName = "com.newrelic.f5"
)

var (
	args               arguments.ArgumentList
	integrationVersion = "0.0.0"
	gitCommit          = ""
	buildDate          = ""
)

func main() {
	// Create Integration
	i, err := integration.New(integrationName, integrationVersion, integration.Args(&args))
	exitOnErr(err)

	if args.ShowVersion {
		fmt.Printf(
			"New Relic %s integration Version: %s, Platform: %s, GoVersion: %s, GitCommit: %s, BuildDate: %s\n",
			strings.Title(strings.Replace(integrationName, "com.newrelic.", "", 1)),
			integrationVersion,
			fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
			runtime.Version(),
			gitCommit,
			buildDate)
		os.Exit(0)
	}

	log.SetupLogging(args.Verbose)

	pathFilter, err := args.Parse()
	exitOnErr(err)

	client, err := client.NewClient(&args)
	exitOnErr(err)

	err = client.LogIn()
	exitOnErr(err)
	defer exitOnErr(client.LogOut())

	collectEntities(i, client, pathFilter)

	exitOnErr(i.Publish())
}

func collectEntities(i *integration.Integration, client *client.F5Client, pathFilter *arguments.PathMatcher) {
	hostPort := args.Hostname + ":" + strconv.Itoa(args.Port)
	// set up and run goroutines for each entity
	var wg sync.WaitGroup
	wg.Add(5)
	go entities.CollectSystem(i, client, &wg, hostPort, args)
	go entities.CollectApplications(i, client, &wg, pathFilter, hostPort, args)
	go entities.CollectVirtualServers(i, client, &wg, pathFilter, hostPort, args)
	go entities.CollectPools(i, client, &wg, pathFilter, hostPort, args)
	go entities.CollectNodes(i, client, &wg, pathFilter, hostPort, args)
	wg.Wait()
}

func exitOnErr(err error) {
	if err != nil {
		log.Error("Encountered fatal error: %v", err)
		os.Exit(1)
	}
}
