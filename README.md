fl – find links
===============

[![Go Report Card](https://goreportcard.com/badge/github.com/Egregors/fl)](https://goreportcard.com/report/github.com/Egregors/fl)

This is the simple tool to search downloadable links in the HTML document. It may be useful if you need to find all video | image | mp3 file links on the page. 

# Install
`go get github.com/Egreogrs/fl`

# Usage
Create a Finder
```
import "github.com/Egregors/fl/pkg/fl"

url := "http://my-url-to-search.com"
extOfFilesILookingFor := []string{"webm", "mp4"}
f := fl.NewFinder(url, extOfFilesILookingFor)
```

get links like slice of strings 
```
var links []string
links, err := f.GetLinks()
```

or channel of strings
```
var ch chan string
ch, err = f.GetChan()
```

# Testing
```
go test ./...
```