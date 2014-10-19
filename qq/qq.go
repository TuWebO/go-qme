// Package qq provides a simple interface to manage queues
package qq

import (
	"fmt"
	"github.com/TuWebO/go-qme/qc"
)

// Queue represents a queue where a client waits.
type Queue struct {
	id 				 string
	name 			 string
	discipline       string // FIFO, LIFO, Proc sharing, Priority, Shortest job first...
	isPreemptive     bool   // Client without priority exit and wait
	preservePosition bool   // If IsPreemptive the client can go back to its older position
	clients		   	 []qc.Client
}

func NewQueue() *Queue {
	q := &Queue{
		id: "miId",
		name: "Myname",
		discipline: "FIFO",
		isPreemptive: true,
		preservePosition: true,
		clients: make([]qc.Client, 0),
	}
	return q
}

// Creates a new Queue
func (q *Queue) PrintName() {
	fmt.Println("My Queue name is :", q.name)
}

func (q *Queue) AddClient() {
	var c qc.Client
	c.Id = "344"
	q.clients = append(q.clients, c)
	for _, client := range q.clients {
		
		fmt.Println("My client's Id is: ", client.Id)		
	}
	fmt.Println("My clients are : ", q, c)
}
/*func New(queueName string) *Queue {
	// @TODO Should we allow queues without a Name?
	return &Queue{Name: queueName}
}

// Implement as a Stringer
func (q Queue) String() string {
	return fmt.Sprintf("Id %v", q.Name)
}*/

/*func (q *Queue) DrawQueue() uint64 {
  //fmt.Println("My Id is:   ", q.id)
  fmt.Println("My Name is: ", q.Name)
  return q.id
}*/
