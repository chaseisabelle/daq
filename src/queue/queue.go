package queue

import (
	"container/list"
	"context"
	"fmt"
	"github.com/chaseisabelle/daq/src/message"
	"sync"
)

type Queue struct {
	mutex sync.RWMutex
	list  *list.List
}

func New() (*Queue, error) {
	return &Queue{
		list: list.New(),
	}, nil
}

func (q *Queue) Enqueue(_ context.Context, m *message.Message) error {
	q.mutex.Lock()

	defer q.mutex.Unlock()

	q.list.PushBack(m)

	return nil
}

func (q *Queue) Dequeue(_ context.Context) (*message.Message, error) {
	q.mutex.RLock()

	defer q.mutex.RUnlock()

	e := q.list.Front()

	if e == nil {
		return nil, nil
	}

	m, ok := q.list.Remove(e).(*message.Message)

	if !ok {
		return m, fmt.Errorf("invalid value: %+v", e.Value)
	}

	return m, nil
}

func (q *Queue) Requeue(_ context.Context, m *message.Message) error {
	q.mutex.Lock()

	defer q.mutex.Unlock()

	q.list.PushFront(m)

	return nil
}
