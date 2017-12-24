# Blogo [![Go Report Card](https://goreportcard.com/badge/github.com/arnaucode/blogo)](https://goreportcard.com/report/github.com/arnaucode/blogo)
Static blog generator, templating engine from markdown and html templates


## Use
Directory structure:

```
/
----blogo
----/blogo-input
--------all the html, js, css files and folders
```

To execute:
```
./blogo
```

Example of blogo.json:

```json
{
    "title": "my blog",
    "indexTemplate": "index.html",
    "postThumbTemplate": "postThumbTemplate.html",
    "posts": [
        {
            "thumb": "post01_thumb.md",
            "md": "post01.md"
        },
        {
            "thumb": "post02_thumb.md",
            "md": "post02.md"
        }
    ],
    "copyRaw": [
      "css",
      "js"
    ]
}
```


Example of input files:
- index.html

```html
<!DOCTYPE html>
<html>
<head>
  <title>[blogo-title]</title>
</head>
<body>

[blogo-content]

</body>
</html>
```

- postThumbTemplate.html

```html
<div class="col-md-3">
  [blogo-index-post-template]
</div>

```

- post01_thumb.md

```
# Post 01 thumb
This is the description of the Post 01. This will appear on the main page, as the post description.
```

- post01.md

```
# Post 01
This is the content of the Post 01. This content will appear when the Post 01 from the main page is clicked.
```


Types of blogo tags:

```
[blogo-title]
[blogo-content]
[blogo-index-post-template]
```
