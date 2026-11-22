package application

import (
	"context"
	"fmt"
	"slices"
	"sync"

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

func (app *HaaraApplication) Run() {
	// Start providers
	app.startProviders()
	go app.server.ServeAndListen()
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
