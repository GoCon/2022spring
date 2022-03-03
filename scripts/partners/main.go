package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v2"
)

type Partner struct {
	Category    string `csv:"category"`
	Key         string `csv:"key"`
	Title       string `csv:"title"`
	Order       int    `csv:"order"`
	Description string `csv:"description"`
	Logo        string
}

func createPartner(ps []*Partner) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for _, p := range ps {
		dirPath := filepath.Join(wd, "content/partners", p.Category)
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
		p.Logo = fmt.Sprintf("/images/partners/%s.png", p.Key)
		f, err := os.OpenFile(filepath.Join(dirPath, fmt.Sprintf("%s.md", p.Key)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		out, err := yaml.Marshal(p)
		if err != nil {
			panic(err)
		}
		body := fmt.Sprintf(`---
%s---
%s`, string(out), p.Description)
		rd := strings.NewReader(body)

		_, err = io.Copy(f, rd)
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	f, err := os.Open("./raw_data/partners.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var partners []*Partner
	if err := gocsv.UnmarshalFile(f, &partners); err != nil {
		panic(err)
	}

	createPartner(partners)
}
