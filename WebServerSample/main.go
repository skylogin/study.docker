package main

import (
  "fmt"
  "log"
  "net/http"
)

func main(){
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
    log.Println("received request")
    fmt.Fprintf(w, "Hello Docker!!")
  })

  log.Println("start server")
  // server := &http.Server{
  //   Addr: ":8080"
  // }
  if err := http.ListenAndServe(":8080", nil); err != nil{
    log.Println(err)
  }
}
