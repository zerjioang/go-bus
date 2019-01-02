// +build example

package benchmark

import (
	"fmt"
	"github.com/zerjioang/go-bus/mutex"
	"net/http"
	"net/http/pprof"
	"os"
	pprof2 "runtime/pprof"
	"testing"
	"time"

	"github.com/zerjioang/go-bus"
)

func TestEventBusProfilingWeb(t *testing.T) {
	go func() {
		r := http.NewServeMux()

		// Register pprof handlers
		r.HandleFunc("/debug/pprof/", pprof.Index)
		r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		r.HandleFunc("/debug/pprof/profile", pprof.Profile)
		r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		r.HandleFunc("/debug/pprof/trace", pprof.Trace)

		http.ListenAndServe(":8080", r)
	}()
	//infinite loop
	for {
		bus := mutex.NewBus()
		bus.Subscribe("test", func(e gobus.EventMessage) {
			//fmt.Printf("%#v\n", e)
		})
		bus.EmitWithMessage("test", gobus.EventPayload{"pi": 3.14159})
		time.Sleep(10 * time.Millisecond)
	}
}

func TestEventBusProfilingFile(t *testing.T) {
	f, err := os.Create("gobus.pprof")
	if err == nil {
		pprof2.StartCPUProfile(f)
		defer pprof2.StopCPUProfile()
	}
	bus := mutex.NewBus()
	bus.Subscribe("test", func(e gobus.EventMessage) {
		//fmt.Printf("%#v\n", e)
	})
	for i := 0; i < 5000; i++ {
		bus.EmitWithMessage("test", gobus.EventPayload{"pi": 3.14159})
		fmt.Println(i)
		//time.Sleep(2 * time.Millisecond)
	}
	bus.Shutdown()
}
