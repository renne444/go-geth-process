package util

import (
	"fmt"
	"testing"
)

func TestPortInUse(t *testing.T){
	fmt.Printf("port 63342 using status : %t\n", PortInUse(63342))
	fmt.Printf("port 62242 using status : %t\n", PortInUse(62242))
}