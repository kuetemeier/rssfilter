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
package cmd

import (
	"time"

	feeds "github.com/gorilla/feeds"
	"github.com/kuetemeier/rssfilter/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// demodataCmd represents the demodata command
var demodataCmd = &cobra.Command{
	Use:   "demodata",
	Short: "Create a RSS feed from demo data.",
	Long: `Create a DEMO RSS feed. Do what ever you like with it.
	
	You can even pipe it into another rssfilter instance.`,
	Run: func(cmd *cobra.Command, args []string) {

		setNr := 1

		if viper.GetBool("demo2") {
			setNr = 2
		}

		now := time.Now()

		if setNr == 1 {
			outputFeed = &feeds.Feed{
				Title:       "rssfilter Demo Feed - No. 1",
				Link:        &feeds.Link{Href: "https://github.com/kuetemeier/rssfilter"},
				Description: "This is the description for the demo feed. (First Set)",
				Author:      &feeds.Author{Name: "Jörg Kütemeier", Email: "joerg@kuetemeier.de"},
				Created:     now,
			}

			outputFeed.Items = []*feeds.Item{
				&feeds.Item{
					Title:       "Demo 1 - Feed Item One",
					Link:        &feeds.Link{Href: "https://demo/blog/item-one/"},
					Description: "A wonderfull demo feed item.",
					Author:      &feeds.Author{Name: "John Müller", Email: "jmueller@demo.de"},
					Created:     now,
				},
				&feeds.Item{
					Title:       "Demo 1 - Feed Item Two",
					Link:        &feeds.Link{Href: "https://demo/blog/item-two/"},
					Description: "More thoughts on demo feed items.",
					Created:     now,
				},
				&feeds.Item{
					Title:       "Demo 1 - Feed Item Three",
					Link:        &feeds.Link{Href: "https://demo/blog/item-three/"},
					Description: "And another <em>great</em> demo feed item.",
					Created:     now,
				},
				&feeds.Item{
					Title:       "Double Item 1 in Demo 1 and 2",
					Link:        &feeds.Link{Href: "https://demo/blog/item-double-1/"},
					Description: "This item has the same title and link in Demo 1 and 2.",
					Created:     now,
				},
				&feeds.Item{
					Title:       "Double Item 2 in Demo 1 and 2 - First version",
					Link:        &feeds.Link{Href: "https://demo/blog/item-double-2/"},
					Description: "This item has the same link, but different title in Demo 1 and 2.",
					Created:     now,
				},
			}

		} else {
			outputFeed = &feeds.Feed{
				Title:       "rssfilter Demo Feed - No. 2",
				Link:        &feeds.Link{Href: "https://github.com/kuetemeier/rssfilter"},
				Description: "This is the description for the demo feed. (Second Set)",
				Author:      &feeds.Author{Name: "Jörg Kütemeier", Email: "joerg@kuetemeier.de"},
				Created:     now,
			}

			outputFeed.Items = []*feeds.Item{
				&feeds.Item{
					Title:       "Demo 2 - Feed Item One",
					Link:        &feeds.Link{Href: "https://demo/blog/item-a/"},
					Description: "Call this a wonderfull demo feed item.",
					Author:      &feeds.Author{Name: "John Müller", Email: "jmueller@demo.de"},
					Created:     now,
				},
				&feeds.Item{
					Title:       "Demo 2 - Feed Item Two",
					Link:        &feeds.Link{Href: "https://demo/blog/item-b/"},
					Description: "And even more thoughts on demo feed items.",
					Created:     now,
				},
				&feeds.Item{
					Title:       "Demo 2 - Feed Item Three",
					Link:        &feeds.Link{Href: "https://demo/blog/item-c/"},
					Description: "Can you believe it? Another <em>great</em> demo feed item.",
					Created:     now,
				},
				&feeds.Item{
					Title:       "Double Item 1 in Demo 1 and 2",
					Link:        &feeds.Link{Href: "https://demo/blog/item-double-1/"},
					Description: "This item has the same title and link in Demo 1 and 2.",
					Created:     now,
				},
				&feeds.Item{
					Title:       "Double Item 2 in Demo 1 and 2 - Second version",
					Link:        &feeds.Link{Href: "https://demo/blog/item-double-2/"},
					Description: "This item has the same link, but different title in Demo 1 and 2.",
					Created:     now,
				},
			}

		}

		app.ExportRSSFeed(outputFeed, viper.GetString("output"), viper.GetString("outputFormat"))

	},
}

func init() {
	rootCmd.AddCommand(demodataCmd)

	demodataCmd.Flags().Bool("demo1", true, "First set of demo data (default)")
	demodataCmd.Flags().Bool("demo2", false, "Second set of demo data")

	viper.BindPFlag("demo1", demodataCmd.Flags().Lookup("demo1"))
	viper.BindPFlag("demo2", demodataCmd.Flags().Lookup("demo2"))

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// demodataCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// demodataCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
