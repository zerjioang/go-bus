<p align="center">
  <img alt="gobus" src="https://user-images.githubusercontent.com/6706342/50519588-f3fb3700-0abb-11e9-965f-f582b46b7b46.png" width="500px"></img>
  <h3 align="center"><b>PubSub Event bus for Go</b></h3>
</p>

<p align="center">
    <a href="https://travis-ci.org/zerjioang/go-bus">
      <img alt="Build Status" src="https://travis-ci.org/zerjioang/go-bus.svg?branch=master">
    </a>
    <a href="https://goreportcard.com/report/github.com/zerjioang/go-bus">
       <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/zerjioang/go-bus">
    </a>
    <a href="https://github.com/zerjioang/go-bus/blob/master/LICENSE">
        <img alt="Software License" src="http://img.shields.io/:license-gpl3-brightgreen.svg?style=flat-square">
    </a>
    <a href="https://godoc.org/github.com/zerjioang/go-bus">
       <img alt="Build Status" src="https://godoc.org/github.com/zerjioang/go-bus?status.svg">
    </a>
</p>

Package **go-bus** is a thread-safe, **zero-alloc**, pub-sub event bus for embebbed go apps and microservices.

## Install

```bash
go get github.com/zerjioang/go-bus
```

## TL;DR

```go
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
	bus.Send("calc", gobus.EventPayload{
		"a": 5,
		"b": 10,
	})

	bus.Send("calc", gobus.EventPayload{
		"a": 8,
		"b": 12,
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
BenchmarkEventBus/instantiation-4         				2000000000	         0.47 ns/op	2143.53 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/instantiation-ptr-4     				2000000000	         0.46 ns/op	2190.94 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/str-to-uint32-4         				50000000	        22.2 ns/op	  45.06 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/subscribe-4                        	 5000000	       314 ns/op	   3.18 MB/s	      43 B/op	       0 allocs/op
BenchmarkEventBus/subscribe-invalid-no-name-4        	300000000	         3.90 ns/op	 256.73 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/subscribe-invalid-no-listener-4    	300000000	         5.16 ns/op	 193.86 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/publish-no-subscriber-4            	 2000000	       616 ns/op	   1.62 MB/s	       1 B/op	       0 allocs/op
BenchmarkEventBus/publish-with-subscriber-4          	 2000000	       697 ns/op	   1.43 MB/s	       0 B/op	       0 allocs/op
BenchmarkEventBus/pub-sub-4                          	  200000	    213943 ns/op	   0.00 MB/s	      51 B/op	       0 allocs/op
PASS
```

Special thanks to http://ernestmicklei.com/2014/11/guava-like-eventbus-for-go/ for it's original approach

## License

All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the following conditions are met:

 * Redistributions of source code must retain the above copyright notice, this list of conditions and the following disclaimer.
 * Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following disclaimer in the documentation and/or other materials provided with the distribution.
 * Uses GPL license described below

This program is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.