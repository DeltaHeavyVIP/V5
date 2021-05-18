package main

import (
	"./entry"
	"./methods"
)

func main() {
	data := entry.GetData()
	lag := methods.Lagrange{data}
	lag.Count()
	newton := methods.Newton{data}
	newton.Count()
}
