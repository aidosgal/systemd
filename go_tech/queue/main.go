package main

import (
	"container/heap"
	"context"
	"sync"
	"time"
)

type Task struct {
	ID       string
	Priority int
	Timeout  time.Duration
	Execute  func(ctx context.Context) error
}
type internalTask struct {
	task  Task
	index int
}

type TaskQueueConfig struct {
	WorkerLimit int
}

type taskQueue struct {
	cfg       TaskQueueConfig
	tasks     heap.Interface
	taskCh    chan *internalTask
	wg        sync.WaitGroup
	mu        sync.Mutex
	cond      *sync.Cond
	running   bool
	shutdown  bool
	cancelCtx context.CancelFunc
}

type TaskQueue interface {
}

func New(cfg TaskQueueConfig) TaskQueue {
	_, _ = context.WithCancel(context.Background())
	q := &taskQueue{
		cfg:    cfg,
		taskCh: make(chan *internalTask),
	}

	return q
}
