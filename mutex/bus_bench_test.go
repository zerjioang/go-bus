package mutex

import (
	"github.com/zerjioang/go-bus"
	"testing"
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
			_ = SharedBus()
		}
	})

	b.Run("instantiation", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = NewBus()
		}
	})

	b.Run("instantiation-ptr", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			_ = NewBusPtr()
		}
	})

	b.Run("str-to-uint32", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		for n := 0; n < b.N; n++ {
			strTouint32("HelloWorld")
		}
	})

	b.Run("subscribe", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("test", testListener)
		}
	})

	b.Run("subscribe-invalid-no-name", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("", testListener)
		}
	})

	b.Run("subscribe-invalid-no-listener", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("test", nil)
		}
	})

	b.Run("emit-no-subscriber", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		for n := 0; n < b.N; n++ {
			bus.EmitWithMessage("test", exampleMessage)
		}
	})

	b.Run("emit-with-message-no-subscriber", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		for n := 0; n < b.N; n++ {
			bus.EmitWithMessage("test", exampleMessage)
		}
	})

	b.Run("emit-with-1-subscriber", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		bus.Subscribe("test", testListener)
		for n := 0; n < b.N; n++ {
			bus.Emit("test")
		}
	})

	b.Run("emit-with-message-1-subscriber", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		bus.Subscribe("test", testListener)
		for n := 0; n < b.N; n++ {
			bus.EmitWithMessage("test", exampleMessage)
		}
	})

	b.Run("pub-sub", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		bus := NewBus()
		for n := 0; n < b.N; n++ {
			bus.Subscribe("test", testListener)
			bus.EmitWithMessage("test", exampleMessage)
		}
	})
}

func testListener(e gobus.EventMessage) {
	//fmt.Printf("%#v\n", e)
}
