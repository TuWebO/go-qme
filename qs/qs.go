// GoQme queue-service client library
package qs

// A service has channels. As a matter of rule, these channels will run on
// parallel, if not, the channel(s) that can't run in parallel should have its
// own service (i.e: A supermarket that opens a new cashier for only those people
// who have less than 10 items.)
type Service struct {
	Channels []Channel // Number of "services (channles)" this service has
}

type Channel struct {
	Id				string
	IsActive	bool
}

// @TODO we should check how many active channels exists
func (s Service) MultiChannelActive() (bool) {
	if len(s.Channels) > 1 {
		return true
	}
	return false
}
