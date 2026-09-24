// Command cmd runs one agent example against the server in
// CONDUCTOR_SERVER_URL, or hosts the external workers example 33 calls.
//
//	go run ./agents/cmd                      # list the examples
//	go run ./agents/cmd 01_basic_agent       # run one
//	go run ./agents/cmd external-workers     # the services 33_external_workers calls
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"examples/agents"

	"github.com/conductor-sdk/conductor-go/sdk/ai"
	"github.com/conductor-sdk/conductor-go/sdk/client"
)

func main() {
	if len(os.Args) < 2 {
		for _, ex := range agents.Catalog {
			fmt.Println(ex.Name)
		}
		fmt.Println("external-workers")
		return
	}
	if os.Args[1] == "external-workers" {
		serveExternalWorkers()
		return
	}
	ex, ok := agents.Lookup(os.Args[1])
	if !ok {
		fmt.Fprintf(os.Stderr, "unknown example %q; run without arguments for the list\n", os.Args[1])
		os.Exit(2)
	}
	rt := ai.NewRuntime(ai.Config{})
	defer rt.Shutdown()
	if _, err := ex.Run(context.Background(), rt, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "run failed:", err)
		os.Exit(1)
	}
}

func serveExternalWorkers() {
	stop, err := agents.StartExternalWorkers(client.NewAPIClientFromEnv())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("external workers polling; Ctrl-C to stop")
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
	stop()
}
