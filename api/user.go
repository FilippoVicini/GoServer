
package main

type User struct{
  // auto marshalling by the json
  FirstName string `json:"first_name"`
LastName string `json:"last_name"`
}
