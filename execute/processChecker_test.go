package execute

import (
	"fmt"
	"testing"
)

func TestComposeInitCmd(t *testing.T){
	var conn Connector
	config := conn.LoadConfig("")
	checker:= NewProcessChecker(config)

	initCmd := checker.composeInitCmd()
	fmt.Print(initCmd)
}

func TestInitialize(t *testing.T){
	var conn Connector
	config := conn.LoadConfig("")
	checker:= NewProcessChecker(config)

	checker.InitProcess()
}

func TestRun(t *testing.T){
	var conn Connector
	config := conn.LoadConfig("")
	checker:= NewProcessChecker(config)

	checker.RunProcess()
}

func TestExecute(t *testing.T){
	var conn Connector
	config := conn.LoadConfig("")
	checker:= NewProcessChecker(config)

	checker.Execute()
}
