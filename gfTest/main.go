package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", func(r *ghttp.Request) {
			r.Response.Write("welcome to gf")
		})
	})
	s.Run()
}
