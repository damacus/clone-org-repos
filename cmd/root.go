/*
Copyright © 2021 Dan Webb<dan.webb@damacus.io>

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
	"errors"
	"os"
	"path/filepath"

	"github.com/damacus/clone-org-repos/checkout"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "clone-org-repos",
	Short: "A tool to clone all repositories in a github org",
	Long:  `clone-org-repos allows you to clone all repositories and update them within a given file path`,
	RunE: func(cmd *cobra.Command, args []string) error {
		token := os.Getenv("GITHUB_TOKEN")
		if token == "" {
			return errors.New("GITHUB_TOKEN must be set")
		}

		org, err := getStringFlag("org", cmd)
		if err != nil {
			return err
		}

		path, err := getStringFlag("path", cmd)
		if err != nil {
			return err
		}

		home, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		if path == "" {
			path = home
		}
		if !filepath.IsAbs(path) {
			path = filepath.Join(home, path)
		}

		checkout.Checkout(token, org, path)
		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringP("path", "p", "", "Path to checkout repositories to, defaults to user's home directory")
	rootCmd.PersistentFlags().StringP("org", "o", "", "Name of the org you wish to checkout")
	cobra.CheckErr(rootCmd.MarkPersistentFlagRequired("org"))
}

func getStringFlag(flagName string, cmd *cobra.Command) (string, error) {
	name, err := cmd.Flags().GetString(flagName)
	if err != nil {
		return "", err
	}
	return name, nil
}
