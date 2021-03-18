package netx

import (
	"fmt"
	"net"
	"time"
)

func TCPPortIsOpen(addr string, timeout int) (bool, error) {
	return portIsOpen("tcp", addr, timeout)
}

func portIsOpen(typ, addr string, timeout int) (bool, error) {
	conn, err := net.DialTimeout(typ, addr, time.Duration(timeout)*time.Second)
	if nil != err {
		return false, fmt.Errorf("netx, TCPPortIsOpen, %w", err)
	}
	defer conn.Close()
	return true, nil
}
