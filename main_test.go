package main

import (
	"flag"
	"fmt"
	"log"
	"sync"
	"testing"

	sdk "github.com/openshift-online/ocm-sdk-go"
)

var (
	logLevel           string
	token              string
	tokenURL           string
	clientID           string
	clientSecret       string
	url                string
	delayInMinutes     int8
	insecure           bool
	demoMode           bool
	concurrentRequests int
)

const orgSummaryDashboad = "org_summary"

func BenchmarkGetDashboard(b *testing.B) {
	flag.Parse()

	fmt.Printf("Staring benchmark against url: %s\n", url)

	logger, err := sdk.NewStdLoggerBuilder().
		Debug(true).
		Build()

	conn, err := sdk.NewConnectionBuilder().
		Logger(logger).
		Tokens(token).
		TokenURL(tokenURL).
		URL(url).
		Insecure(insecure).
		Build()

	if err != nil {
		log.Fatalf("An error occurred while attempting to create and SDK connection: %v",
			err)
	}

	wg := sync.WaitGroup{}

	for n := 0; n < concurrentRequests; n++ {
		go func() {
			wg.Add(1)
			if err := GetDashboard(conn); err != nil {
				b.Errorf("An error occurred getting dashboards: %v", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
