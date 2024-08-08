package main

import "fmt"

type IPAddr [4]byte

func (receiver IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", receiver[0], receiver[1], receiver[2], receiver[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
