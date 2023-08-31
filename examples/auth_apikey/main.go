package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/minchao/neuvector-sdk/auth"
	apiclient "github.com/minchao/neuvector-sdk/client"
	logapi "github.com/minchao/neuvector-sdk/client/log"
)

var (
	host   string
	apiKey string
)

func init() {
	flag.StringVar(&host, "host", "localhost:10443", "NeuVector Controller API host and port")
	flag.StringVar(&apiKey, "api-key", "", "API Key")
}

// Example of using API Key authentication.
//
// Usage: `go run main.go --host <Host> --api-key <API_Key>`
func main() {
	flag.Parse()

	tp := httptransport.New(host, "/", []string{"https"})
	tp.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := apiclient.New(tp, nil)

	// New an auth apiKey writer
	authInfo := auth.NewApiKeyWriter(apiKey)

	// Call the LogEvent API
	resp, err := client.Log.GetV1LogEvent(&logapi.GetV1LogEventParams{
		Context: context.Background(),
	}, authInfo)
	handleError(err)

	// Print the response payload
	payload, err := resp.Payload.MarshalBinary()
	handleError(err)
	fmt.Println(string(payload))
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
