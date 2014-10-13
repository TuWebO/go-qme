// GoQme queue-queue client library
package qq

import (
	"fmt"
)

type Queue struct {
	Name string
}

type QueueInfo struct {
	Id               string
	Discipline       string // FIFO, LIFO, Proc sharing, Priority, Shortest job first...
	IsPreemptive     bool   // Client without priority exit and wait
	PreservePosition bool   // If IsPreemptive the client can go back to its older position
}

func New(queueName string) *Queue {
	// @TODO Should we allow queues without a Name?
	return &Queue{Name: queueName}
}

// Implement as a Stringer
func (q Queue) String() string {
	return fmt.Sprintf("Id %v", q.Name)
}

/*func (q *Queue) DrawQueue() uint64 {
  //fmt.Println("My Id is:   ", q.id)
  fmt.Println("My Name is: ", q.Name)
  return q.id
}*/
