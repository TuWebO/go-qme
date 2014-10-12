// GoQme qsystem that a queue is working on. client library

package qsystem

// @TODO should be this an interface implemented by queues, services and channels?
type Qsystem struct {
  ArrivingPathern uint // How a client comes (i.e. one by one, by lotes...)
  ArrivingTime    uint // Time in secconds between the last in the queue and the new one
  ResponseTime    uint // Time in secconds between the last exited client and the next one
  Lambda          int  // Number of arrivings by unit time
  Mu              int  // Available services by unit time
}
