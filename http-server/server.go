package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		str := `<h1 style="color:red;">今天天气不错</h1>`
		writer.Write([]byte(str))
	})
	http.ListenAndServe("localhost:8000", nil)
}
