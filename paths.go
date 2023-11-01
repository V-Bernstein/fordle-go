package main

import "fmt"
//import "fordle-go/fordle/utils"
import "net/http"

func main() {
    http.HandleFunc("/words/validate/{word}", validateWord)
}


func validateWord(w http.ResponseWriter, req *http.Request) {
    fmt.Printf("%s", req)
    fmt.Fprintf(w, true)
}
