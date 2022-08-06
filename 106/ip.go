package main

import (
	"fmt"
	"net"
)

func IPtest() {
	var ip net.IP
	b := ip.IsLoopback()
	fmt.Printf("b is %v\n", b)
}
func ParseIP() {
	//ipv4Addr := net.ParseIP("192.0.2.1")
	// This mask corresponds to a /24 subnet for IPv4.
	ipv4Mask := net.CIDRMask(24, 32)
	fmt.Println(ipv4Mask)

	ipv6Addr := net.ParseIP("2001:db8:a0b:12f0::1")
	// This mask corresponds to a /32 subnet for IPv6.
	ipv6Mask := net.CIDRMask(32, 128)
	fmt.Println(ipv6Addr.Mask(ipv6Mask))

	fmt.Println(net.IPv4Mask(255, 255, 255, 0))

	mask := net.ParseIP("255.255.255.0")
	fmt.Printf("mask is %v\n", mask)
	i, j := mask.MarshalText()
	fmt.Printf("i is %v, j is %v\n", i, j)

	var a, b, c, d byte
	fmt.Sscanf("255.255.255.0", "%d.%d.%d.%d", &a, &b, &c, &d)
	fmt.Printf("a is %v, b is %v[%v][%v]\n", a, b, c, d)
	mm := net.IPv4Mask(a, b, c, d)
	fmt.Printf("mm is %v\n", mm)
	m, n := mm.Size()
	fmt.Printf("m,n is %v,%v\n", m, n)
	fmt.Printf("---------IPtest------------\n")
	IPtest()
}
