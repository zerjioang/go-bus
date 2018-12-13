package mutex

import (
	"sync"

	"github.com/zerjioang/go-bus"
)

// The Bus allows publish-subscribe-style communication between components/modules
type Bus struct {
	listeners     map[string][]gobus.EventListener
	listenerMutex sync.RWMutex
	wg            sync.WaitGroup
}

func NewBus() Bus {
	bus := Bus{}
	return bus
}

// Subscribe adds an EventListener to be called when an event is posted.
func (e *Bus) Subscribe(id string, listener gobus.EventListener) {
	if id == "" {
		return
	}
	if listener == nil {
		return
	}

	var present bool
	var list []gobus.EventListener

	//read current map status
	/*
		in the same lock period
		* check if map is empty
		* if not empty, get requested id
	*/
	e.listenerMutex.RLock()
	empty := e.listeners == nil
	if !empty {
		list, present = e.listeners[id]
	}
	e.listenerMutex.RUnlock()

	//add requested listener to its list
	//no lock required for now
	if !present {
		list = []gobus.EventListener{}
	}
	list = append(list, listener)

	//create new map, only if its empty
	/*
		in the same lock period,
			* we create the listener holder
			* add the list we already have
	*/
	e.listenerMutex.Lock()
	if empty {
		e.listeners = make(map[string][]gobus.EventListener)
	}
	e.listeners[id] = list
	e.listenerMutex.Unlock()
}

// Send sends an event to all subscribed listeners.
// Parameter data is optional ; Send can only have one map parameter.
func (e *Bus) Send(topic string, data map[string]interface{}) {
	if topic == "" {
		return
	}
	if data == nil {
		return
	}

	e.wg.Add(1)
	go func() {
		e.listenerMutex.RLock()
		list, present := e.listeners[topic]
		e.listenerMutex.RUnlock()
		if present {
			e.sendEvent(list, topic, data)
			e.wg.Done()
		}
	}()
}

func (e *Bus) sendEvent(receivers []gobus.EventListener, id string, data map[string]interface{}) {
	event := gobus.NewEventMessage(id, data)
	for _, each := range receivers[:] { // iterate over unmodifyable copy
		each(event)
	}
}

/*
wait to all messages to be processed
*/
func (e *Bus) Shutdown() {
	e.wg.Wait()
}
