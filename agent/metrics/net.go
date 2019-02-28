package metrics

// import (
// 	"fmt"
// 	"log"
// 	"net"
// )

// func GetNetwork() {
// 	iface, err := net.InterfaceByName("en0")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Print(iface.Addrs())
// 	var ifaceIP net.IP
// 	addrs, _ := iface.Addrs()
// 	for _, addr := range addrs {
// 		var ip net.IP
// 		switch v := addr.(type) {
// 		case *net.IPNet:
// 			ifaceIP = v.IP
// 		case *net.Addr:
// 			ifaceIP = v.IP
// 		}
// 	}
// 	fmt.Print(ifaceIP)
// }
