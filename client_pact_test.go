package main

import (
	"log"
	"testing"

	"github.com/pact-foundation/pact-go/dsl"
)

type data struct {
	Greeting string
}

func TestConsumer(t *testing.T) {

	// Create Pact connecting to local Daemon
	pact := &dsl.Pact{
		Consumer: "client",
		Provider: "server",
		Host:     "localhost",
	}
	defer pact.Teardown()

	// Pass in test case. This is the component that makes the external HTTP call
	var test = func() (err error) {
		_, err = MakeRequest("localhost", pact.Server.Port, "Robert")
		return err
	}

	// Set up our expected interactions.
	pact.
		AddInteraction().
		Given("Request for a user is made").
		UponReceiving("A request to get a users name").
		WithRequest(dsl.Request{
			Method:  "GET",
			Path:    dsl.String("/hi"),
			Query:   dsl.MapMatcher{"name": dsl.String("Robert")},
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.MapMatcher{"Greeting": dsl.String("Hello Robert")},
		})

	// Run the test, verify it did what we expected and capture the contract
	if err := pact.Verify(test); err != nil {
		log.Fatalf("Error on Verify: %v", err)
	}
}
