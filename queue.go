package queue

import "sync"

type Queue struct {
	q    chan *task
	once sync.Once
	size int
}

func NewQueue(size int) *Queue {
	return &Queue{size: size}
}

type task struct {
	retChan chan error
	f       func() error
}

func (m *Queue) init() {
	if m.size < 1 {
		m.size = 1
	}
	m.q = make(chan *task, m.size)
	for i := 0; i < m.size; i++ {
		go func() {
			for {
				t := <-m.q
				err := t.f()
				go func() {
					t.retChan <- err
				}()
			}
		}()
	}
}

func (m *Queue) Add(f func() error) chan error {
	m.once.Do(m.init)
	t := &task{}
	t.f = f
	t.retChan = make(chan error, 2)
	m.q <- t
	return t.retChan
}
