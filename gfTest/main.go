package main

import (
	"github.com/fatih/color"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

func main() {
	//color.New(color.FgYellow).Print("oooooooo")
	//return
	s := g.Server()
	g.Log().Debug("this is log")
	g.Log().Info("this is info")
	g.Log().Notice("this is notice")
	g.Log().Warning("this is waning")
	g.Log().Info("this is blue info log")
	g.Log().Info("this is origin info log")
	//g.Log().Critical("this is critical")
	//g.Log().Error("this is err")
	//g.Log().Panic("this is panic, so break")
	//g.Log().Fatal("this is fatal, can not print")

	s.Group("/", func(group *ghttp.RouterGroup) {
		group.GET("/", func(r *ghttp.Request) {
			r.Response.Write("welcome to gf")
		})
	})
	s.Run()
}

func colorT()  {
	color.Red("1111")
}