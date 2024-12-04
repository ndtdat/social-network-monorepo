package suid

import "net"

// InterfaceAddrs defines the interface used for retrieving network addresses
type InterfaceAddrs func() ([]net.Addr, error)

var defaultInterfaceAddrs = net.InterfaceAddrs
