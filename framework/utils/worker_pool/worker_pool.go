package worker_pool

import (
	"context"
	"sync"
)

type WorkerPool[T any] struct {
	workerFunc  func(T)
	numWorkers  int
	taskQueue   chan T
	doneChannel chan struct{}
	isWait      bool
	wg          sync.WaitGroup
	ctx         context.Context
	cancelFunc  context.CancelFunc
}

// NewWorkerPool is a method to deal with goroutine
// limit the maximum number of goroutine at a time, make sure that we always process data at the max capacity of the
// optimized of the goroutine
//
// How to use:
//
//		pool := worker_pool.NewWorkerPool[int64](taskFunc, 50)
//		pool.Start()
//	 defer pool.Stop()
func NewWorkerPool[T any](workerFunc func(T), numWorkers int) *WorkerPool[T] {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool[T]{
		workerFunc:  workerFunc,
		numWorkers:  numWorkers,
		taskQueue:   make(chan T),
		doneChannel: make(chan struct{}, numWorkers),
		wg:          sync.WaitGroup{},
		ctx:         ctx,
		cancelFunc:  cancel,
	}
}

func (wp *WorkerPool[T]) Start() {
	for i := 0; i < wp.numWorkers; i++ {
		wp.wg.Add(1)
		go wp.worker()
	}
}

func (wp *WorkerPool[T]) Stop() {
	close(wp.taskQueue)
	wp.wg.Wait()
	wp.cancelFunc()
}

func (wp *WorkerPool[T]) SubmitTask(task T) {
	select {
	case <-wp.ctx.Done():
		return
	case wp.taskQueue <- task:
	}
}

func (wp *WorkerPool[T]) worker() {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			return
		case task, ok := <-wp.taskQueue:
			if !ok {
				return
			}
			wp.workerFunc(task)

			if wp.isWait {
				wp.doneChannel <- struct{}{}
			}
		}
	}
}

func (wp *WorkerPool[T]) SetWaitTask(isWait bool) {
	wp.isWait = isWait
}

func (wp *WorkerPool[T]) WaitForNTask(n int, done chan struct{}) {
	for i := 0; i < n; i++ {
		<-wp.doneChannel
	}

	done <- struct{}{}
}
