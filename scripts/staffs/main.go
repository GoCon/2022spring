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

type Staff struct {
	Key          string `csv:"key"`
	Name         string `csv:"name"`
	Organization string `csv:"organization"`
	Twitter      string `csv:"twitter" yaml:"-"`
	Github       string `csv:"github" yaml:"-"`
	Site         string `csv:"site" yaml:"-"`
	PhotoURL     string `yaml:"photoURL"`
	Socials      []Social
}

type Social struct {
	Icon string
	Link string
	Name string
}

const imageDirPath = "static/images/staffs"

func createStaff(ss []*Staff) {
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

	dirPath := filepath.Join(wd, "content/staffs")
	for _, s := range ss {
		s.Key = fmt.Sprintf("st_%s", s.Key) // key名にprefixを付ける

		if s.Twitter != "" {
			s.Socials = append(s.Socials, Social{
				Icon: "twitter",
				Link: fmt.Sprintf("https://twitter.com/%s", s.Twitter),
				Name: s.Twitter,
			})
		}
		if s.Github != "" {
			s.Socials = append(s.Socials, Social{
				Icon: "github",
				Link: fmt.Sprintf("https://github.com/%s", s.Github),
				Name: s.Github,
			})
		}
		if s.Site != "" {
			s.Socials = append(s.Socials, Social{
				Icon: "link",
				Link: s.Site,
				Name: s.Site,
			})
		}

		s.PhotoURL = fmt.Sprintf("/images/staffs/%s", photoURLMap[s.Key])
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
%s---`, string(out))
		rd := strings.NewReader(body)

		_, err = io.Copy(f, rd)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	f, err := os.Open("./raw_data/staffs.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var staffs []*Staff
	if err := gocsv.UnmarshalFile(f, &staffs); err != nil { // Load clients from file
		panic(err)
	}
	createStaff(staffs)
}
