// Monitor agent
// Gets simple metrics from OS on which it is running those can be accessd via http page
// Listens on 0.0.0.0:10808 by default
// After startup it will register in AWS DynamoDB providing hostname, os, ip and port
//

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
var env = flag.String("e", "docker", "runing on docker or VM")

func main() {
	flag.Parse()
	ipPort := *ip + ":" + *port
	// fmt.Print(flag.Args())
	// register.ConsulRegister("23", *ip)
	reg, err := register.RegisterAgent(*ip, *port, *env)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(reg)
	http.HandleFunc("/", web.IndexMetrics)
	log.Printf("Listening on %s", ipPort)

	// http.ListenAndServe(ip + ":8080", handler http.Handler)
	log.Fatal(http.ListenAndServe(ipPort, nil))
}
