package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("HelloServer handler")
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", req.URL.Path[1:], "hello")
}

type Hello struct {
	data string
}

func (h *Hello) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println("HelloServer handler")
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", req.URL.Path[1:], h.data)
}

func main() {
	//test1()
	//test2()
	test3()
}

func test1() {
	// 用struct实现
	//hello := &Hello{"hello"}
	//err := http.ListenAndServe(":8080", hello)
	// 用函数实现
	err := http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}

func test2() {
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.bilibili.com/",
		"http://www.weibo.com/",
	}
	for _, url := range urls {
		resp, err := http.Head(url)
		checkError(err, "url")
		fmt.Println(url, ": ", resp.Status)
	}
}

func test3() {
	resp, err := http.Get("http://www.baidu.com/")
	checkError(err, "Http Get")
	bodyData, err := io.ReadAll(resp.Body)
	checkError(err, "IO ReadAll")
	fmt.Printf("Get: %q", string(bodyData))
}

func checkError(err error, info string) {
	if err != nil {
		fmt.Println("Error ", ": ", info, ", ", err.Error())
	}
}
