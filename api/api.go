package main

import (
	"encoding/json"
	"log"
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
  // Setting header letting client know we will be using json
  w.Header().Set("Content-Type", "application/json")
  
  // Json package from go to encode (receives a writer)
  // encode users into json

  err:= json.NewEncoder(w).Encode(users)
  if err != nil{
    log.Fatal(err)
    return
  }

  // Return another header with OK status
  w.WriteHeader(http.StatusOK)
}


func (s *api) createUserHandler(w http.ResponseWriter, r *http.Request){

  var payload User
  // decoding again the json
  err:= json.NewDecoder(r.Body).Decode(&payload)
  if err != nil{
    log.Fatal(err)
    return
  }

  // creating a new user
  u := User{
    FirstName :payload.FirstName,
    LastName: payload.LastName,
  }
  // appending it to the local list of users 
  users = append(users, u)

  w.WriteHeader(http.StatusCreated)
}
