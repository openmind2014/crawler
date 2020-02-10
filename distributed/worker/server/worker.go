package main

import (
	"crawler/distributed/rpcsupport"
	"crawler/distributed/worker"
	"flag"
	"fmt"
	"log"
)

var port = flag.Int("port", 0, "The port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
