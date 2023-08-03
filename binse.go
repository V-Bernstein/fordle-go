package main

import "fmt"

func main() {
    var searchSpace = []string{"ab","bc", "cd","de","fgh","fordle","lol"}
    var target = "fordl"
    var isInSpace = find(searchSpace, target)
    fmt.Printf("Is in space? %v\n", isInSpace)
}

func find(space []string, target string) bool {
    var high = len(space) - 1
    var low = 0
    var mid = high / 2

    for low <= high {
        var word = space[mid]
	if (target == word) {
	    return true
	} else if (word < target) {
	    low = mid + 1
	} else {
	    high = mid - 1
	}	
	mid = (high + low) / 2
    }
    return false
}
