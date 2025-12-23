// Pacakage main is the entry point of the application.
// Every executable Go program must have package main.
package main	

// Import required standard library packages.
// Each package has a single, clear responsibility.
import (
	// context is used to control cancellation and timeouts.
	// We use it to allow in-flight HTTP requests to finish
	// before shutting down the server.
	"context"


	// fmt provides formatted I/O functions.
	// Used here for logging messages to the console.
	"fmt"


	// net/http provides HTTP server and client implementions.
	// This is the core package for building web services in Go.
	"net/http"


	// os provides acces to operating system functionality.
	// Used here to receive OS-level signals (Crtl+c, SIGTERM).
	"os"


	// os/signal allows us to listen for operating system signals.
	// This is critical for graceful shutdowns in containers and kubernetes.
	"os/signal"

	// syscall provides low-level OS primitives.
	// SIGTERM is sent by Kubernetes before stopping a pod.
	"syscall"


	// time provides functionality for measuring and displaying time.
	// Used to simulate request processing and define shutdown timeouts.
	"time"

)

/*

helloHandler handles incoming HTTP requests.

Think of like this:
- A worker serving a customer
- The sleep simulates real work (DB calls, API calls, etc.)

This helps demonstrate how graceful shutdown waits
for ongoing requests to finish before exiting.
*/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request received")

	// Simulate long-running work
	time.Sleep(5 * time.Second)

	// Write HTTP response back to the client
	fmt.Fprintln(w, "Hello DevOps Graceful shutdown in action")

	fmt.Println("Request completed")
}

func main() {

	/*

		Create an HTTP request multiplexer (router).

		The mux decides which handler function 
		should handle which URL path.

	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	/*

		Configure the HTTP server.

		Port 909 is used insted of 8080
		(because Jenkins commonly uses 8080).

	*/
	server := &http.Server{
		Addr:	":9090",
		Handler: mux,
	}

	/* 

		Create a channel to receive OS signals.

		- os.Interrupt → Crtl +c (local development)
		- sysclall.SIGTERM → Kubernetes / Docker stop signal

	*/
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	/*

		Start the HTTP server in a separate goroutine.

		This allows the main goroutine to stay  free
		and listen for dhutdown signals.

	*/
	go func() {
		fmt.Println("HTTP server started on port 9090")

		// ListenAndServe bl;ocks until the server stops.
		// http.ErrServerClosed is returned on graceful shutdown,
		// which is expected and should NOT be teated as error.
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Server error: %v\n", err)
		}
}()

/* 

		Block the main goroutine until a shutdown dignal is reecived.
		This keeps the application running.

*/
<-stop
fmt.Println("\nShutdown signal received")

/*

		Create a context with timeout.

		This gives the server a fixed amount of time
		to finish processing ongoing requests.
*/
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()


/*
		Gracefully shut down the servser.

		This does three things:
		1. Stops accepting new requests
		2. Waits for active requests to finish
		3. Exits cleanly within the timeout
*/
if err := server.Shutdown(ctx); err != nil {
	fmt.Println("Forced shutdown:", err)
} else {
	fmt.Println("Server shut down gracefully")
	}

}

