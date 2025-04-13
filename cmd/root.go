/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/rstms/rspamd-rescan/rescan"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const Version = "0.0.4"

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rspamd-rescan FILE",
	Short: "scan an email message with rspamd and update its headers",
	Long: `
Read FILE and send it to a remote rspamd server for scanning.  Read the JSON
data returned by rspamd and modify the message headers with the result.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
	    err := rescan.Rescan(args[0]);
	    cobra.CheckErr(err);
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.config/rspamd-rescan/rspamd-rescan.yaml)")

	rootCmd.PersistentFlags().StringP("log-file", "l", "/var/log/rspamd-rescan.log", "log filename")
	viper.BindPFlag("log_file", rootCmd.PersistentFlags().Lookup("log-file"))

	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable diagnostic output")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	rootCmd.PersistentFlags().String("cert", filepath.Join(home, "/ssl/filterctl.pem"), "client certificate PEM file")
	viper.BindPFlag("cert", rootCmd.PersistentFlags().Lookup("cert"))

	rootCmd.PersistentFlags().String("key", filepath.Join(home, "/ssl/filterctl.key"), "client certificate key file")
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))

	rootCmd.PersistentFlags().String("ca", "/etc/ssl/keymaster.pem", "certificate authority file")
	viper.BindPFlag("ca", rootCmd.PersistentFlags().Lookup("ca"))

	rootCmd.PersistentFlags().String("url", "http://localhost:3000", "server url")
	viper.BindPFlag("url", rootCmd.PersistentFlags().Lookup("url"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		configDir, err := os.UserConfigDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".rspamd-rescan" (without extension).
		viper.AddConfigPath(filepath.Join(configDir, "rspamd-rescan"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("rspamd-rescan")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
