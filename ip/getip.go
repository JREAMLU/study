package ip

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

const (
	LO       = "lo"
	LOOPBACK = "loopback"
	ETH0     = "eth0"
	ETH1     = "eth1"
)

// func main() {
// 	ip, err := externalIP()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(ip)
// }

func GetServerIP() string {
	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}

	var ipMap = make(map[string]string)
	ipMap[LO] = ""
	ipMap[LOOPBACK] = ""
	ipMap[ETH0] = ""
	ipMap[ETH1] = ""

	for _, iface := range list {
		// fmt.Printf("%d name=%s %v\n", i, iface.Name, iface)
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		for _, addr := range addrs {
			// fmt.Printf(" %d %v\n", j, addr)
			name := strings.Split(strings.ToLower(iface.Name), " ")
			ip := strings.Split(addr.String(), "/")

			switch name[0] {
			case LO:
			case LOOPBACK:
			case ETH0:
			case ETH1:
			}
			if name[0] == LO || name[0] == LOOPBACK || name[0] == ETH0 || name[0] == ETH1 {
				ipMap[name[0]] = ip[0]
			}
		}
	}

	for k, v := range ipMap {
		fmt.Println(k, ":", v)
	}

	if ipMap[ETH0] != "" {
		return ipMap[ETH0]
	}
	if ipMap[ETH1] != "" {
		return ipMap[ETH1]
	}
	if ipMap[LO] != "" {
		return ipMap[LO]
	}
	if ipMap[LOOPBACK] != "" {
		return ipMap[LOOPBACK]
	}

	return "127.0.0.1"
}

func ExternalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

func Add(i int) int {
	return i + 1
}
