package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fvbock/endless"
)

var addr = ":8888"

// 开启服务
func server() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("hello world")
	})
	fmt.Println("监听了8888端口")
	http.ListenAndServe(":8888",nil)
}

// 带中间件的服务
func serverWithMiddleWare()  {
	http.Handle("/foo", middleWare(http.HandlerFunc(router1)))
	port := 8888
	serverPort := fmt.Sprintf(":%d",port)
	fmt.Printf("服务监听了%s\n",serverPort)
	http.ListenAndServe(serverPort, nil)
}

func middleWare(m http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 中间件里面做的内容
		w.Write([]byte("hello world \n"))
		// 继续本身的方法里面的内容
		m.ServeHTTP(w,r)
	})
}

func router1(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("router1"))

	fmt.Println("路由1")
}

func fileServer() {
	// todo 路由改成其他的就不行？？
	http.Handle("/",http.FileServer(http.Dir("../")))
	fmt.Println("监听了8888端口")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		fmt.Println(err)
	}
}
func simpleGet() {
	// 要加http(s)
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	body := resp.Body
	defer body.Close()
	content, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}

func customGet() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		fmt.Println(err)
	}
	// 自己加些东西
	req.Header.Add("Auth","wanna")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := resp.Body
	defer body.Close()
	content, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(content))
}


// 优雅的热重启
func artisanShutdown() {
	r := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("yes i do"))
	})
	s := &http.Server{
		Addr: addr,
		Handler: r,
	}
	go func() {
		log.Println("start")
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()
	q := make(chan os.Signal)
	// 监听系统的interrupt（退出操作）
	signal.Notify(q,os.Interrupt)
	<-q
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Println("shutdown",err)
	}
	// 每次死掉后就重新唤醒一个
	for  {
		ns := endless.NewServer(addr, r)
		ns.BeforeBegin = func(add string) {
			log.Printf("Actual pid is %d", syscall.Getpid())
		}
		err := ns.ListenAndServe()
		if err != nil {
			log.Printf("Server err: %v", err)
		}
		log.Println("server exit")
	}

}
func main() {
	//server()
	//serverWithMiddleWare()
	//fileServer()
	//simpleGet()
	//customGet()
	artisanShutdown()
}

