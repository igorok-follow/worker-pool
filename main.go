package main

import (
	"context"
	"sync"
	"worker-pool/w_pool"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	receiver := make(chan *w_pool.Task)

	pool := w_pool.NewPool(100, wg, receiver)
	pool.Run(ctx)

	toReceiver(receiver, cancel)

	pool.Wg.Wait()
}

func toReceiver(receiver chan *w_pool.Task, cancel context.CancelFunc) {
	go func() {
		for i := 0; i < 1000; i++ {
			receiver <- w_pool.NewTask(i, func(k interface{}) error {
				return nil
			})
		}

		close(receiver)

		cancel()
	}()
}
