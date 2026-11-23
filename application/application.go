package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"slices"
	"syscall"
	"time"

	"github.com/haaragard/golang-stdlib-server/container"
	"github.com/haaragard/golang-stdlib-server/provider"
	"github.com/haaragard/golang-stdlib-server/server"
)

type HaaraApplication struct {
	container *container.HaaraContainer
	providers []provider.Provider
	server    *server.HaaraHttpServer
}

func NewHaaraApplication(server *server.HaaraHttpServer) *HaaraApplication {
	return &HaaraApplication{
		providers: make([]provider.Provider, 0),
		container: container.NewHaaraContainer(),
		server:    server,
	}
}

func (app *HaaraApplication) GetContainer() *container.HaaraContainer {
	return app.container
}

// Is a thread blocking feature.
func (app *HaaraApplication) Run() {
	// Get application context and OS signals
	context, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGTERM, // - sent to a process to request its termination. Most commonly used, and we’ll be focusing on it later.
		syscall.SIGKILL, // - “quit immediately”, can not be interfered with.
		syscall.SIGINT,  // - interrupt signal (such as Ctrl+C)
		syscall.SIGQUIT, // - quit signal (such as Ctrl+D)
	)
	defer stop()

	// Start providers
	app.startProviders()

	// Dispatch server ServeAndListen to GoRoutine.
	go app.server.ServeAndListen()

	<-context.Done()
	app.Stop(context)

	// Ends program
	os.Exit(0)
}

func (app *HaaraApplication) Stop(ctx context.Context) {
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("\nStopping server...")
	if err := app.server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("error on server shutdown: %v", err)
	}

	// Graceful Shutdown for providers
	app.shutdownProviders()
}

func (app *HaaraApplication) RegisterProvider(provider provider.Provider) {
	provider.SetContainer(app.container)
	app.providers = append(app.providers, provider)
}

func (app *HaaraApplication) startProviders() {
	fmt.Println("\nRunning start for every provider...")
	for _, provider := range app.providers {
		provider.OnStart()
	}
	fmt.Println("Start for providers is finished.\n")
}

func (app *HaaraApplication) shutdownProviders() {
	providersCopy := make([]provider.Provider, len(app.providers))
	copy(providersCopy, app.providers)

	slices.Reverse(providersCopy)
	fmt.Println("\nRunning shutdown for every provider...")
	for _, provider := range providersCopy {
		provider.OnShutdown()
	}
	fmt.Println("Shutdown for providers is finished.\n")
}
