package popular

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type post struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func popularWord(url string) string {
	r, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	var data []byte
	var posts []post

	data, err = ioutil.ReadAll(r.Body)
	err = json.Unmarshal(data, &posts)
	if err != nil {
		panic(err)
	}

	wordCount := make(map[string][]int, 1)

	for _, p := range posts {
		words := strings.Fields(p.Body)
		for _, w := range words {
			userIDs := wordCount[w]
			present := false
			for _, u := range userIDs {
				if p.UserID == u {
					present = true
				}
			}
			if present == false {
				wordCount[w] = append(wordCount[w], p.UserID)
			}
		}
	}

	var mostPopular string
	largest := 0
	for w, u := range wordCount {
		if len(u) > largest {
			largest = len(u)
			mostPopular = w
		}
	}

	return mostPopular
}
