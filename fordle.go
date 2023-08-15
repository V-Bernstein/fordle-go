package main

import "fmt"
import "os"
import "fordle-go/fordle/utils"

func main() {
    if (len(os.Args) < 3) {
	fmt.Println("Error, borked args")
	return
    }
    var s1 = os.Args[1]
    var s2 = os.Args[2]
    test := utils.IsAnagram(s1, s2)
    fmt.Printf("Ana? %v\n", test)
}
