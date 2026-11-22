package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"slices"
	"sync"
	"syscall"

	"github.com/haaragard/golang-stdlib-server/container"
	"github.com/haaragard/golang-stdlib-server/provider"
	"github.com/haaragard/golang-stdlib-server/server"
)

type HaaraApplication struct {
	wg        sync.WaitGroup
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
	if err := app.Stop(context); err != nil {
		log.Fatalf("could not shutdown: %v\n", err)
	}

	// Ends program
	os.Exit(0)
}

func (app *HaaraApplication) Stop(ctx context.Context) error {
	// Graceful Shutdown for providers
	app.shutdownProviders()

	fmt.Println("Waiting for all HaaraApplication GoRoutines to finish.")
	app.wg.Wait()

	return app.server.Shutdown(ctx)
}

func (app *HaaraApplication) RegisterProvider(provider provider.Provider) {
	provider.SetContainer(app.container)
	provider.SetWaitGroup(&app.wg)
	app.providers = append(app.providers, provider)
}

func (app *HaaraApplication) startProviders() {
	for _, provider := range app.providers {
		app.wg.Add(1)
		go provider.OnStart()
	}
}

func (app *HaaraApplication) shutdownProviders() {
	providersCopy := make([]provider.Provider, len(app.providers))
	copy(providersCopy, app.providers)

	slices.Reverse(providersCopy)
	for _, provider := range providersCopy {
		app.wg.Add(1)
		go provider.OnShutdown()
	}
}
