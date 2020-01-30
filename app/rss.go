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

// Package app includes functions to import, export and filter rss.
package app

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

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

// FilterFeed creates a new, filtered feed from the input.
func FilterFeed(feed *gofeed.Feed) (*feeds.Feed, error) {
	now := time.Now()

	out := &feeds.Feed{
		Title: "rssfilter Feed",
		//Link:        &feeds.Link{Href: ""},
		//Description: "",
		//Author:      &feeds.Author{Name: "", Email: ""},
		//Created:     now,
	}

	if len(strings.TrimSpace(feed.Title)) != 0 {
		out.Title = strings.TrimSpace(feed.Title)
	}

	out.Link = &feeds.Link{}

	if len(strings.TrimSpace(feed.Link)) != 0 {
		out.Link.Href = strings.TrimSpace(feed.Link)
	}

	if len(strings.TrimSpace(feed.FeedLink)) != 0 {
		out.Link.Rel = strings.TrimSpace(feed.FeedLink)
	}

	if len(strings.TrimSpace(feed.Description)) != 0 {
		out.Description = strings.TrimSpace(feed.Description)
	}

	if feed.PublishedParsed != nil {
		out.Created = *feed.PublishedParsed
	}

	if feed.UpdatedParsed != nil {
		out.Updated = *feed.UpdatedParsed
	} else {
		out.Updated = now
	}

	if (feed.Image != nil) && (len(strings.TrimSpace(feed.Image.URL))) != 0 {
		i := &feeds.Image{Url: feed.Image.URL}
		if len(strings.TrimSpace(feed.Image.Title)) != 0 {
			i.Title = feed.Image.Title
		}
		out.Image = i
	}

	if len(strings.TrimSpace(feed.Copyright)) != 0 {
		out.Copyright = strings.TrimSpace(feed.Copyright)
	}

	if feed.Author != nil {
		a := &feeds.Author{}

		if len(strings.TrimSpace(feed.Author.Name)) != 0 {
			a.Name = strings.TrimSpace(feed.Author.Name)
		}

		if len(strings.TrimSpace(feed.Author.Email)) != 0 {
			a.Email = strings.TrimSpace(feed.Author.Email)
		}

		out.Author = a
	}

	count := viper.GetInt("count")

	for i, item := range feed.Items {
		new := &feeds.Item{}

		if count > -1 {
			if i >= count {
				break
			}
		}

		// Id
		if len(strings.TrimSpace(item.GUID)) != 0 {
			new.Id = strings.TrimSpace(item.GUID)
		}

		// Title
		if len(strings.TrimSpace(item.Title)) != 0 {
			new.Title = strings.TrimSpace(item.Title)
		}

		// Description
		if len(strings.TrimSpace(item.Description)) != 0 {
			new.Description = strings.TrimSpace(item.Description)
		}

		// Content
		if len(strings.TrimSpace(item.Content)) != 0 {
			new.Content = strings.TrimSpace(item.Content)
		}

		// Link
		link := &feeds.Link{}
		if len(strings.TrimSpace(item.Link)) != 0 {
			link.Href = strings.TrimSpace(item.Link)
		}
		new.Link = link

		// Publish date / time
		if item.PublishedParsed != nil {
			new.Created = *item.PublishedParsed
		}

		fmt.Println(item.Updated)
		// Update date / time
		if item.UpdatedParsed != nil {
			new.Updated = *item.UpdatedParsed
		}

		// Add the new Item to the output feed
		out.Items = append(out.Items, new)
	}

	return out, nil

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
