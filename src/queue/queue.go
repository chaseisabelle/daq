package queue

import (
	"container/list"
	"fmt"
	"github.com/chaseisabelle/daq/src/message"
)

type Queue struct {
	list *list.List
}

func New() (*Queue, error) {
	return &Queue{
		list: list.New(),
	}, nil
}

func (q *Queue) Enqueue(m *message.Message) error {
	q.list.PushBack(m)

	return nil
}

func (q *Queue) Dequeue() (*message.Message, error) {
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

func (q *Queue) Requeue(m *message.Message) error {
	q.list.PushFront(m)

	return nil
}
