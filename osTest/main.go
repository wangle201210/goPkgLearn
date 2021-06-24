package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func main()  {
	//imgUrl := "https://panel-file.qa.medlinker.com/meeting/background.png"
	//imgUrl := "Users/med/mine/goPkgLearn/osTest/a.txt"
	imgUrl := "asbnf"
	u, err := url.Parse(imgUrl)
	if err != nil {
		println("1",err.Error())
		return
	}
	println(u.Host)
	return
	response, err := http.Get(imgUrl)
	if err != nil {
		println("1",err.Error())
		return
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		println("not ok")
	}
	println(ioutil.ReadAll(response.Body))
	//
	////imgUrl := "/Users/med/mine/goPkgLearn/osTest/a.txt"
	//open, err := os.Open(imgUrl)
	//if err != nil {
	//	println(err.Error())
	//	println(6)
	//	return
	//}
	//defer open.Close()
	//s := make([]byte,1024)
	//_,err = open.Read(s)
	//fmt.Println(err)
	//println("a:",s)
	////获取远端图片
	//res, err := http.Get(imgUrl)
	//if err != nil {
	//	fmt.Println("A error occurred!")
	//	return
	//}
	//defer res.Body.Close()
	//
	//// 读取获取的[]byte数据
	//data, _ := ioutil.ReadAll(res.Body)
	//println(data)
	//imageBase64 := base64.StdEncoding.EncodeToString(data)
	//fmt.Println("base64", imageBase64)
}
