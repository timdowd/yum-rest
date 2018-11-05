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

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// RootCmd represents the base command when called without any subcommands
	RootCmd = &cobra.Command{
		Use:   "thing",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application. For example:
	
	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
	}
)

func init() {

	RootCmd.PersistentFlags().StringP("server-ip", "s", "0.0.0.0", "IP the server will listen on")
	RootCmd.PersistentFlags().StringP("servicename", "S", "", "Name of the service that is running")
	RootCmd.PersistentFlags().StringP("rest-port", "R", "", "Port the rest server will listen on")
	RootCmd.PersistentFlags().StringP("rpc-port", "r", "", "Port the rpc server will listen on")
	RootCmd.PersistentFlags().StringP("rpc-address", "a", "", "Address the rpc server will listen on")
	RootCmd.PersistentFlags().StringP("traceserviceaccountfile", "f", "", "Google Service Account file path")
	RootCmd.PersistentFlags().StringP("projectid", "P", "", "Google cloud project id, e.g. -P phdigidev")

	_ = viper.BindPFlag("server-ip", RootCmd.PersistentFlags().Lookup("server-ip"))
	_ = viper.BindPFlag("servicename", RootCmd.PersistentFlags().Lookup("servicename"))
	_ = viper.BindPFlag("rest-port", RootCmd.PersistentFlags().Lookup("rest-port"))
	_ = viper.BindPFlag("rpc-port", RootCmd.PersistentFlags().Lookup("rpc-port"))
	_ = viper.BindPFlag("rpc-address", RootCmd.PersistentFlags().Lookup("rpc-address"))
	_ = viper.BindPFlag("traceserviceaccountfile", RootCmd.PersistentFlags().Lookup("traceserviceaccountfile"))
	_ = viper.BindPFlag("projectid", RootCmd.PersistentFlags().Lookup("projectid"))

}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Printf("Start")
		os.Exit(1)
	}
}
