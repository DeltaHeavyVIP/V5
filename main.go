package main

import (
	"./entry"
	"fmt"
	"os"
)

func main() {
	data := entry.GetData()
	fmt.Fprintf(os.Stdout, "",data)
}
