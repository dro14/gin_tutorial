package main

// import (
//   "fmt"
//   "strings"
//   "net/http"
//   "io/ioutil"
// )

// func main() {

//   url := "http://localhost:8080/videos"
//   method := "GET"

//   payload := strings.NewReader(`{
//     "title": "Golang / Go Gin Framework Crash Course 02 | Middlewares 101: Authorization, Logging and Debugging",
//     "description": "In this video we are going to start working with Middlewares providing Authorization, Logging and Debugging capabilities to our API using Golang's Gin HTTP Framework.",
//     "url": "https://youtu.be/Ypwv1mFZ5vU"
// }`)

//   client := &http.Client {
//   }
//   req, err := http.NewRequest(method, url, payload)

//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
//   req.Header.Add("Content-Type", "text/plain")

//   res, err := client.Do(req)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   defer res.Body.Close()

//   body, err := ioutil.ReadAll(res.Body)
//   if err != nil {
//     fmt.Println(err)
//     return
//   }
//   fmt.Println(string(body))
// }