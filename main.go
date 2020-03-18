package main

import (
	"flag"
	"net/http"

	sdk "github.com/openshift-online/ocm-sdk-go"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
)

func init() {
	flag.StringVar(
		&token,
		"token",
		"",
		"Token used to connect to the API.",
	)
	flag.StringVar(
		&tokenURL,
		"token-url",
		"https://sso.redhat.com/auth/realms/redhat-external/protocol/openid-connect/token",
		"Token URL used to connect to the API.",
	)
	flag.StringVar(
		&clientID,
		"client-id",
		"cloud-services",
		"Client identifier used to connect to the API.",
	)
	flag.StringVar(
		&clientSecret,
		"client-secret",
		"",
		"Client secret used to connect to the API.",
	)
	flag.StringVar(
		&url,
		"url",
		"https://api.openshift.com",
		"URL used to connect to the API.",
	)
	flag.BoolVar(
		&insecure,
		"insecure",
		false,
		"Whether to use an insecure connection.",
	)
	flag.IntVar(
		&concurrentRequests,
		"concurrent-requests",
		10,
		"The number of concurrent requests to make.",
	)
}

func GetDashboard(conn *sdk.Connection) error {
	response, err := conn.ClustersMgmt().V1().
		Dashboards().
		Dashboard(orgSummaryDashboad).
		Get().
		SendContext(context.Background())
	if err != nil {
		return err
	}
	if response.Status() != http.StatusOK {
		return errors.Errorf("Recieved unexpected status: %d", response.Status())
	}
	return nil
}
