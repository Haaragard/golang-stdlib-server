package provider

import (
	"github.com/haaragard/golang-stdlib-server/container"
)

type Provider interface {
	SetContainer(*container.HaaraContainer)
	OnStart()
	OnShutdown()
}
