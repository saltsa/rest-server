package main

import (
	"fmt"
	"net"
)

// findListener creates a listener.
func findListener(addr string) (listener net.Listener, err error) {
	// listen manually
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("listen on %v failed: %w", addr, err)
	}

	log.Infof("start server on %v", addr)
	return listener, nil
}
