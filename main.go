package main

import (
	"fmt"
	"github.com/ofer987/rename/fileutils"
	"os"
)

func main() {
	DisplayHello()

	os.Stdout.WriteString("The source file is:")
	fmt.Printf("%+v\n", os.Args[1])
	src := os.Args[1]
	dest := src
	fmt.Scanf("%s", &dest)
	fmt.Printf("the new file name is %+v\n", dest)

	fileutils.MoveFile(src, dest)
}

func DisplayHello() {
	fmt.Printf("Hello World\n")
}
