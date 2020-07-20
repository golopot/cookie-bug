package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func main() {
	server := fasthttp.Server{
		Handler:                       requestHandler,
		DisableHeaderNamesNormalizing: true,
	}

	err := server.ListenAndServe(":8080")

	if err != nil {
		panic(err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx) {
	fmt.Println("== Header ==")
	fmt.Println(string(ctx.Request.Header.Header()))

	fmt.Println("== VisitAllCookie() ==")
	ctx.Request.Header.VisitAllCookie(func(k []byte, v []byte) {
		fmt.Println(string(k), string(v))
	})

	fmt.Println("== Header.Cookie() ==")
	fmt.Println("CookieFoo:", string(ctx.Request.Header.Cookie("foo")))
	ctx.WriteString("ok\n")
}
