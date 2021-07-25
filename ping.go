package main

import (
	"time"
)

type Ping struct {
	seq          uint32
	send_tick    uint64
	receive_tick uint64
	latency      uint64
}

func CreatePing(seq uint32) Ping {
	return Ping{
		seq:       seq,
		send_tick: uint64(time.Now().UnixNano()),
	}
}

func (s Ping) Serialize() []byte {
	return []byte{
		byte(s.seq >> 24),
		byte(s.seq >> 16),
		byte(s.seq >> 8),
		byte(s.seq),
		//---
		byte(s.send_tick >> 56),
		byte(s.send_tick >> 48),
		byte(s.send_tick >> 40),
		byte(s.send_tick >> 32),
		byte(s.send_tick >> 24),
		byte(s.send_tick >> 16),
		byte(s.send_tick >> 8),
		byte(s.send_tick),
	}
}
