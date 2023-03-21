package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func main() {
	url := "https://twitter.com/devabanoub"
	urls := GetImageLinksFromUrl(url)
	for i := 0; i < len(urls); i++ {
		err := DownloadImagesFromUrl("D:/projects/side/burnout/getimg/main.go", fmt.Sprintf("%s%s", url, urls[i]))
		if err != nil {
			panic(err)
		}
	}
}

func GetImageLinksFromUrl(url string) []string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	regex, _ := regexp.Compile("[a-zA-Z0-9/_.:-]+.(jpg|png)")
	links := regex.FindAllString(string(body), -1)
	return links
}

func DownloadImagesFromUrl(path string, url string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	out, err := os.Create(path + filepath.Base(url))
	if err != nil {
		return err
	}

	defer out.Close()

	_, err = io.Copy(out, response.Body)
	return err
}
