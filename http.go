// a go program to illustrate of HTTP request
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func main() {
    // Define the URL to send the GET request
    url := "https://example.com" // Replace with the URL you want to request

    // Send the GET request
    response:= http.Get(url)
   
   // defer response.Body.Close()

    // Check the HTTP status code
  

    // Read the response body
    body := ioutil.ReadAll(response.Body)
  

    // Print the response body as a string
    fmt.Println("Response:")
    fmt.Println(string(body))
}
