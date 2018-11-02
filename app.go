// Package main provides ...
package main

import (
	"leaseapp/routers"
)

func main() {
	route := router.InitRouter()
	route.Run(":9000")
}
