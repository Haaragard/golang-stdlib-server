package provider

type Provider interface {
	OnStart()
	OnShutdown()
}
