package execute

import (
	"fmt"
	"testing"
)

func TestLoadConfig(t *testing.T){
	var conn Connector
	conf := conn.LoadConfig("")
	fmt.Print(conf)
}
