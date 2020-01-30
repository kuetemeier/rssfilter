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
	"fmt"

	"github.com/kuetemeier/rssfilter/app"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// filterCmd represents the filter command
var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "(default) filter rss streams",
	Long: `This is the default command.
	
	It can filter and manipulate an rss stream in many ways
	`,
	Run: run,
}

func init() {
	rootCmd.AddCommand(filterCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// filterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// filterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func run(cmd *cobra.Command, args []string) {
	if viper.GetBool("verbose") {
		fmt.Println("INFO: Applying filter")
		fmt.Println("  Input        : " + viper.GetString("input"))
		fmt.Println("  Output       : " + viper.GetString("output"))
		fmt.Println("  OutputFormat : " + viper.GetString("outputFormat"))
	}

	feed, err := app.ImportRSSFeed(viper.GetString("input"))

	if err != nil {
		panic(err)
	}

	out, err := app.FilterFeed(feed)

	if err != nil {
		panic(err)
	}

	err = app.ExportRSSFeed(out, viper.GetString("output"), viper.GetString("outputFormat"))

	if err != nil {
		panic(err)
	}
}
