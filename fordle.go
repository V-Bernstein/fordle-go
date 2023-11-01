package main

import (
    "fmt"
    "fordle-go/fordle/types"
    "fordle-go/fordle/utils"
    "github.com/gorilla/mux"
    "io"
    "math/rand"
    "net/http"
    "os"
    "time"
)

func main() {
    if (len(os.Args) < 3) {
	fmt.Println("Error, borked args")
	return
    }
    var s1 = os.Args[1]
    var s2 = os.Args[2]
    test := utils.IsAnagram(s1, s2)
    fmt.Printf("Ana? %v\n", test)

    r := mux.NewRouter()

    r.HandleFunc("/words/validate/{word}", isWordValid)
    r.HandleFunc("/words/start", getTargetWord)
    r.HandleFunc("/words/guess", makeGuess)

    http.Handle("/", r)
    listenErr := http.ListenAndServe(":8080", nil)
    if listenErr != nil {
	fmt.Printf("Listener encountered an error!")
	return
    }
}

func isWordValid(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	return
    }

    vars := mux.Vars(r)
    word, found := vars["word"]

    if !found {
	fmt.Printf("Required arg not passed")
	http.Error(w, http.StatusText(400), http.StatusBadRequest)
	return
    }

    words := utils.GetWords()
    valid := utils.Find(words, word)

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "%t", valid)
}

func getTargetWord(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet{
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	return
    }
    words := utils.GetWords()
    rand.Seed(time.Now().UnixNano())
    randomIdx := rand.Intn(len(words))

    w.WriteHeader(http.StatusOK)
    io.WriteString(w, words[randomIdx])
}

func makeGuess(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	return
    }

    target := r.URL.Query().Get("target")
    guess := r.URL.Query().Get("guess")
    fmt.Printf("%s, %s", target, guess)

    /* TODO */
    tgtLetterMap := make(map[byte][]int)
    for idx, letter := range target {
	ltrs, contains := tgtLetterMap[letter]
	if contains {
	    fmt.Printf("Appending")
	    append(ltrs, idx)
	} else {
	    fmt.Printf("Appending to nil slice")
	    append(ltrs, idx)
	}
    }

    /* 
for (i in target.indices) { // Fill the target letter map
            val letter = target[i]
            if (tgtLetterMap.containsKey(letter)) {
                val indices = tgtLetterMap[letter]!!
                indices.add(i) // Add index to map
            } else {
                val tempList: MutableList<Int> = mutableListOf(i)
                tgtLetterMap[letter] = tempList
            }
        }
    */
    


    w.WriteHeader(http.StatusOK)
    io.WriteString(w, "nope") 
}

