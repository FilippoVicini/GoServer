package main

import (
  "encoding/json"
  "net/http"
)



type api struct {
  addr string
}

// In memory storage for users
var users = []User{}


func (s *api) ServeHTTP(w http.ResponseWriter, r *http.Request){
  switch r.Method{
  case http.MethodGet:
      switch  r.URL.Path {
       case "/":
         w.Write([]byte("Get Method"))
        case "/index":
          w.Write([]byte("Get index"))
        
      }
    case http.MethodPost:
      w.Write([]byte("Post method"))
   
  }
}

func (s *api) getUserHandler(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Users list"))
}

func (s *api) createUserHandler(w http.ResponseWriter, r *http.Request){
  w.Write([]byte("Created user"))
}
