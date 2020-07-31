package queue

import (
	"fmt"
	"testing"
)

func TestQueue_New(t *testing.T) {
	q := NewQueue()
	t.Log(q.Length())
}

func TestQueue_Push(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	if q.Length() > 0 {
		t.Log(q.l)
	} else {
		t.Error("error")
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(1)
	n, _ := q.Pop()
	if n == 1 {
		t.Log("ok")
	} else {
		t.Error("error")
	}
	if n, _ = q.Pop(); n == 2 {
		t.Log("ok")
	} else {
		t.Error("error")
	}
}

func TestQueue_Pop_Empty(t *testing.T) {
	q := NewQueue()
	_, err := q.Pop()
	if err != nil {
		t.Log(err)
	} else {
		t.Error("get value from empty queue wrong")
	}
}

func TestQueue_Remove(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(2)
	PrintAll(q)
	n := q.Remove(2, false)
	fmt.Println(n)
	PrintAll(q)
	q.Push(2)
	PrintAll(q)
	n = q.Remove(2, true)
	fmt.Println(n)
	PrintAll(q)

	q = NewQueue()
	q.Push(2)
	q.Push(2)
	q.Push(2)
	q.Push(2)
	PrintAll(q)
	n = q.Remove(2, true)
	fmt.Println(n)
	PrintAll(q)
}

func TestQueueAll(t *testing.T) {
	q := NewQueue()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	t.Log(q.All())
}

func PrintAll(q *Queue) {
	if q.Length() == 0 {
		fmt.Println("")
		return
	}
	fmt.Println(q.All())
}
