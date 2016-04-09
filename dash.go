package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/dominikh/arp"
)

var (
	dashFlag = flag.String("d", "74:c2:46:fc:84:19", "MAC address of DashButton")
	countFlag = flag.Int("c", 0, "number of clicks to handle, 0 for unlimited")
	itfFlag = flag.String("i", "wlan0", "name of interface to listen on")
	num = 0
)

func handler(w arp.ResponseSender, r *arp.Request) {
	if (r.SenderHardwareAddr.String() == *dashFlag && r.SenderIP.String() == "0.0.0.0") {
		num += 1
		fmt.Printf("Dash %d!\n", num)
		if (*countFlag > 0 && num == *countFlag) {
			os.Exit(0)
		}
	}
}

func main() {
	flag.Parse()
	if (*countFlag == 0) {
		fmt.Printf("Handling clicks on %s\n", *dashFlag)
	} else {
		fmt.Printf("Handling %d clicks on %s\n", *countFlag, *dashFlag)
	}
	e := arp.ListenAndServe(*itfFlag, arp.HandlerFunc(handler))
	if (e.Error() == "operation not permitted") {
		fmt.Printf("Either first run 'sudo setcap cap_net_raw+pe goDash' or run as root\n")
	} else {
		fmt.Printf("Got error %s, aborting\n", e)
	}
}
	

	
