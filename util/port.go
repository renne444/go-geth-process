package util

import (
	"net"
	"strconv"
	"time"
)

func PortInUse(port int)bool{
	conn,_ := net.DialTimeout("tcp","127.0.0.1:"+strconv.Itoa(port),time.Second)
	if conn != nil {
		conn.Close()
		return true
	}

	return false
}
