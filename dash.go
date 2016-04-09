package main

import (
	"flag"
	"fmt"
	"github.com/dominikh/arp"
	"os"
)

var (
	dash  string
	count int
	itf   string
	num   = 0
)

func handler(w arp.ResponseSender, r *arp.Request) {
	if r.SenderHardwareAddr.String() == dash && r.SenderIP.String() == "0.0.0.0" {
		num += 1
		fmt.Printf("Dash %d!\n", num)
		if count > 0 && num == count {
			os.Exit(0)
		}
	}
}

func main() {
	flag.StringVar(&dash, "d", "74:c2:46:fc:84:19", "MAC address of DashButton")
	flag.IntVar(&count, "c", 0, "number of clicks to handle, 0 for unlimited")
	flag.String(&itf, "i", "wlan0", "name of interface to listen on")
	flag.Parse()

	if *countFlag == 0 {
		fmt.Printf("Handling clicks on %s\n", dash)
	} else {
		fmt.Printf("Handling %d clicks on %s\n", count, dash)
	}

	e := arp.ListenAndServe(itf, arp.HandlerFunc(handler))

	if e.Error() == "operation not permitted" {
		fmt.Printf("Either first run 'sudo setcap cap_net_raw+pe goDash' or run as root\n")
	} else {
		fmt.Printf("Got error %s, aborting\n", e)
	}
}
