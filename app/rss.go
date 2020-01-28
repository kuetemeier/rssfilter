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
	"io/ioutil"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
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
