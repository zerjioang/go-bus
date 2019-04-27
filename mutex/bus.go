package mutex

import (
	"hash/fnv"
	"sync"
	"unsafe"

	"github.com/zerjioang/go-bus"
)

// The Bus allows publish-subscribe-style communication between components/modules
type Bus struct {
	listeners     map[uint32][]gobus.EventListener
	listenerMutex sync.RWMutex
	wg            sync.WaitGroup
}

var (
	onceCheck        sync.Once
	sharedBus *Bus
	h = fnv.New32a()
	hlock sync.Mutex
)

func SharedBus() *Bus {
	onceCheck.Do(func() {
		sharedBus = NewBusPtr()
	})
	return sharedBus
}

func NewBus() Bus {
	bus := Bus{}
	//bus.listenerMutex = new(sync.RWMutex)
	//bus.wg = new(sync.WaitGroup)
	return bus
}

func NewBusPtr() *Bus {
	bus := new(Bus)
	//bus.listenerMutex = new(sync.RWMutex)
	//bus.wg = new(sync.WaitGroup)
	return bus
}

func strTouint32(s string) uint32 {
	//h := fnv.New32a()
	hlock.Lock()
	h.Reset()
	h.Write(strToByte(s))
	v := h.Sum32()
	hlock.Unlock()
	return v
}

func strToByte(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

// Subscribe adds an EventListener to be called when an event is posted.
func (e *Bus) Subscribe(topic string, listener gobus.EventListener) {
	if topic == "" || listener == nil {
		return
	}
	var list []gobus.EventListener

	e.listenerMutex.Lock()
	if e.listeners == nil {
		e.listeners = make(map[uint32][]gobus.EventListener)
	}
	id := strTouint32(topic)
	list, _ = e.listeners[id]
	list = append(list, listener)
	e.listeners[id] = list
	e.listenerMutex.Unlock()
}

// EmitWithMessage sends an event to all subscribed listeners.
// Parameter data is optional ; EmitWithMessage can only have one map parameter.
func (e *Bus) EmitWithMessage(topic string, data map[string]interface{}) {
	if topic == "" {
		return
	}
	if data == nil {
		return
	}

	e.wg.Add(1)
	go func() {
		e.listenerMutex.RLock()
		id := strTouint32(topic)
		list, present := e.listeners[id]
		e.listenerMutex.RUnlock()
		if present {
			e.sendEvent(list, topic, data)
			e.wg.Done()
		}
	}()
}

// Send sends an event to all subscribed listeners.
// Parameter data is optional ; Send can only have one map parameter.
func (e *Bus) Emit(topic string) {
	e.wg.Add(1)
	go func() {
		e.listenerMutex.RLock()
		id := strTouint32(topic)
		list, present := e.listeners[id]
		e.listenerMutex.RUnlock()
		if present {
			e.sendEvent(list, topic, nil)
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
	if e != nil {
		e.wg.Wait()
	}
}
