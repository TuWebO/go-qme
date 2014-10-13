// GoQme qs. System enviroment where clients are.
package qs

import (
	"fmt"
	//"github.com/TuWebO/go-qme/qq"
)

type Qsystem struct {
	Id              uint
	ArrivingPathern uint      // How a client comes (i.e. one by one, by lotes...)
	ArrivingTime    uint      // Time in secconds between the last in the queue and the new one
	ResponseTime    uint      // Time in secconds between the last exited client and the next one
	Lambda          int       // (λ) Number of arrivings by unit time
	Mi              int       // (μ) Number of services by unit time (if the service is occupied)
	Rho             int       // (ρ) System congestion
	Services        []Service // Service
	ClientCapacity  int       // Number of clients a system could have
}

// A service has servers (channels). As a matter of rule, these servers will run on
// parallel (they can attend any client in the queue), but not always (i.e: A
// supermarket that opens a new cashier for only those people who have less than 10 items.)
// A multichannel system with one queue by channel will behave as a multiple system
// with one server each (monochannel system)
type Service struct {
	Id       string
	IsActive bool
	Servers  []Server
	//Queues   []Queue
}

type Server struct {
	Id       string
	IsActive bool
}
