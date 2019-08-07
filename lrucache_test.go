package LruCache

import (
	"container/list"
	"testing"
)

type Elem struct {
	key   int
	value string
}

func Test_New(t *testing.T) {
	lc := New(5)
	if lc.Len() != 0 {
		t.Error("case 1 failed")
	}
}

func Test_Put(t *testing.T) {
	lc := New(0)
	lc.Put(1, "1")
	if lc.Len() != 0 {
		t.Error("case 1.1 failed")
	}

	lc = New(5)
	lc.Put(1, "1")
	lc.Put(2, "2")
	lc.Put(1, "3")
	if lc.Len() != 2 {
		t.Error("case 2.1 failed")
	}

	l := list.New()
	l.PushBack(&Elem{
		key:   1,
		value: "3",
	})
	l.PushBack(&Elem{
		key:   2,
		value: "2",
	})

	e := l.Front()
	for c := lc.Head(); c != nil; c = c.Next() {
		v := e.Value.(*Elem)
		if c.key.(int) != v.key {
			t.Error("case 2.2 failed:", c.key.(int), v.key)
		}
		if c.value.(string) != v.value {
			t.Error("case 2.3 failed: ", c.value.(string), v.value)
		}
		e = e.Next()
	}

	lc.Put(3, "4")
	lc.Put(4, "5")
	lc.Put(5, "6")
	lc.Put(2, "7")
	if lc.Len() != 5 {
		t.Error("case 3.1 failed")
	}

	l = list.New()
	l.PushBack(&Elem{2, "7"})
	l.PushBack(&Elem{5, "6"})
	l.PushBack(&Elem{4, "5"})
	l.PushBack(&Elem{3, "4"})
	l.PushBack(&Elem{1, "3"})

	r1 := list.New()
	r1.PushBack(&Elem{1, "3"})
	r1.PushBack(&Elem{3, "4"})
	r1.PushBack(&Elem{4, "5"})
	r1.PushBack(&Elem{5, "6"})
	r1.PushBack(&Elem{2, "7"})

	e = l.Front()
	for c := lc.Head(); c != nil; c = c.Prev() {
		v := e.Value.(*Elem)
		if c.key.(int) != v.key {
			t.Error("case 3.4 failed: ", c.key.(int), v.key)
		}
		if c.value.(string) != v.value {
			t.Error("case 3.5 failed: ", c.value.(string), v.value)
		}
		e = e.Next()
	}

	e = r1.Front()
	for c := lc.Tail(); c != nil; c = c.Prev() {
		v := e.Value.(*Elem)
		if c.key.(int) != v.key {
			t.Error("case 3.4 failed: ", c.key.(int), v.key)
		}
		if c.value.(string) != v.value {
			t.Error("case 3.5 failed: ", c.value.(string), v.value)
		}
		e = e.Next()
	}

	lc.Put(6, "8")
	if lc.Len() != 5 {
		t.Error("case 4.1 failed")
	}

	l = list.New()
	l.PushBack(&Elem{6, "8"})
	l.PushBack(&Elem{2, "7"})
	l.PushBack(&Elem{5, "6"})
	l.PushBack(&Elem{4, "5"})
	l.PushBack(&Elem{3, "4"})

	e = l.Front()
	for c := lc.Head(); c != nil; c = c.Next() {
		v := e.Value.(*Elem)
		if c.key.(int) != v.key {
			t.Error("case 4.2 failed: ", c.key.(int), v.key)
		}
		if c.value.(string) != v.value {
			t.Error("case 4.3 failed: ", c.value.(string), v.value)
		}
		e = e.Next()
	}
}

func Test_Delete(t *testing.T) {
	lc := New(5)
	lc.Put(3, "4")
	lc.Put(4, "5")
	lc.Put(5, "6")
	lc.Put(2, "7")
	lc.Put(6, "8")
	lc.Delete(5)

	l := list.New()
	l.PushBack(&Elem{6, "8"})
	l.PushBack(&Elem{2, "7"})
	l.PushBack(&Elem{4, "5"})
	l.PushBack(&Elem{3, "4"})
	if lc.Len() != 4 {
		t.Error("case 1.1 failed")
	}

	e := l.Front()
	for c := lc.Head(); c != nil; c = c.Next() {
		v := e.Value.(*Elem)
		if c.key.(int) != v.key {
			t.Error("case 1.2 failed: ", c.key.(int), v.key)
		}
		if c.value.(string) != v.value {
			t.Error("case 1.3 failed: ", c.value.(string), v.value)
		}
		e = e.Next()
	}
	lc.Delete(6)

	l = list.New()
	l.PushBack(&Elem{2, "7"})
	l.PushBack(&Elem{4, "5"})
	l.PushBack(&Elem{3, "4"})
	if lc.Len() != 3 {
		t.Error("case 2.1 failed")
	}

	e = l.Front()
	for c := lc.Head(); c != nil; c = c.Next() {
		v := e.Value.(*Elem)
		if c.key.(int) != v.key {
			t.Error("case 2.2 failed: ", c.key.(int), v.key)
		}
		if c.value.(string) != v.value {
			t.Error("case 2.3 failed: ", c.value.(string), v.value)
		}
		e = e.Next()
	}

	lc.Delete(3)

	l = list.New()
	l.PushBack(&Elem{2, "7"})
	l.PushBack(&Elem{4, "5"})
	if lc.Len() != 2 {
		t.Error("case 3.1 failed")
	}

	e = l.Front()
	for c := lc.Head(); c != nil; c = c.Next() {
		v := e.Value.(*Elem)
		if c.key.(int) != v.key {
			t.Error("case 3.2 failed: ", c.key.(int), v.key)
		}
		if c.value.(string) != v.value {
			t.Error("case 3.3 failed: ", c.value.(string), v.value)
		}
		e = e.Next()
	}
}
