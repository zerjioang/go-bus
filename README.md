# Eventbus for Go

Package **go-bus** is a thread-safe, concurrent, pub-sub event bus for embebbed go apps and microservices.

## TL;DR

```go
func TestExampleCalculationBusWithChain(t *testing.T){
	//define a global bus instance
	bus := mutex.NewBus()

	//register our subscriber 1: the calculation engine
	bus.Subscribe("calc", func(message gobus.EventMessage) {
		println("new calculation request arrived")
		a := message.Get("a").(int)
		b := message.Get("b").(int)

		println("received a = ", a)
		println("received b = ", b)
		result := a+b

		bus.Send("calc-print", gobus.EventPayload{
			"result": result,
		})
	})

	//register our subscriber 2: the calculation printer
	bus.Subscribe("calc-print", func(message gobus.EventMessage) {
		result:= message.Get("result").(int)
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
```

```bash
new calculation request arrived
new calculation request arrived
received a =  5
received b =  10
received a =  8
received b =  12
the sum is 15
the sum is 20
```

## Benchmarking

Always do benchmarking with your own data. Here are mine:

### For mutex based version

```bash
BenchmarkEventBus/instantiation-4                      2000000000	         0.41 ns/op	2453.96 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/subscribe-4             	              5000000	          242 ns/op	   4.12 MB/s	      43 B/op	       0 allocs/op
BenchmarkEventBus/subscribe-invalid-no-name-4         	500000000	         3.27 ns/op	 306.27 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/subscribe-invalid-no-listener-4     	300000000	         4.42 ns/op	 226.48 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/publish-4                           	 50000000	         28.1 ns/op	  35.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/pub-sub-4                           	   100000	       139794 ns/op	   0.01 MB/s	      46 B/op	       0 allocs/op
```

Special thanks to http://ernestmicklei.com/2014/11/guava-like-eventbus-for-go/ for it's original approach

## License

All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 * Uses GPL license described below

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.