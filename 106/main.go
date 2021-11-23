package main

import (
	"fmt"
	"net"
)

func main() {
	//ipv4Addr := net.ParseIP("192.0.2.1")
	// This mask corresponds to a /24 subnet for IPv4.
	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println(ipv4Mask)

	ipv6Addr := net.ParseIP("2001:db8:a0b:12f0::1")
	// This mask corresponds to a /32 subnet for IPv6.
	ipv6Mask := net.CIDRMask(32, 128)
	fmt.Println(ipv6Addr.Mask(ipv6Mask))

	fmt.Println(net.IPv4Mask(255, 255, 255, 0))

}
