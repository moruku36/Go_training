package main

import (
	"fmt"
	"net/http" //HTTPプロトコルを利用してくれるパッケージ
)

func handler(writer http.ResponseWriter, request *http.Request) {
	//Hello worldを渡してあげるよ
	fmt.Fprintf(writer, "Hello World!,%s", request.URL.Path[1:])
}

func main() {
	//パスを指定してどういった動きをするかハンドリングする
	http.HandleFunc("/", handler)
	//サーバーを自分のPCの中で立ち上げている。ポート8080として立ち上げる
	http.ListenAndServe(":8080", nil)
}
