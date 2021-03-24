package main

import (
	"wkb_comments/conf"
	"wkb_comments/routes"
)
func main() {
	conf.Init()
	r := routes.NewRouter()
	_ = r.Run(":8080") // listen and serve on 0.0.0.0:8080
}