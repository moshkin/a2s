package utils

import (
	"errors"
	"fmt"
	"net"
)

func IsValidIp(ip string) (bool, error) {
	if ip := net.ParseIP(ip); ip == nil {
		return false, errors.New(fmt.Sprintf("Invalid IP address {%s}", ip))
	}

	return true, nil
}

func IsValidPort(port int) (bool, error) {
	if port < 1 || port > 65535 {
		return false, errors.New(fmt.Sprintf("Port number {%d} is out of the valid range.", port))
	}

	return true, nil
}
