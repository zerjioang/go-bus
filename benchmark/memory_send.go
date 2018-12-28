// +build example

package main

import (
	"github.com/pkg/profile"
	"github.com/zerjioang/go-bus"
	"github.com/zerjioang/go-bus/mutex"
)

func main() {
	defer profile.Start(profile.MemProfile).Stop()
	bus := mutex.NewBus()
	bus.Subscribe("test", func(e gobus.EventMessage) {
		//fmt.Printf("%#v\n", e)
	})
	e := gobus.EventPayload{"pi": 3.14159}
	for i := 0; i < 5000000; i++ {
		bus.Send("test", e)
	}
	bus.Shutdown()
}
