# Blogo
Static blog generator, templating engine from markdown and html templates

Types of blogo tags:
- index.html
```
[blogo-title]
[blogo-index]
```
- postTemplate.html
```
[blogo-post-title]
[blogo-post-md]
```


Example of config.json:
```json
{
    "title": "my blog",
    "indexTemplate": "index.html",
    "indexPostTemplate": "indexPostTemplate.html",
    "postTemplate": "postTemplate.html",
    "posts": [
        {
            "title": "Post 01",
            "thumb": "post01thumb.md",
            "md": "post01.md"
        },
        {
            "title": "Post 02",
            "thumb": "post02thumb.md",
            "md": "post02.md"
        }
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

[blogo-index]

</body>
</html>
```

- indexPostTemplate.html

```html
<div class="col-md-3">
  [blogo-index-post-template]
</div>

```

- postTemplate.html

```html
<!DOCTYPE html>
<html>
<head>
  <title>[blogo-post-title]</title>
</head>
<body>

[blogo-post-content]

</body>
</html>
```
