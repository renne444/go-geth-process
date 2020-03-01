package main

import (
	"flag"
	"fmt"
	process "gethProcess/execute"
)

func main(){

	var configPath string
	flag.StringVar(&configPath,"c","","config file path")
	flag.Parse()

	fmt.Println("path = "+ configPath)
	var conn process.Connector
	conn.Start(configPath)
}
