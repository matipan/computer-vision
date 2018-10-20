package main

import "image"

type mqueue struct {
	data []image.Point
	size int
}

// New creates a new mqueue with the specified size.
func New(size uint) *mqueue {
	return &mqueue{
		data: []image.Point{},
		size: int(size),
	}
}

func (q *mqueue) Clear() {
	q.data = []image.Point{}
}

func (q *mqueue) Push(p image.Point) {
	if len(q.data) == q.size {
		q.data = q.data[1 : q.size-1]
	}
	q.data = append(q.data, p)
}

func (q *mqueue) Range(f func(p image.Point)) {
	for _, p := range q.data {
		f(p)
	}
}

func (q *mqueue) RangeWithPrevious(f func(current image.Point, previous image.Point)) {
	for i := 1; i < len(q.data); i++ {
		f(q.data[i], q.data[i-1])
	}
}
