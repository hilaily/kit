package netx

import (
	"fmt"
	"net"
	"time"
)

func TCPPortIsOpen(host, port string, timeout int) (bool, error) {
	return portIsOpen("tcp", host, port, timeout)
}

func portIsOpen(typ, host, port string, timeout int) (bool, error) {
	addr := net.JoinHostPort(host, port)
	conn, err := net.DialTimeout(typ, addr, time.Duration(timeout)*time.Second)
	if nil != err {
		return false, fmt.Errorf("netx, TCPPortIsOpen, %w", err)
	}
	defer conn.Close()
	return true, nil
}
