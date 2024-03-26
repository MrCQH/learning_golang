package main

import (
	"io"
	"log"
	"net/http"
)

const form = `
    <html><body>
        <form action="#" method="post" name="bar">
            <input type="text" name="in" />
            <input type="submit" value="submit"/>
        </form>
    </body></html>
`

type HandleFnc func(http.ResponseWriter, *http.Request)

// 闭包增强方法，防止服务器发送错误
func logPanic(fn HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, req *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", req.RemoteAddr, x)
			}
		}()
		fn(writer, req)
	}
}

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data */
		//request.ParseForm();
		//io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, request.FormValue("in"))
		panic("This is an Error!!!")
	}
}
func main() {
	http.HandleFunc("/test1", logPanic(SimpleServer))
	http.HandleFunc("/test2", logPanic(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
