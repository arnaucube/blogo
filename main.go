package main

import (
	"fmt"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const directory = "input"

func main() {
	readConfig("input/config.json")
	fmt.Println(config)

	// generate index page
	indexTemplate := readFile(directory + "/" + config.IndexTemplate)
	indexPostTemplate := readFile(directory + "/" + config.IndexPostTemplate)
	var blogoIndex string
	blogoIndex = ""
	for _, post := range config.Posts {
		mdpostthumb := readFile(directory + "/" + post.Thumb)
		htmlpostthumb := string(blackfriday.Run([]byte(mdpostthumb)))

		//put the htmlpostthumb in the blogo-index-post-template
		m := make(map[string]string)
		m["blogo-index-post-template"] = htmlpostthumb
		r := putHTMLToTemplate(indexPostTemplate, m)
		blogoIndex = blogoIndex + r
	}
	//put the blogoIndex in the index.html
	m := make(map[string]string)
	m["blogo-title"] = config.Title
	m["blogo-index"] = blogoIndex
	r := putHTMLToTemplate(indexTemplate, m)
	writeFile("index.html", r)

	// generate posts pages
	postTemplate := readFile(directory + "/" + config.PostTemplate)

	for _, post := range config.Posts {
		mdcontent := readFile(directory + "/" + post.Md)
		htmlcontent := string(blackfriday.Run([]byte(mdcontent)))

		m := make(map[string]string)
		m["blogo-post-title"] = post.Title
		m["blogo-post-content"] = htmlcontent

		r := putHTMLToTemplate(postTemplate, m)
		//fmt.Println(r)

		filename := strings.Split(post.Md, ".")[0]
		writeFile(filename+".html", r)
	}

	//copy raw
	fmt.Println("copying raw:")
	for _, dir := range config.CopyRaw {
		copyRaw(directory+"/"+dir, dir)
	}
}

func putHTMLToTemplate(template string, m map[string]string) string {
	lines := getLines(template)
	var resultL []string
	for _, line := range lines {
		inserted := false
		for k, v := range m {
			if strings.Contains(line, k) {
				resultL = append(resultL, v)
				inserted = true
			}
		}
		if inserted == false {
			resultL = append(resultL, line)
		}
	}
	result := concatStringsWithJumps(resultL)
	return result
}
