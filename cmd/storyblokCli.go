/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/apjames93/mui-storyblok-cli/api/muiclient"
	"github.com/apjames93/mui-storyblok-cli/pkg/config"
	"github.com/spf13/cobra"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// storyblokCliCmd represents the storyblokCli command
var storyblokCliCmd = &cobra.Command{
	Use:   "storyblokCli",
	Short: "sets storyBlokSpaceID and storyBlokOathToken in config",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(themeURL, authToken, spaceID)

		if themeURL == "" ||
			authToken == "" ||
			spaceID == "" ||
			accessToken == "" {
			fmt.Println("you need to have all fags added")
			os.Exit(1)
		}

		c := config.New(spaceID, authToken, themeURL, accessToken)

		if err := muiclient.GitRepo(c); err != nil {
			fmt.Printf("git error %s", err.Error())
			os.Exit(1)
		}

		// we aeclonerp

	},
}

// %+v
func init() {
	rootCmd.AddCommand(storyblokCliCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// storyblokCliCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// storyblokCliCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
