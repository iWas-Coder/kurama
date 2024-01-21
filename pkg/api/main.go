package main

import (
  "log"
  "net/http"
)

type server struct {}

func (self *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(`Hello, World!`))
}

func main() {
  s := &server{}
  http.Handle("/", s)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
