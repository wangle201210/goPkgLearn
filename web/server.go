package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type recvData struct {
	m sync.RWMutex
	data map[string]struct{}
}

var rc recvData

func Run() {
	log.Println("run start")
	http.HandleFunc("/index", Index)
	http.HandleFunc("/send", Send)
	rc = recvData{
		data: make(map[string]struct{}),
	}
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Println("run end with err:", err)
		return
	}
	return
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("now in index")
	get := r.URL.Query().Get("data")
	if len(get) == 0 {
		get = "nothing"
	}
	w.Write([]byte(get))
	return
}

type sendData struct {
	Data string `json:"data"`
}

type respData struct {
	Errcode int64       `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func Send(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	readAll, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("read body err:", err.Error())
		return
	}
	data := new(sendData)
	if err = json.Unmarshal(readAll, data); err != nil {
		log.Println("Unmarshal body err:", err.Error())
		return
	}
	resp := new(respData)
	if have(data.Data) {
		resp.Data = "exist"
	} else {
		save(data.Data)
	}
	marshal, err := json.Marshal(resp)
	if err != nil {
		log.Println("Marshal err:", err.Error())
		return
	}
	w.Write(marshal)
	return
}

func have(s string) bool {
	rc.m.RLock()
	defer rc.m.RUnlock()
	_, e := rc.data[s]
	return e
}

func save(s string) {
	rc.m.Lock()
	defer rc.m.Unlock()
	rc.data[s] = struct{}{}
	return
}

func clean()  {
	rc.m.Lock()
	defer rc.m.Unlock()
	rc.data = map[string]struct{}{}
	return
}