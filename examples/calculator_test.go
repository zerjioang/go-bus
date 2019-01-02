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
message chain: none
difficulty: 1/5
*/

func TestExampleCalculationBus(t *testing.T) {
	//define a global bus instance
	bus := mutex.NewBus()

	//register our subscriber: the calculation engine
	bus.Subscribe("calc", func(message gobus.EventMessage) {
		println("new calculation request arrived")
		a := message.Get("a").(int)
		b := message.Get("b").(int)
		println("received a = ", a)
		println("received b = ", b)
		println("the sum is", a+b)
	})

	//register our data publisher
	bus.EmitWithMessage("calc", gobus.EventPayload{
		"a": 5,
		"b": 10,
	})

	bus.EmitWithMessage("calc", gobus.EventPayload{
		"a": 8,
		"b": 12,
	})

	bus.Shutdown()
}
