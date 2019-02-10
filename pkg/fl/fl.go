// Package fl – Find Links. Tools for searching downloadable links on the web page
//
// @egregors 2019
package fl

import (
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

var defaultExts = []string{"webm", "mp4"}

func unique(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// Finder for finding
type Finder struct {
	url     string
	exts    []string
	links   []string
	linksCh chan string
}

func (f *Finder) getDoc() (*html.Node, error) {
	r, err := http.Get(f.url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	doc, err := html.Parse(r.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (f *Finder) findAll(n *html.Node) ([]string, error) {
	url := strings.Split(f.url, "/")
	domain := url[0] + "//" + url[1] + url[2]

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				for _, ext := range f.exts {
					if strings.Contains(a.Val, ext) {
						f.links = append(f.links, domain+a.Val)
					}
				}
			}
		}
	}

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		f.findAll(child)
	}
	return f.links, nil
}

func (f *Finder) find() error {
	doc, err := f.getDoc()
	if err != nil {
		return err
	}
	_, err = f.findAll(doc)
	if err != nil {
		return err
	}
	return nil
}

// GetLinks return string slice with links
func (f *Finder) GetLinks() ([]string, error) {
	if len(f.links) == 0 {
		err := f.find()
		if err != nil {
			return nil, err
		}
	}
	f.links = unique(f.links)

	return f.links, nil
}

// GetChan return chan of strings with links
func (f *Finder) GetChan() (chan string, error) {
	if f.linksCh == nil {
		f.linksCh = make(chan string)
		links, err := f.GetLinks()

		if err != nil {
			return nil, err
		}

		go func() {
			for _, link := range links {
				f.linksCh <- link
			}
			close(f.linksCh)
		}()
	}
	return f.linksCh, nil
}

// NewFinder create new Finder
func NewFinder(pageURL string, exts []string) Finder {
	if len(exts) == 0 {
		exts = defaultExts
	}
	return Finder{url: pageURL, exts: exts}
}
