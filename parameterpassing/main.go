package main

import "fmt"

func main() {
	s := NewServer()
	s.PrintOptions()

	s2 := NewServer(MaxCon(8), Timeout(-1))
	s2.PrintOptions()
}

var defaultOptions = options{
	maxCon:        4,
	transportType: "UDP",
	timeout:       3000,
}

type options struct {
	maxCon        int
	transportType string
	timeout       int
}

type ServerOption func(o options) options

type Server struct {
	opts options
}

func (s Server) PrintOptions() {
	fmt.Println("Server options..")
	fmt.Println(s.opts.maxCon)
	fmt.Println(s.opts.transportType)
	fmt.Println(s.opts.timeout)
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
