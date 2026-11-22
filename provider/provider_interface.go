package provider

import (
	"sync"

	"github.com/haaragard/golang-stdlib-server/container"
)

type Provider interface {
	SetWaitGroup(*sync.WaitGroup)
	SetContainer(*container.HaaraContainer)
	OnStart()
	OnShutdown()
}
