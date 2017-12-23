package main

import (
	"encoding/json"
	"io/ioutil"
)

type Post struct {
	Title string `json:"title"`
	Thumb string `json:"thumb"`
	Md    string `json:"md"`
}
type Config struct {
	Title             string `json:"title"`
	IndexTemplate     string `json:"indexTemplate"`
	IndexPostTemplate string `json:"indexPostTemplate"`
	PostTemplate      string `json:"postTemplate"`
	Posts             []Post `json:"posts"`
	CopyRaw           []string `json:"copyRaw"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
}
