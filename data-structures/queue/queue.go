package queue

import "github.com/midnight-vivian/go-data-structures/data-structures/container"

type Queuer interface {
	Enqueue(v interface{})
	Dequeue() interface{}
	container.Container
}
