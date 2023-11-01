package utils

import (
    "bufio"
    "log"
    "os"
)

func IsAnagram(s1 string, s2 string) bool {
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

func Find(space []string, target string) bool {
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

/* TODO: Return an error object as well; offload error handling to caller`*/
func GetWords() []string {
    filename := "words"
    file, err := os.Open(filename)
    if err != nil {
	log.Fatal(err)
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
 	lines = append(lines, scanner.Text())
    }

    if err := scanner.Err(); err != nil {
	log.Fatal(err)
    }
    return lines
}
