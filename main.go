package main

import (
	"flag"
	"fmt"
	process "gethProcess/execute"
)

func main(){

	configPath := flag.String("c","","config file path")
	fmt.Println("path = "+ *configPath)
	var conn process.Connector
	conn.Start(*configPath)
}
