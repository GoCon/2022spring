package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v2"
)

const GitHubURLPrefix = "https://github.com/"

type Speaker struct {
	Key      string `csv:"key"`
	ID       string `csv:"id"`
	Name     string `csv:"name"`
	Company  string `csv:"company"`
	Twitter  string `csv:"twitter" yaml:"-"`
	Site     string `csv:"site" yaml:"-"`
	PhotoURL string `yaml:"photoURL"`
	Bio      string `csv:"bio" yaml:"-"`
	Socials  []Social
}

type Social struct {
	Icon string
	Link string
	Name string
}

const imageDirPath = "static/images/speakers"

func createSpeaker(ss []*Speaker) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	files, err := os.ReadDir(imageDirPath)
	if err != nil {
		panic(err)
	}

	photoURLMap := make(map[string]string, len(ss))
	for _, f := range files {
		baseName := f.Name()
		key := strings.TrimSuffix(baseName, filepath.Ext(baseName))
		if photoURLMap[key] != "" {
			log.Fatalf("multiple files found for key: %s", key)
		}
		photoURLMap[key] = baseName
	}

	dirPath := filepath.Join(wd, "content/speakers")
	for _, s := range ss {
		if s.Twitter != "" {
			s.Socials = append(s.Socials, Social{
				Icon: "twitter",
				Link: fmt.Sprintf("https://twitter.com/%s", s.Twitter),
				Name: s.Twitter,
			})
		}
		if s.Site != "" {
			s.Socials = append(s.Socials, Social{
				Icon: "link",
				Link: s.Site,
				Name: s.Site,
			})
		}

		s.PhotoURL = fmt.Sprintf("/images/speakers/%s", photoURLMap[s.Key])
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		f, err := os.OpenFile(filepath.Join(dirPath, fmt.Sprintf("%s.md", s.Key)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		out, err := yaml.Marshal(s)
		if err != nil {
			panic(err)
		}
		body := fmt.Sprintf(`---
%s---
%s`, string(out), s.Bio)
		rd := strings.NewReader(body)

		_, err = io.Copy(f, rd)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	f, err := os.Open("./raw_data/speakers.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var speakers []*Speaker
	if err := gocsv.UnmarshalFile(f, &speakers); err != nil { // Load clients from file
		panic(err)
	}
	createSpeaker(speakers)
}
