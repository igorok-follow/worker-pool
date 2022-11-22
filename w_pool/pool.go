package w_pool

import (
	"context"
	"sync"
)

type Pool struct {
	BroadCast chan *Task
	WorkerNum int
	Wg        *sync.WaitGroup
	receiver  chan *Task
}

func NewPool(workerNum int, wg *sync.WaitGroup, receiver chan *Task) *Pool {
	return &Pool{
		BroadCast: make(chan *Task, 100),
		WorkerNum: workerNum,
		Wg:        wg,
		receiver:  receiver,
	}
}

func (p *Pool) Run(ctx context.Context) {
	for i := 0; i < p.WorkerNum; i++ {
		w := NewWorker(i, p.BroadCast)
		w.Start(ctx, p.Wg)
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				close(p.BroadCast)
				return
			case t := <-p.receiver:
				p.BroadCast <- t
			}
		}
	}()
}
