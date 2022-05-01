# Blogo [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucube/blogo)](https://goreportcard.com/report/github.com/arnaucube/blogo)
Static blog generator, templating engine from markdown and html templates

![blogo](https://raw.githubusercontent.com/arnaucube/blogo/master/blogo-diagram.png "blogo-diagram")

## Usage
```
Usage of blogo:
  -d dev mode
  -p port (only for dev mode) (default "8080")
```

So, for local usage:
```
blogo -d
```
Which will re-generate the html files each time that a input file is modified, and will serve the html generated site at `http://127.0.0.1:8080`.

For a single use, you can simply use `blogo` command, which will generate the html files without serving them.

A complete usage example can be found in this repo: https://github.com/arnaucube/blogoExample

### Config example
```json
{
    "title": "Blogo Blog",
    "relativePath": "",
    "absoluteUrl": "https://blog.website.com",
    "indexTemplate": "index.html",
    "postThumbTemplate": "postThumbTemplate.html",
    "posts": [
	{
	    "thumb": "article0_thumb.md",
	    "md": "article0.md",
	    "metaimg": "img/article0-img.png",
	    "metadescr": "description of the article 0"
	}
    ],
    "copyRaw": [
	"css",
	"img",
	"js"
    ]
}
```


---

Blogo is used in https://arnaucube.com/blog
