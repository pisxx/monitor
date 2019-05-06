package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	resp, err := client.Get("http://169.254.169.254/latest/meta-data/public-ipv4")
	fmt.Println(resp)
}
