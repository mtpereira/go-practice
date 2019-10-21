package popular

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func popularWord(url string) (words []string, count int, err error) {
	r, err := http.Get(url)
	if err != nil {
		return
	}

	var data []byte
	var posts []post

	data, err = ioutil.ReadAll(r.Body)
	err = json.Unmarshal(data, &posts)
	if err != nil {
		return
	}

	wordCount := map[string]map[int]int{}

	for _, p := range posts {
		words := strings.Fields(p.Body)
		for _, w := range words {
			if _, ok := wordCount[w]; !ok {
				wordCount[w] = make(map[int]int)
			}
			wordCount[w][p.UserID]++
		}
	}

	wordsPerCount := make(map[int][]string)
	largest := 0
	for w, uc := range wordCount {
		count := 0
		for _, c := range uc {
			if c > 0 {
				count++
			}
		}
		wordsPerCount[count] = append(wordsPerCount[count], w)
		if count > largest {
			largest = count
		}
	}

	sort.Strings(wordsPerCount[largest])
	return wordsPerCount[largest], largest, nil
}
