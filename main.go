package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type target struct {
	hostname string
	port     int
}

func polltarget(prot string, target string) (string, bool) {
	//var status string
	//var err error
	conn, err := net.Dial(prot, target)
	if err != nil {
		fmt.Println(err)
		return "failed to dial", false
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return "failed to write to target", false
	}
	return status, true

}

func main() {

	var h = flag.String("target", "golang.org", "target of the probe")
	var p = flag.String("port", "80", "TCP port to check")
	var x string

	flag.Parse()

	// parse port here
	// test for , or - , else single port

	var ports []string

	if strings.Contains(*p, "-") {

	}

	if strings.Contains(*p, ",") {

		ports = strings.Split(*p, ",")

	}

	for _, port := range ports {
		t := target{hostname: *h}

		//t.hostname = *h
		t.port, _ = strconv.Atoi(port)

		fmt.Println(t.hostname)
		fmt.Println(t.port)
		// x = *t + ":" + *p

		x = t.hostname + ":" + strconv.Itoa(t.port)

		status, result := polltarget("tcp", x)
		if result {
			fmt.Println(status)
		}
		fmt.Println(status)
	}
}
