package main

import "net/http"



func main(){
  api := &api{addr: ":8080"}


  // Mux is a router
  mux := http.NewServeMux()


  srv := &http.Server{
    Addr: api.addr,
    Handler: mux,
  }

  //
  mux.HandleFunc("GET /users", api.getUserHandler)
  srv.ListenAndServe()
}
