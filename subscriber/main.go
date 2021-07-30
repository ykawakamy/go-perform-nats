package main

import (
	"flag"
	"strconv"
	"sync"
	"time"

	config "go-perform-nats/config"
	util "go-perform-nats/util"

	"github.com/nats-io/nats.go"
)

func main() {
	// 引数
	arg_loop_ns := flag.Int("d", 5000*1000*1000, "duration ns")
	arg_iter := flag.Int("i", 20, "iteration count")
	arg_thread := flag.Int("c", 10, "thread(gorouting) count")
	arg_topic_expr := flag.String("t", "test", "topic expression (e.g. /test/ or /test/:1,2,3 or /test/:1-3) ")
	arg_server_url := flag.String("s", config.Config_DefaultUrl, "connection string")
	arg_topic_divide := flag.Int("v", 1, "divide value")
	flag.Parse()

	//
	var iter = *arg_iter
	var thread = *arg_thread
	var loop_ns int64 = int64(*arg_loop_ns)

	// パフォーマンス測定用カウンタ生成
	pcMap := util.CreatePerformCounterMap()
	hist := util.CreateHistogram()

	// TopicSupplierFactory生成
	factory := util.CreateFactory().
		ParseTopicExpression(*arg_topic_expr).
		SetDistoribution(*arg_topic_divide)

	// Subscriberのセットアップ
	println("connecting")
	wg := sync.WaitGroup{}
	wg.Add(thread)
	for i := 0; i < thread; i++ {
		thread_id := strconv.Itoa(i)
		go func() {
			supplier := factory.Build()

			nc, err := nats.Connect(*arg_server_url)
			if err != nil {
				panic("Connect error:[" + err.Error() + "] thread:" + thread_id)
			}

			for _, topic := range supplier.GetAll() {
				_, err := nc.Subscribe(topic, func(m *nats.Msg) {
					ping := util.DeserialPing(m.Data)
					key := thread_id + "<>" + m.Subject
					pcMap.Perform(key, &ping)
					hist.IncreamentPing(&ping)
				})

				if err != nil {
					panic("Subscribe error:[" + err.Error() + "] thread:" + thread_id + " topic:" + topic)
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()
	println("connected")

	// パフォーマンス測定
	pcMap.CollectAndReset()
	var st int64 = time.Now().UnixNano()
	var et int64 = 0
	println("start-benchmark")
	for i := 0; i < iter; i++ {
		// 別スレッドで処理するため、メインスレッドはスリープする。
		time.Sleep(time.Duration(loop_ns) * time.Nanosecond)
		et = time.Now().UnixNano()

		snap := pcMap.CollectAndReset()
		snap.Print(et - st)
		st = et
	}
	hist.Print()
	println("end-benchmark")
}
