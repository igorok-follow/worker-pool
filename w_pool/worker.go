package w_pool

import (
	"context"
	"fmt"
	"strconv"
	"sync"
)

type Worker struct {
	Id       int
	tasks    chan *Task
	progress int
}

func NewWorker(id int, broadcast chan *Task) *Worker {
	return &Worker{
		Id:    id,
		tasks: broadcast,
	}
}

func (w *Worker) Start(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				fmt.Println("progress of " + strconv.Itoa(w.Id) + " worker is: " + strconv.Itoa(w.progress) + " tasks")
				return
			case t := <-w.tasks:
				w.progress++
				t.Run(w.Id)
			}
		}
	}()
}
