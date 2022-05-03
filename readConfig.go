package main

import (
	"encoding/json"
	"io/ioutil"
)

//Post is the struct for each post of the blog
type Post struct {
	Thumb     string `json:"thumb"`
	Md        string `json:"md"`
	MetaImg   string `json:"metaimg"`
	MetaDescr string `json:"metadescr"`
}

//Config gets the config.json file into struct
type Config struct {
	Title             string   `json:"title"`
	MetaImg           string   `json:"metaimg"`
	MetaDescr         string   `json:"metadescr"`
	RelativePath      string   `json:"relativePath"`
	PostsDir          string   `json:"postsDir"`
	AbsoluteUrl       string   `json:"absoluteUrl"`
	IndexTemplate     string   `json:"indexTemplate"`
	PostThumbTemplate string   `json:"postThumbTemplate"`
	Posts             []Post   `json:"posts"`
	CopyRaw           []string `json:"copyRaw"`
	OutputDir         string   `json:"outputDir"`
}

var config Config

func readConfig(path string) {
	file, err := ioutil.ReadFile(path)
	check(err)
	content := string(file)
	json.Unmarshal([]byte(content), &config)
	if config.PostsDir != "" {
		config.PostsDir += "/"
	}
	if config.OutputDir == "" {
		config.OutputDir = defaultOutputDir
	}
}
