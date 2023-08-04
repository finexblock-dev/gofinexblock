package goaws

import (
	"fmt"
	"net"
)

func OwnPrivateIP() (ip string, err error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	for _, i := range interfaces {
		addresses, err := i.Addrs()
		if err != nil {
			panic(err)
		}

		for _, addr := range addresses {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			if ip != nil && ip.IsPrivate() {
				return ip.String(), nil
			}
		}
	}

	return "", fmt.Errorf("no private ip found")
}