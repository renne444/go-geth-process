package util

import (
	"fmt"
	"os"
	"testing"
)

func TestPathDir(t *testing.T)  {
	emptyDir := "/home/yuki/dev/workplace/sol/doc/test"
	nonEmptyDir := "/home/yuki/dev/workplace/sol/doc"
	notDir := "/home/yuki/worplace"

	fmt.Println(FileExists(emptyDir))
	fmt.Println(FileExists(nonEmptyDir))
	fmt.Println(FileExists(notDir))


	fmt.Println(IsDir(emptyDir))
	fmt.Println(IsDir(nonEmptyDir))
	fmt.Println(IsDir(notDir))

	fmt.Println(IsNonEmptyDir(emptyDir))
	fmt.Println(IsNonEmptyDir(nonEmptyDir))
	fmt.Println(IsNonEmptyDir(notDir))

	fmt.Println(CreateFolder(notDir))
	os.Remove(notDir)
}
