package main

import "fmt"

func main() {
	s := NewServer()
	s.PrintOptions()

	s2 := NewServer(MaxCon(8), Timeout(-1))
	s2.PrintOptions()
}

func magic() {
	fun()
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

type ServerOption func(o *options) error

type Server struct {
	opts options
}

func (s Server) PrintOptions() {
	fmt.Println(s.opts.maxCon)
	fmt.Println(s.opts.transportType)
	fmt.Println(s.opts.timeout)
}

func NewServer(os ...ServerOption) Server {
	opts := defaultOptions
	for _, o := range os {
		o(&opts)
	}
	return Server{
		opts: opts,
	}
}

func MaxCon(n int) ServerOption {
	return func(o *options) error {
		o.maxCon = n
		return nil
	}
}

func Timeout(n int) ServerOption {
	return func(o *options) error {
		o.timeout = n
		return nil
	}
}
