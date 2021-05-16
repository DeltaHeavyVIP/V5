package main

import (
	"./entry"
	"./methods"
)

func main() {
	data := entry.GetData()
	lag := methods.Lagrange{}
	lag.Count(data)
	newton := methods.Newton{}
	newton.Count(data)
}
