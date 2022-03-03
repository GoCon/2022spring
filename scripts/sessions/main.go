package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	"gopkg.in/yaml.v2"
)

type Session struct {
	Key          string `csv:"key"`
	Title        string `csv:"title"`
	ID           string `csv:"id"`
	Format       string `csv:"format"`
	TalkType     string `csv:"talkType" yaml:"talkType"`
	Level        string `csv:"level"`
	Tag          string `csv:"tag" yaml:"-"`
	Tags         []string
	Speaker      string `csv:"speaker" yaml:"-"`
	Speakers     []string
	Partner      *string `csv:"partner"`
	VideoID      *string `csv:"videoId" yaml:"videoId"`
	Presentation *string `csv:"presentation"`
	Draft        bool
	Abstract     string `csv:"abstract" yaml:"-"`
}

func generateSession(session *Session, dirPath string) {
	f, err := os.OpenFile(filepath.Join(dirPath, fmt.Sprintf("%s.md", session.Key)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if session.Speaker != "" {
		session.Speakers = append(session.Speakers, session.Speaker)
	}
	if session.Tag != "" {
		session.Tags = append(session.Tags, session.Tag)
	}
	if *session.VideoID == "" {
		session.VideoID = nil
	}
	if *session.Presentation == "" {
		session.Presentation = nil
	}
	if *session.Partner == "" {
		session.Partner = nil
	}

	out, err := yaml.Marshal(session)
	if err != nil {
		panic(err)
	}

	body := fmt.Sprintf(`---
%s---
%s`, string(out), session.Abstract)

	_, err = f.WriteString(body)
	if err != nil {
		panic(err)
	}
}

func generate(sessions []*Session) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dirPath := filepath.Join(wd, "content/sessions")
	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}
	for _, session := range sessions {
		generateSession(session, dirPath)
	}
}

func main() {
	f, err := os.Open("./raw_data/sessions.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var sessions []*Session
	if err := gocsv.UnmarshalFile(f, &sessions); err != nil {
		panic(err)
	}
	generate(sessions)
}
