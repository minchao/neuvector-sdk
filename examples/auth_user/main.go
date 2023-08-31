package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/strfmt/conv"
	"github.com/go-openapi/swag"

	"github.com/minchao/neuvector-sdk/auth"
	apiclient "github.com/minchao/neuvector-sdk/client"
	"github.com/minchao/neuvector-sdk/client/authentication"
	logapi "github.com/minchao/neuvector-sdk/client/log"
	"github.com/minchao/neuvector-sdk/models"
)

var (
	host     string
	username string
	password string
)

func init() {
	flag.StringVar(&host, "host", "localhost:10443", "NeuVector Controller API host and port")
	flag.StringVar(&username, "username", "", "Username")
	flag.StringVar(&password, "password", "", "Password")
}

// Example of using username/password authentication.
//
// Usage: `go run main.go --host <Host> --username <Username> --password <Password>`
func main() {
	flag.Parse()

	tp := httptransport.New(host, "/", []string{"https"})
	tp.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := apiclient.New(tp, nil)

	// Login
	authResp, err := client.Authentication.PostV1Auth(&authentication.PostV1AuthParams{
		Body: &models.RESTAuthData{
			Password: &models.RESTAuthPassword{
				Username: swag.String(username),
				Password: conv.Password(strfmt.Password(password)),
			},
		},
		Context: context.Background(),
	})
	handleError(err)

	// New an auth token writer
	authInfo := auth.NewTokenWriter(*authResp.Payload.Token.Token)

	// Call the LogEvent API
	resp, err := client.Log.GetV1LogEvent(&logapi.GetV1LogEventParams{
		Context: context.Background(),
	}, authInfo)
	handleError(err)

	// Print the response payload
	payload, err := resp.Payload.MarshalBinary()
	handleError(err)
	fmt.Println(string(payload))

	// Logout
	//
	// Note: username based connections have a limited number of concurrent sessions,
	// so it is important to delete the username token when finished.
	// ref: https://open-docs.neuvector.com/automation/automation#authentication-for-rest-api
	_, err = client.Authentication.DeleteV1Auth(&authentication.DeleteV1AuthParams{
		Context: context.Background(),
	}, authInfo)
	handleError(err)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
