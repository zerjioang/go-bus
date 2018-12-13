package examples

import (
	"github.com/zerjioang/go-bus"
	"github.com/zerjioang/go-bus/mutex"

	"testing"
)

/*
this is a stupid example showing how a complex calculations can be triggered using the event bus
in this example, we are just calculating the sum of 2 integers

Concurrency: none
message chain: yes. two items
difficulty: 2/5
*/

func TestExampleCalculationBusWithChain(t *testing.T) {
	//define a global bus instance
	bus := mutex.NewBus()

	//register our subscriber 1: the calculation engine
	bus.Subscribe("calc", func(message gobus.EventMessage) {
		println("new calculation request arrived")
		a := message.Get("a").(int)
		b := message.Get("b").(int)

		println("received a = ", a)
		println("received b = ", b)
		result := a + b

		bus.Send("calc-print", gobus.EventPayload{
			"result": result,
		})
	})

	//register our subscriber 2: the calculation printer
	bus.Subscribe("calc-print", func(message gobus.EventMessage) {
		result := message.Get("result").(int)
		println("the sum is", result)
	})

	bus.Send("calc", gobus.EventPayload{
		"a": 8,
		"b": 12,
	})

	//register our data publisher
	bus.Send("calc", gobus.EventPayload{
		"a": 5,
		"b": 10,
	})

	bus.Shutdown()
}
