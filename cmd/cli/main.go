package main

import (
	"log"
	"os"

	"github.com/Egregors/fl/pkg/fl"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("[FATAL] You should pass the valid URL of web page")
		os.Exit(1)
	}

	url := os.Args[1]
	extOfFilesILookingFor := []string{"webm", "mp4"}
	f := fl.NewFinder(url, extOfFilesILookingFor)

	links, err := f.GetLinks()
	if err != nil {
		panic(err)
	}

	if len(links) == 0 {
		log.Print("[INFO] Nothing to show :(")
	} else {

		for _, link := range links {
			log.Print("[INFO]", link)
		}
	}
}
