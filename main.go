package main

import (
	"fmt"
	"net/http"
)

func Middleware1(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("M1 in")
		next.ServeHTTP(w, r)
		fmt.Println("M1 out")
	})

}

func Middleware2(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("M2 in")
		next.ServeHTTP(w, r)
		fmt.Println("M2 out")
	})

}

func Middleware3(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("M3 in")
		next.ServeHTTP(w, r)
		fmt.Println("M3 out")
	})

}

type Chain struct {
	middlewares []func(handler http.Handler) http.Handler
}

func Pipeline(next http.Handler) http.Handler {

	//return Middleware1(Middleware2(Middleware3(next)))

	return AddMiddlewares(Middleware1, Middleware2, Middleware3).Then(next)

}

func AddMiddlewares(m ...func(handlerFunc http.Handler) http.Handler) Chain {

	c := Chain{}

	c.middlewares = append(c.middlewares, m...)

	return c

}

func (c Chain) Then(next http.Handler) http.Handler {

	for i := range c.middlewares {

		prev := c.middlewares[len(c.middlewares)-1-i]

		next = prev(next)
	}

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

	http.Handle("/Login", Pipeline(http.HandlerFunc(LoginHandler)))

	http.Handle("/Register", http.HandlerFunc(RegisterHandler))

	http.ListenAndServe(":8080", nil)

}
