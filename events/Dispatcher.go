package events

import (
	"fmt"
)

type Dispatcher struct {
	Events map[string]*IEventHandler
}

func (dispatcher *Dispatcher) Register(eventHandler *IEventHandler, name string) error {
	if _, ok := dispatcher.Events[name]; ok {
		return fmt.Errorf("The %s event is already registered", name)
	}

	dispatcher.Events[name] = eventHandler

	return nil
}

func (dispatcher *Dispatcher) Dispatch(name string, payload *[]byte) error {
	var eventHandler IEventHandler

	if _, ok := dispatcher.Events[name]; !ok {
		return fmt.Errorf("The %s event is not registered", name)
	}

	if _, ok := dispatcher.Events[name]; ok {
		eventHandler = *dispatcher.Events[name]
		_, err := eventHandler.Handle(*payload)
		return err
	} else {
		return fmt.Errorf("The handler %s is not registered", name)
	}
}
