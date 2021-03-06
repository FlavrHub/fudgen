package main

import (
	"io/ioutil"

	"./recipes"
	//"./units"
	"encoding/json"
	"fmt"
	"math/rand"
	//"io/ioutil"
	"net/http"
)

type message struct {
	Error    string         `json:"error"`
	Recipe   recipes.Recipe `json:"recipe"`
	Schedule [][][]int      `json:"sched"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, "Hello Web")
	})
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		m := getRandomRecipe()
		bytes, err := json.Marshal(m)
		if err != nil {
			bytes = []byte("{error:'I dont even know'}")
		}
		w.Write(bytes)
	})
	http.ListenAndServe(":8890", nil)
}

func getRandomRecipe() (m message) {
	fileNumber := rand.Intn(3)
	b, err := ioutil.ReadFile("../outlines/" + fmt.Sprintf("%d", fileNumber) + ".yml")
	if err != nil {
		panic(err)
	}
	r, err := recipes.ParseYaml(string(b))

	//r, err := recipes.RandomRecipe(recipes.RandomParameters{})
	m.Schedule, err = recipes.Schedule(r)
	if err != nil {
		m.Error = err.Error()
	}
	m.Recipe = *r
	return
}
