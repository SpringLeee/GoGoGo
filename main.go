package main

import (
	"fmt"
	"net/http"
)

func Middleware1(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("M1 in")
		next(w, r)
		fmt.Println("M1 out")
	}
}

func Middleware2(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("M2 in")
		next(w, r)
		fmt.Println("M2 out")
	}
}

func Middleware3(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("M3 in")
		next(w, r)
		fmt.Println("M3 in")
	}
}

type middleware func(http.Handler) http.Handler
type chain struct {
	middlewares []middleware
}

func Pipeline(next http.HandlerFunc) http.HandlerFunc {

	return Middleware1(Middleware2(Middleware3(next)))

}

func Pipeline1(middlewares ...middleware) chain {

	//return

}

func (c chain) Then(next http.HandlerFunc) http.HandlerFunc {

	// for i := range c.middlewares {
	// 	d :=c.middlewares[i]
	//   	d(next)
	// }

	return next

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Login...")
	w.Write([]byte("Login..."))

}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Register...")
	w.Write([]byte("Register..."))

}

func main() {

	http.HandleFunc("/Login", Pipeline(LoginHandler))

	http.HandleFunc("/Register", Pipeline1(Middleware1, Middleware2).Then(RegisterHandler))

	http.ListenAndServe(":9099", nil)

}
