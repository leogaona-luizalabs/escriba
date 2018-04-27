// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/luizalabs/escriba/bot"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "escriba",
	Short: "Chatbot que organiza artigos",
	Long:  "Chatbot que organiza artigos",
	Run: func(cmd *cobra.Command, args []string) {
		slackToken := viper.GetString("slack_token")
		approvals := viper.GetInt("draft_approvals")
		mysqlDSN := viper.GetString("mysql_dsn")
		bot.Start(slackToken, mysqlDSN, approvals)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().String("slack_token", "", "slack_api_token")
	rootCmd.Flags().String("mysql_dsn", "", "mysql data source name")
	rootCmd.Flags().Int("draft_approvals", 2, "number of approvals for a draft to be eligible for publication")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.AutomaticEnv() // read in environment variables that match
}
