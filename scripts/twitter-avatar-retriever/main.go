package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/gocarina/gocsv"
)

type Speaker struct {
	Twitter string `csv:"twitter" yaml:"-"`
}

var (
	consumerKey    = os.Getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
	token          = os.Getenv("TWITTER_TOKEN")
	tokenSecret    = os.Getenv("TWITTER_TOKEN_SECRET")
)

const imageDirPath = "../../static/images/speakers"

func downloadProfileImage(screenName, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	f, err := os.OpenFile(filepath.Join(imageDirPath, fmt.Sprintf("%s.png", screenName)), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := os.MkdirAll(imageDirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// すでにプロフィール画像が存在する screenName を記録する
	files, err := os.ReadDir(imageDirPath)
	if err != nil {
		panic(err)
	}
	alreadyExistsScreenNames := map[string]bool{}
	for _, f := range files {
		baseName := f.Name()
		screenName := strings.TrimSuffix(baseName, filepath.Ext(baseName))
		alreadyExistsScreenNames[screenName] = true
	}

	// CSV から screenName の一覧を取得する
	f, err := os.Open("../../raw_data/speakers.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var speakers []*Speaker
	if err := gocsv.UnmarshalFile(f, &speakers); err != nil {
		panic(err)
	}
	screenNames := make([]string, 0, len(speakers))
	for _, speaker := range speakers {
		screenName := speaker.Twitter
		// screenName が設定されていないか、既にプロフィール画像が存在していればスキップ
		if screenName == "" || alreadyExistsScreenNames[screenName] {
			continue
		}
		screenNames = append(screenNames, screenName)
	}

	// Twitter からアバター画像を取得する
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	oauthToken := oauth1.NewToken(token, tokenSecret)

	httpClient := config.Client(oauth1.NoContext, oauthToken)
	client := twitter.NewClient(httpClient)

	users, _, err := client.Users.Lookup(&twitter.UserLookupParams{
		ScreenName: screenNames,
	})
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		// retrieve original size image. see: https://developer.twitter.com/en/docs/twitter-api/v1/accounts-and-users/user-profile-images-and-banners
		imageURL := strings.Replace(user.ProfileImageURL, "_normal", "", 1)
		fmt.Println(imageURL)
		err := downloadProfileImage(user.ScreenName, imageURL)
		if err != nil {
			panic(err)
		}
	}
}
