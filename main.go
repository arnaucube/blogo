package main

import (
	"fmt"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

const directory = "blogo-input"

func main() {
	readConfig(directory + "/blogo.json")
	fmt.Println(config)

	// generate index page
	indexTemplate := readFile(directory + "/" + config.IndexTemplate)
	indexPostTemplate := readFile(directory + "/" + config.PostThumbTemplate)
	var blogoIndex string
	blogoIndex = ""
	for _, post := range config.Posts {
		mdpostthumb := readFile(directory + "/" + post.Thumb)
		htmlpostthumb := string(blackfriday.Run([]byte(mdpostthumb)))

		//put the htmlpostthumb in the blogo-index-post-template
		m := make(map[string]string)
		m["[blogo-index-post-template]"] = htmlpostthumb
		r := putHTMLToTemplate(indexPostTemplate, m)
		filename := strings.Split(post.Md, ".")[0]
		r = "<a href='" + filename + ".html'>" + r + "</a>"
		blogoIndex = blogoIndex + r
	}
	//put the blogoIndex in the index.html
	m := make(map[string]string)
	m["[blogo-title]"] = config.Title
	m["[blogo-content]"] = blogoIndex
	r := putHTMLToTemplate(indexTemplate, m)
	writeFile("index.html", r)

	// generate posts pages

	for _, post := range config.Posts {
		mdcontent := readFile(directory + "/" + post.Md)
		htmlcontent := string(blackfriday.Run([]byte(mdcontent)))

		m := make(map[string]string)
		m["[blogo-title]"] = config.Title
		m["[blogo-content]"] = htmlcontent

		r := putHTMLToTemplate(indexTemplate, m)
		//fmt.Println(r)

		filename := strings.Split(post.Md, ".")[0]
		writeFile(filename+".html", r)
	}

	//copy raw
	fmt.Println("copying raw:")
	for _, dir := range config.CopyRaw {
		copyRaw(directory+"/"+dir, ".")
	}
}

func putHTMLToTemplate(template string, m map[string]string) string {
	lines := getLines(template)
	var resultL []string
	for _, line := range lines {
		inserted := false
		for k, v := range m {
			if strings.Contains(line, k) {
				//in the line, change [tag] with the content
				lineReplaced := strings.Replace(line, k, v, -1)
				resultL = append(resultL, lineReplaced)
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
