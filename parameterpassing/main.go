package main

import "fmt"

func main() {
	s := NewServer()
	s.PrintOptions()

	s2 := NewServer(MaxCon(4), Transport(TCP))
	s2.PrintOptions()
}

type transport int
const (
	TCP transport = iota
	UDP
)

var defaultOptions = options{
	maxCon:        4,
	transportType: UDP,
	timeout:       3000,
}

type options struct {
	maxCon        int
	transportType transport
	timeout       int
}

type ServerOption func(o options) options

type Server struct {
	opts options
}

func (s Server) PrintOptions() {
	fmt.Println("Server options..")
	fmt.Printf("maxCon: %v\n", s.opts.maxCon)
	fmt.Printf("transport type: %v\n", ([]string{"TCP", "UDP"}[s.opts.transportType]))
	fmt.Printf("timeout: %v\n", s.opts.timeout)
}

func NewServer(os ...ServerOption) Server {
	opts := defaultOptions
	for _, o := range os {
		opts = o(opts)
	}
	return Server{
		opts: opts,
	}
}

func MaxCon(n int) ServerOption {
	return func(o options) options {
		o.maxCon = n
		return o
	}
}

func Timeout(n int) ServerOption {
	return func(o options) options {
		o.timeout = n
		return o
	}
}

func Transport(t transport) ServerOption {
	return func(o options) options {
		o.transportType = t
		return o
	}
}
