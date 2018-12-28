package mutex

import (
	"fmt"

	"github.com/zerjioang/go-bus"

	"sync"
	"testing"
	"time"
)

func TestEventBus(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		_ = NewBus()
	})

	t.Run("instantiation-ptr", func(t *testing.T) {
		busPtr := NewBusPtr()
		if busPtr == nil {
			t.Error("failed to instantiate a bus via pointer")
		}
	})

	t.Run("str-to-uint32", func(t *testing.T) {
		if StrTouint32("HelloWorld") != 926844193 {
			t.Error("hash function failed")
		}
	})

	t.Run("pubsub", func(t *testing.T) {
		bus := NewBus()
		bus.Subscribe("test", func(e gobus.EventMessage) {
			fmt.Printf("A: %#v\n", e)
		})
		bus.Subscribe("test", func(e gobus.EventMessage) {
			fmt.Printf("B: %#v\n", e)
		})
		bus.Send("test", gobus.EventPayload{"pi": 3.14159})
	})

	t.Run("pubsub-concurrent", func(t *testing.T) {
		bus := NewBus()

		var wg sync.WaitGroup
		wg.Add(10)

		for i := 0; i < 2; i++ {
			go bus.Subscribe("test", func(e gobus.EventMessage) {
				fmt.Printf("%#v\n", e)
			})
		}

		fmt.Println("waiting to all suscribers to be ready")
		time.Sleep(time.Second * 2)

		for i := 0; i < 10; i++ {
			go func(thread *sync.WaitGroup, idx int) {
				fmt.Println("sending message to bus", idx)
				go bus.Send("test", gobus.EventPayload{"index": idx})
				fmt.Println("message sent to bus", idx)
				defer thread.Done()
			}(&wg, i)
		}
		wg.Wait()
	})
}
