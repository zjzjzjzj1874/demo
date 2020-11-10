package main

import (
	"container/list"
	"fmt"
	"sync"
)

var lock sync.Mutex

type Queue struct {
	List *list.List
}

func NewQueue() *Queue {
	return &Queue{
		List: list.New(),
	}
}

func (q *Queue) Push(v interface{}) {
	defer lock.Unlock()
	lock.Lock()
	q.List.PushFront(v)
}

func (q *Queue) Pop() interface{} {
	defer lock.Unlock()
	lock.Lock()
	var v interface{}
	iter := q.List.Back()
	if iter != nil {
		v = iter.Value
		q.List.Remove(iter)
	}
	return v
}

func (q *Queue) PushBack(v interface{}) {
	defer lock.Unlock()
	lock.Lock()
	q.List.PushBack(v)
}

func (q *Queue) Dump() {
	for iter := q.List.Back(); iter != nil; iter = iter.Prev() {
		fmt.Println("item:", iter.Value)
	}
}

func (q *Queue) Len() int {
	return q.List.Len()
}
