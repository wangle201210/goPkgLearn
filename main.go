package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"

	//"github.com/go-kratos/kratos/pkg/naming"
	//"github.com/go-kratos/kratos/pkg/naming/discovery"
)

func quickSort(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}
	mid := arr[0]
	var left, right []int
	for i := 1; i < l; i++ {
		if arr[i] < mid {
			left = append(left, arr[i])
		} else {
			right = append(right, arr[i])
		}
	}
	left = quickSort(left)
	right = quickSort(right)
	res := append(left, mid)
	return append(res, right...)
}

func quickTest(c int) {
	startAt := time.Now()
	var arr []int
	for i := 0; i < c; i++ {
		arr = append(arr, rand.Intn(c))
	}
	arr = quickSort(arr)
	for i := 0; i < c-1; i++ {
		if arr[i] > arr[i+1] {
			panic("验证未通过")
		}
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	timeUsed := time.Since(startAt) / time.Millisecond
	peak := m.Alloc / (1 << 20)
	fmt.Printf("%+v", m)
	fmt.Printf("验证通过！\n内存使用：%d,\n耗时：%d\n", peak, timeUsed)
}

func lenTest() {
	const s = "12345678"
	var a byte = 1 << len(s) / 128
	var b byte = 1 << len(s[:]) / 128
	const m = 8
	var n = 8
	var c byte = 1 << m / 128
	var d byte = 1 << n / 128
	print(a, b, c, d)
}

func floatTest()  {
	const percent = 0.1395
	res := percent * 100
	fmt.Printf("%T\n",res)
	fmt.Println(res)
}

//func reg() {
//	ip := "127.0.0.1"
//	port := "9000"
//	region := "regionName"
//	zone := "zoneName"
//	host := "hostName"
//	DeployEnv := "DeployEnvName"
//
//	dis := discovery.New(&discovery.Config{
//		Region: region,
//		Zone: zone,
//		Host: host,
//		Env: DeployEnv,
//		Nodes:[]string{"127.0.0.1:7171"},
//	})
//	ins := &naming.Instance{
//		Zone:     "zoneName",
//		Env:      "DeployEnvName",
//		AppID:    "infra.discovery1",
//		Hostname: "hostName",
//		Addrs: []string{
//			"grpc://" + ip + ":" + port,
//		},
//	}
//
//	cancel, err := dis.Register(context.Background(), ins)
//	if err != nil {
//		panic(err)
//	}
//
//	defer cancel()
//}
func main() {
	//reg()
	//floatTest()
	//lenTest()
	//quickTest(100000)
}
