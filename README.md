# MyBlogInGo
Rewrite my blog generator in Go

## Components
- [blackfriday V2](https://github.com/russross/blackfriday/tree/v2) is a Markdown processor implemented in Go.
- [bfchroma](https://github.com/Depado/bfchroma/), Integrating Chroma syntax highlighter as a Blackfriday renderer.

## FAQ
Q1. Checked the header format and content is right, but still got "Header not exist" error. Why?

Check whether the line break is "CRLF".