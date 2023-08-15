package utils

import "testing"

func TestIsAnagram_unequal_length(t *testing.T) {
    ana := IsAnagram("salve", "lol")
    if ana {
        t.Fatal("Unequal length words aren't anagrams!")
    }
}

func TestIsAnagram_missing_letter(t *testing.T) {
    ana := IsAnagram("salve", "seals")
    if ana {
        t.Fatal("These words aren't anagrams!")
    }
}

func TestIsAnagram_true(t *testing.T) {
    ana := IsAnagram("listen", "silent")
    if !ana {
        t.Fatal("These should be anagrams!")
    }
}

func TestFind_empty_space(t *testing.T) {
    searchSpace := []string{}
    target := "fordl"
    found := Find(searchSpace, target) 
    if found {
        t.Fatal("Search space is empty!")
    }
}

func TestFind_not_in_space(t *testing.T) {
    searchSpace := []string{"ab","bc", "cd","de","fgh","fordle","lol"}
    target := "ford"
    found := Find(searchSpace, target) 
    if found {
        t.Fatal("Word not in space!")
    }
}

func TestFind_found(t *testing.T) {
    searchSpace := []string{"ab","bc", "cd","de","fgh","fordle","lol"}
    target := "fordle"
    found := Find(searchSpace, target) 
    if !found {
        t.Fatal("Word in space!")
    } 
}

func TestGetWords_length(t *testing.T) {
    words := GetWords()
    if len(words) != 2309 {
	t.Fatal("Unsuccessful in getting all the words")
    }
}

func TestGetWords_isAlphabetical(t *testing.T) {
    words := GetWords()
    for i := 0; i < len(words) - 1; i++ {
	cur := words[i]
	next := words[i+1]
	if next <= cur {
	    t.Fatal("The words are jumbled")
	}
    }
}
