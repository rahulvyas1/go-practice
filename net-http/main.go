package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/users", usersHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func usersHandleFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("We got a request on /users")
	fmt.Fprint(w, "Hi, thanks for calling my /users API with HTTP Method '%v' ", r.Method)
}
