package main

import (
	"log"
	"net/http"
)

type server struct{
  addr string
}

func(s *server) ServeHTTP(w http.ResponseWriter, r *http.Request){
  // Routing eg, / will return the index
  switch r.Method{
  case http.MethodGet:
    switch r.URL.Path {
    // index path
    case "/":
      w.Write([]byte("Index page"))
      return
    // user route
    case "/users":
      w.Write([]byte("Users page"))
      return
    // no route is found
    default:
      w.Write([]byte("404 not found"))
      return
    }
  }
}

func main(){
  s:= &server{addr:":8080"}
  err := http.ListenAndServe(s.addr,s) 
  
  if err != nil{
  log.Fatal(err)  
  }

}

// To run the server simply `go run main.go `
// The client will then connect to port 8080 and the server will respond
// For example we can call the server by doing `curl http://localhost:8080` and the server will respond


// 
