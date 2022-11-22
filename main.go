package main

import (
	"context"
	"flag"
	"sync"
	"worker-pool/w_pool"
)

func main() {
	wnum := flag.Int("wnum", 5, "number of workers")
	tnum := flag.Int("tnum", 5, "number of tasks")
	flag.Parse()
	w := *wnum
	t := *tnum

	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	receiver := make(chan *w_pool.Task)

	pool := w_pool.NewPool(w, wg, receiver)
	pool.Run(ctx)

	toReceiver(cancel, receiver, t)

	pool.Wg.Wait()
}

func toReceiver(cancel context.CancelFunc, receiver chan *w_pool.Task, t int) {
	go func() {
		for i := 0; i < t; i++ {
			receiver <- w_pool.NewTask(i, func(k interface{}) error {
				return nil
			})
		}

		close(receiver)

		cancel()
	}()
}
