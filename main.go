package main

import (
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {

	var loop_ns int64 = 5000 * 1000 * 1000

	supplier := CreateTopicNameSupplier()

	// Connect to a server
	println("connecting")
	nc, _ := nats.Connect(config_NatsUrl)
	println("connected")

	var seq uint32 = 0
	for i := 0; i < 20; i++ {
		var cnt int64 = 0
		var st int64 = time.Now().UnixNano()
		var et int64 = 0

		for et-st < loop_ns {
			et = time.Now().UnixNano()

			_ = nc.Publish(supplier.Get(), CreatePing(seq).Serialize())
			seq++
			cnt++
		}
		fmt.Printf("%s: %d ns. %d times. %f ns/op\n", "", et-st, cnt, float64(et-st)/float64(cnt))
	}
}
