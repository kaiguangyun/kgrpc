package helper

import (
	"net"
)

// get computer macs
func GetLocalMacs() (macs []string, err error) {
	interfaces, err := net.Interfaces()

	if err != nil {
		return macs, err
	}

	for _, inter := range interfaces {
		mac := inter.HardwareAddr.String()
		macs = append(macs, mac)
	}
	return macs, err
}

// get computer ips
func GetLocalIps() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ips, err
	}

	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()
				ips = append(ips, ip)
			}
		}
	}
	return ips, err
}
