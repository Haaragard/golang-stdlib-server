package container

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type HaaraContainer struct {
	mutex    sync.RWMutex
	services map[reflect.Type]any
}

func NewHaaraContainer() *HaaraContainer {
	return &HaaraContainer{services: make(map[reflect.Type]any)}
}

// Register accepts any dynamic pointer/value at runtime.
func (hc *HaaraContainer) RegisterDynamic(instance any) error {
	if instance == nil {
		return errors.New("cannot register a nil instance")
	}

	t := reflect.TypeOf(instance)
	if t.Kind() != reflect.Pointer {
		return fmt.Errorf("expected a pointer, got %s", t.Kind())
	}

	hc.mutex.Lock()
	defer hc.mutex.Unlock()

	hc.services[t] = instance
	return nil
}

// Register accepts any dynamic pointer/value at runtime.
func (hc *HaaraContainer) RegisterAs(reference reflect.Type, instance any) error {
	if instance == nil {
		return errors.New("cannot register a nil instance")
	}

	t := reflect.TypeOf(instance)
	if t.Kind() != reflect.Pointer {
		return fmt.Errorf("expected a pointer, got %s", t.Kind())
	}

	hc.mutex.Lock()
	defer hc.mutex.Unlock()

	hc.services[reference] = instance
	return nil
}

func (hc *HaaraContainer) Resolve(reference reflect.Type) any {
	hc.mutex.RLock()
	defer hc.mutex.RUnlock()

	instance, ok := hc.services[reference]
	if !ok {
		panic(fmt.Errorf("service not registered for type: %v", reference))
	}

	return instance
}
