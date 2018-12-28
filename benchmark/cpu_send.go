package main

import (
	"os"
	"runtime/pprof"

	"github.com/zerjioang/go-bus"
	"github.com/zerjioang/go-bus/mutex"
)

func main() {
	f, err := os.Create("gobus_cpu.pprof")
	if err == nil {
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
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
