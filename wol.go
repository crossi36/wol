package main

import (
	"fmt"
	"net"
	"os"

	"github.com/crossi36/wol/magicpacket"
)

func main() {
	var localAddr *net.UDPAddr

	macAddr := os.Args[1]
	bcastAddr := "255.255.255.255:7"

	udpAddr, err := net.ResolveUDPAddr("udp", bcastAddr)
	if err != nil {
		panic(err)
	}

	mp, err := magicpacket.New(macAddr)
	if err != nil {
		panic(err)
	}

	bs, err := mp.Marshal()
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", localAddr, udpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	n, err := conn.Write(bs)
	if err == nil && n != 102 {
		err = fmt.Errorf("magic packet sent was %d bytes (expected 102 bytes sent)", n)
	}
	if err != nil {
		panic(err)
	}

	fmt.Printf("Magic packet sent successfully to %s\n", macAddr)
}
