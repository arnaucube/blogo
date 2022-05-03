package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

const version = "v0_20220503"
const directory = "blogo-input"
const defaultOutputDir = "public"

func main() {
	devMode := flag.Bool("d", false, "dev mode")
	port := flag.String("p", "8080", "port (only for dev mode)")
	flag.Parse()

	fmt.Println("Blogo version:", version)
	readConfig(directory + "/blogo.json")
	_ = os.Mkdir(config.OutputDir, os.ModePerm)

	fmt.Println("devMode:", *devMode)
	generateHTML()
	if !*devMode {
		return
	}

	go watch("./blogo-input")
	// TODO public not watched, until a way to force browser refresh
	// go watch("./public")

	// serve files
	fs := http.FileServer(http.Dir(config.OutputDir))
	fmt.Printf("Blog being served in: \n http://127.0.0.1:%s\n http://localhost:%s\n",
		*port, *port)
	log.Fatal(http.ListenAndServe(":"+*port, fs))
}

func generateHTML() {
	readConfig(directory + "/blogo.json")
	fmt.Println(config)

	mdExtensions := parser.NoIntraEmphasis | parser.Tables | parser.FencedCode |
		parser.Autolink | parser.Strikethrough | parser.SpaceHeadings | parser.HeadingIDs |
		parser.BackslashLineBreak | parser.DefinitionLists

	// generate index page
	indexTemplate := readFile(directory + "/" + config.IndexTemplate)
	indexPostTemplate := readFile(directory + "/" + config.PostThumbTemplate)
	var blogoIndex string
	blogoIndex = ""
	for _, post := range config.Posts {
		mdpostthumb := readFile(directory + "/" + config.PostsDir + post.Thumb)
		mdParser := parser.NewWithExtensions(mdExtensions)
		htmlpostthumb := markdown.ToHTML([]byte(mdpostthumb), mdParser, nil)

		//put the htmlpostthumb in the blogo-index-post-template
		m := make(map[string]string)
		m["[blogo-index-post-template]"] = string(htmlpostthumb)
		r := putHTMLToTemplate(indexPostTemplate, m)
		filename := strings.Split(post.Md, ".")[0]
		r = "<a href='" + config.RelativePath + "/" + filename + ".html'>" + r + "</a>"
		blogoIndex = blogoIndex + r
	}
	//put the blogoIndex in the index.html
	m := make(map[string]string)
	m["[blogo-title]"] = config.Title
	m["[blogo-content]"] = blogoIndex
	m["[blogo-summary]"] = config.MetaDescr
	m["[blogo-img]"] = config.AbsoluteUrl + "/" + config.MetaImg
	m["[blogo-link]"] = config.AbsoluteUrl
	r := putHTMLToTemplate(indexTemplate, m)
	writeFile(config.OutputDir+"/"+"index.html", r)

	// generate posts pages

	for _, post := range config.Posts {
		mdcontent := readFile(directory + "/" + config.PostsDir + post.Md)
		mdParser := parser.NewWithExtensions(mdExtensions)
		htmlcontent := markdown.ToHTML([]byte(mdcontent), mdParser, nil)

		firstline := strings.Split(mdcontent, "\n")[0]
		title := strings.Replace(firstline, "#", "", -1)

		filename := strings.Split(post.Md, ".")[0]

		m := make(map[string]string)
		m["[blogo-title]"] = title + " - " + config.Title
		m["[blogo-content]"] = string(htmlcontent)
		m["[blogo-summary]"] = post.MetaDescr
		m["[blogo-link]"] = config.AbsoluteUrl + "/" + filename + ".html"
		m["[blogo-img]"] = config.AbsoluteUrl + "/" + post.MetaImg

		r := putHTMLToTemplate(indexTemplate, m)
		writeFile(config.OutputDir+"/"+filename+".html", r)
	}

	//copy raw
	fmt.Println("copying raw:")
	for _, dir := range config.CopyRaw {
		copyRaw(directory+"/"+dir, config.OutputDir+"/")
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
