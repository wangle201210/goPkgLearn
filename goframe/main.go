package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)


func handleFun(r *ghttp.Request) {
	r.Response.Write("here is localhost")
}


func handleFunAbc(r *ghttp.Request) {
	r.Response.Write("here is abc")
}

func handleFunCanAbc(r *ghttp.Request) {
	r.Response.Write("here is can abc")
}

type test struct {
}

func (s *test) Index(r *ghttp.Request) {
	r.Response.Write("index")
}


func (s *test) Show(r *ghttp.Request) {
	r.Response.Write("show")
}

func main() {
	s := g.Server()
	// 多域名绑定同路由做不同的事情
	s.Domain("localhost").BindHandler("/", handleFun)
	s.Domain("127.0.0.1").BindHandler("/", handleFun)
	// 参数获取
	s.BindHandler("/{class}-{course}/:name/*do", func(r *ghttp.Request){
		r.Response.Writef("%s\n %s\n %s\n %s\n %s\n",
			r.Router.Uri,
			r.Get("class"),
			r.Get("course"),
			r.Get("name"),
			r.Get("do"),
		)
	})
	// gin 不支持这种覆盖的写法
	// beego和gf 支持这种覆盖的写法
	s.BindHandler("/used/abc",handleFunAbc)
	s.BindHandler("/used/*abc",handleFunCanAbc)
	// 牛皮的写法
	s.BindObject("/{.struct}-{.method}",&test{})
	//s.BindHandler("/test/index",&test{}.Index)

	s.SetPort(9999)
	s.Run()
}
