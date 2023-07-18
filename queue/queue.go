package queue

import (
	"container/list"
	"errors"
	"sync"
)

/*
	简单实现的线程安全的队列
*/

// 简易实现队列
type Queue struct {
	l  *list.List
	mu sync.RWMutex
}

func (q *Queue) Push(num int64) {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.l.PushBack(num)
}

func (q *Queue) Pop() (int64, error) {
	q.mu.Lock()
	defer q.mu.Unlock()
	if q.l.Len() == 0 {
		return 0, errors.New("empty queue")
	}
	i := q.l.Front()
	q.l.Remove(i)
	return i.Value.(int64), nil
}

func (q *Queue) Length() int {
	return q.l.Len()
}

// 获取 queue 中的所有元素
func (q *Queue) All() []int64 {
	q.mu.RLock()
	defer q.mu.RUnlock()
	res := []int64{}
	if q.l.Len() == 0 {
		return res
	}

	for f := q.l.Front(); f != nil; f = f.Next() {
		res = append(res, f.Value.(int64))
	}
	return res
}

// all: 是否删除队列里所有的这个元素
// 返回删除了多少元素
func (q *Queue) Remove(num int64, all bool) int {
	q.mu.Lock()
	defer q.mu.Unlock()
	count := q.Length()
	if q.l.Len() == 0 {
		return 0
	}
	res := make([]*list.Element, 0, 5)
	for e := q.l.Front(); e != nil; e = e.Next() {
		if e.Value.(int64) == num {
			res = append(res, e)
			if !all {
				break
			}
		}
	}

	for i := range res {
		q.l.Remove(res[i])
	}

	return count - q.Length()
}

func NewQueue() *Queue {
	q := &Queue{}
	q.l = list.New()
	return q
}
