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

type Category struct {
	Key  string `csv:"key"`
	Name string `csv:"name"`
	ID   string `csv:"id"`
}

const fileName = "categories"

func createCategory(c []*Category) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Join(wd, "data")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	f, err := os.OpenFile(filepath.Join(dirPath, fmt.Sprintf("%s.yml", fileName)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	out, err := yaml.Marshal(c)
	if err != nil {
		panic(err)
	}

	rd := strings.NewReader(string(out))

	_, err = io.Copy(f, rd)
	if err != nil {
		panic(err)
	}
}

func main() {
	f, err := os.Open(fmt.Sprintf("./raw_data/%s.csv", fileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var categories []*Category
	if err := gocsv.UnmarshalFile(f, &categories); err != nil {
		panic(err)
	}
	createCategory(categories)
}
