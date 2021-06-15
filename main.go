package main

import "calculator/router"

func main() {
	r := router.SetUpRoute()
	r.Run(":9000")
}
