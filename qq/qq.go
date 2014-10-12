// GoQme queue-queue client library
package qq


import (
  "fmt"
)

type Queue struct {
  id   uint64
  name string
}

// Implement as a Stringer
func (q Queue) String() string {
  return fmt.Sprintf("Id %v | Name %v", q.id, q.name)
}

// Helpers
func NewQueue(id uint64, name string) Queue {
  // @TODO Should we allow queues without a Name?
  return Queue{id, name}
}

func (q *Queue) DrawQueue() uint64 {
  fmt.Println("My Id is:   ", q.id)
  fmt.Println("My Name is: ", q.name)
  return q.id
}

func (q *Queue) Setid(id uint64) {
  q.id = id
}
