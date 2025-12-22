package main 

import (
	"context"		// Used to control cancellation across goroutines
	"fmt"           // Used for printing messages
	"os"		    // Provides OS-level funtions
	"os/signal"		// Used to recieve OS signals (Ctrl+c, SIGTERM)
	"syscall"		// Provides system call constants like SIGTERM
	"time"			// Used for sleeping and time durations
)

func main() {
	fmt.Println("Starting application...")

	// ---------------------------------------------------------------
	// 1. Create a cancellable context
	// ---------------------------------------------------------------
	// context.Background() → root context (never cancelled by itself)
	// context.WithCancel() → creates a child context we can cancle manually
	// ctx      → shared signal passed to goroutines
	// cancel   → function used to trigger shutdown
	ctx, cancel := context.WithCancel(context.Background())


	// ----------------------------------------------------------------
	// 2. Create a channel to recieve OS shutdown  signals
	// ----------------------------------------------------------------
	// os.Signal → type of data the channel carries
	// buffer size = 1 → allows one signal witout blocking
	signalChan := make(chan os.Signal, 1)


	// ----------------------------------------------------------------
	// 3. Register signals to listen for
	// ----------------------------------------------------------------
	// os.Interrupt → Crtl + c (SGINT)
	// syscall.SIGTERM → sent by Docker/Kubernetes during shutdown
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)


	// ------------------------------------------------------------------
	// 4. Start application work in a separate goroutine
	// ------------------------------------------------------------------
	// This simulates a long-running service
	go runApp(ctx)


	// -------------------------------------------------------------------
	// 5. Block here until a shutdown signal is recieved
	// -------------------------------------------------------------------
	sig := <-signalChan
	fmt.Println("\nReceived signal:", sig)


	// --------------------------------------------------------------------
	// 6. Trigger graceful shutdown
	// --------------------------------------------------------------------
	// Calling a cancel() notifies all goroutines using ctx
	cancel()


	// ---------------------------------------------------------------------
	// 7. Give the application time to clean up
	// ---------------------------------------------------------------------
	fmt.Println("Shutting down gracefully...")
	time.Sleep(2 * time.Second)

	fmt.Println("Shutdown complete.")
}

// ---------------------------------------------------------------------------
// Application work function
// ---------------------------------------------------------------------------
func runApp(ctx context.Context) {
	for {
		select {
			
			// ----------------------------------------------------------------
			// Case 1: Context canceled → stop work  safely
			// ----------------------------------------------------------------
		case <-ctx.Done():
			// ctx.Done() is closed when cancel() is called
			fmt.Println("Stopping work...")
			return // Exit goroutine cleanly


		// ----------------------------------------------------------------------
		// Case 2: Default work continues
		// ----------------------------------------------------------------------
		default:
			fmt.Println("Working...")
			time.Sleep(1 *time.Second)
		}
	}
}