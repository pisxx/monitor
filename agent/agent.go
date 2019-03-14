package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/pisxx/monitor/agent/register"

	"github.com/pisxx/monitor/agent/web"
)

var ip = flag.String("ip", "0.0.0.0", "IP on which agent will listen")
var port = flag.String("p", "10808", "Port on which agent will listen")

func main() {
	flag.Parse()
	ipPort := *ip + ":" + *port
	// fmt.Print(flag.Args())
	// register.ConsulRegister("23", *ip)
	reg, err := register.RegisterAgent(*ip, *port)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(reg)
	http.HandleFunc("/", web.IndexMetrics)
	log.Printf("Listening on %s", ipPort)

	// http.ListenAndServe(ip + ":8080", handler http.Handler)
	log.Fatal(http.ListenAndServe(ipPort, nil))
}
