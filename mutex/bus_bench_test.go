package mutex_test

import (
	"testing"
	"time"

	"github.com/zerjioang/go-bus"
	"github.com/zerjioang/go-bus/mutex"
)

/*
* Benchmark functions start with Benchmark not Test.

* Benchmark functions are run several times by the testing package.
  The value of b.N will increase each time until the benchmark runner
  is satisfied with the stability of the benchmark. This has some important
  ramifications which we’ll investigate later in this article.

* Each benchmark is run for a minimum of 1 second by default.
  If the second has not elapsed when the Benchmark function returns,
  the value of b.N is increased in the sequence 1, 2, 5, 10, 20, 50, …
  and the function run again.

* the for loop is crucial to the operation of the benchmark driver
  it must be: for n := 0; n < b.N; n++

* Add b.ReportAllocs() at the beginning of each benchmark to know allocations
* Add b.SetBytes(1) to measure byte transfers

  Optimization info: https://stackimpact.com/blog/practical-golang-benchmarks/
*/

var (
	exampleMessage = gobus.EventPayload{"pi": 3.14159}
)

func BenchmarkEventBus(b *testing.B) {

	b.Run("shared-bus", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = mutex.SharedBus()
		}
	})

	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = mutex.NewBus()
		}
	})

	b.Run("instantiation-ptr", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = mutex.NewBusPtr()
		}
	})

	b.Run("str-to-uint32", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			mutex.StrTouint32("HelloWorld")
		}
	})

	b.Run("str-to-uint32-concurrent", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			go func() {
				if mutex.StrTouint32("HelloWorld") != 926844193 {
					b.Error("hash function failed")
				}
			}()
		}
		time.Sleep(10 * time.Second)
	})

	b.Run("subscribe", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := mutex.NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("test", testListener)
		}
	})

	b.Run("subscribe-invalid-no-name", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := mutex.NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("", testListener)
		}
	})

	b.Run("subscribe-invalid-no-listener", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := mutex.NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("test", nil)
		}
	})

	b.Run("publish-no-subscriber", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := mutex.NewBus()
		for n := 0; n < b.N; n++ {
			bus.Send("test", exampleMessage)
		}
	})

	b.Run("publish-with-subscriber", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := mutex.NewBus()
		bus.Subscribe("test", testListener)
		for n := 0; n < b.N; n++ {
			bus.Send("test", exampleMessage)
		}
	})

	b.Run("pub-sub", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := mutex.NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("test", testListener)
			bus.Send("test", exampleMessage)
		}
	})
}

func testListener(e gobus.EventMessage) {
	//fmt.Printf("%#v\n", e)
}
