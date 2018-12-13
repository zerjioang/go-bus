package gobus

// EventPayload represents the context of an event, a simple map.
type EventPayload map[string]interface{}

// EventMessage is send on the bus to all subscribed listeners
type EventMessage struct {
	id      string
	payload EventPayload
}

func NewEventMessage(id string, data map[string]interface{}) EventMessage {
	return EventMessage{id: id, payload: data}
}

func (m *EventMessage) Get(key string) interface{} {
	return m.payload[key]
}

// EventListener is the signature of functions that can handle an EventMessage.
type EventListener func(EventMessage)
