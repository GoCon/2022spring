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

type Slot struct {
	Key      string `csv:"key"`
	Start    string `csv:"start"`
	Duration int    `csv:"duration"`
	Row      Row
}

type Row struct {
	Start int `csv:"row.start"`
	End   int `csv:"row.end"`
}

const fileName = "slots"

func generate(c []*Slot) {
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

	var slots []*Slot
	if err := gocsv.UnmarshalFile(f, &slots); err != nil {
		panic(err)
	}
	generate(slots)
}
