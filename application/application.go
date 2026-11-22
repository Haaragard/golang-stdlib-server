package application

import (
	"slices"

	"github.com/haaragard/golang-stdlib-server/container"
	"github.com/haaragard/golang-stdlib-server/provider"
	"github.com/haaragard/golang-stdlib-server/server"
)

type HaaraApplication struct {
	providers []provider.Provider
	container *container.HaaraContainer
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

func (app *HaaraApplication) Run() error {
	// Start providers
	app.startProviders()

	// Graceful Shutdown for providers
	defer app.shutdownProviders()

	err := app.server.ServeAndListen()
	if err != nil {
		return err
	}

	return nil
}

func (app *HaaraApplication) RegisterProvider(provider provider.Provider) {
	app.providers = append(app.providers, provider)
}

func (app *HaaraApplication) startProviders() {
	for _, provider := range app.providers {
		provider.OnStart()
	}
}

func (app *HaaraApplication) shutdownProviders() {
	providersCopy := make([]provider.Provider, len(app.providers))
	copy(providersCopy, app.providers)

	slices.Reverse(providersCopy)
	for _, provider := range providersCopy {
		provider.OnShutdown()
	}
}
