package main

import "fmt"
import "os"

func main() {
    if (len(os.Args) < 3) {
	fmt.Println("Error, borked args")
	return
    }
    var s1 = os.Args[1]
    var s2 = os.Args[2]
    var isAnagram = isAnagram(s1, s2)
    fmt.Printf("Ana? %v\n", isAnagram)
   
    var searchSpace = []string{"ab","bc", "cd","de","fgh","fordle","lol"}
    var target = "fordl"
    var isInSpace = find(searchSpace, target)
    fmt.Printf("Is in space? %v\n", isInSpace)
}

func isAnagram(s1 string, s2 string) bool {
    if (len(s1) != len(s2)) {
	return false
    }

    var letters map[rune]uint8 = make(map[rune]uint8)

    for _,letter := range s1 {
	count, ok := letters[letter]
	if (ok) { /* Letter is already in map */
	    letters[letter] = count + 1
	} else {
	    letters[letter] = 1
	}
    } 

    for _,letter := range s2 {
	count, ok := letters[letter]
	if (ok) {
	    if (count == 0) {
	        return false /* No more letters remaining */
	    }
	    letters[letter] = count - 1
	} else {
	    return false /* This letter is not in s1 */
	}
    }
    return true
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

/* TODO: Convert the file into the string slice*/
func getWords() []string {
    return make([]string,1,2)
}
