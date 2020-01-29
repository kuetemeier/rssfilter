/*
Copyright © 2020 Jörg Kütemeier <joerg@kuetemeier.de>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	feeds "github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
	"github.com/spf13/viper"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ImportRSSFeed imports RSS feeds from different sources.
func ImportRSSFeed(source string) (*gofeed.Feed, error) {
	if (source == "STDIN") || (source == "") {
		return importRSSFeedFromStdin()
	}
	if strings.HasPrefix(source, "http") {
		return importRSSFeedFromURL(source)
	}
	return importRSSFeedFromFile(source)
}

func importRSSFeedFromURL(feedURL string) (*gofeed.Feed, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(feedURL)

	return feed, err
}

func importRSSFeedFromFile(filename string) (*gofeed.Feed, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fp := gofeed.NewParser()
	feed, err := fp.Parse(file)

	return feed, err
}

func importRSSFeedFromStdin() (*gofeed.Feed, error) {
	var data []byte

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseString(string(data))

	return feed, err
}

func info(msg string) {
	if viper.GetBool("verbose") {
		fmt.Println("INFO: " + msg)
	}
}

// ExportRSSFeed exports the given feed as string
func ExportRSSFeed(feed *feeds.Feed, destination string, format string) error {
	f := strings.ToLower(format)
	var res string
	var err error

	info("Output Format: " + format)
	info("Output Destination: " + destination)

	if f == "rss" {
		res, err = feed.ToRss()
	} else if f == "atom" {
		res, err = feed.ToAtom()
	} else if f == "json" {
		res, err = feed.ToJSON()

	} else {
		return errors.New("Unknow RSS Feed export format '" + format + "'")
	}

	fmt.Println(res)

	return err
}
