package main

import (
	"./entry"
	"fmt"
	"os"
)

func main() {
	entry.GetData()
	fmt.Fprintf(os.Stdout, "Некорректно введено значение x")
}
