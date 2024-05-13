/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

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
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/bristolgolang/rapid-go/internal/server"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rapid-go",
	Short: "A small microservice for demonstrating rapid development",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		connStr := viper.GetString("postgres_connection_string")
		slog.Info("connecting to database", slog.String("connection_string", connStr))
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			slog.Error("failed to connect to database", slog.String("error", err.Error()))
			return err
		}
		defer db.Close()

		err = db.PingContext(cmd.Context())
		if err != nil {
			slog.Error("failed to ping database", slog.String("error", err.Error()))
			return err
		}

		s := server.NewServer(db)

		router := http.NewServeMux()
		router.Handle("GET /greet/{name}", loggingMiddleware(http.HandlerFunc(s.Greet)))
		router.HandleFunc("GET /ready", ready)

		server := &http.Server{
			Addr:    ":" + viper.GetString("port"),
			Handler: router,
		}

		slog.Info("starting server", slog.String("port", viper.GetString("port")))
		return server.ListenAndServe()
	},
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("recieved request", slog.String("path", r.URL.Path), slog.String("method", r.Method), slog.String("remote_addr", r.RemoteAddr))
		startTime := time.Now()
		next.ServeHTTP(w, r)
		slog.Info("request completed", slog.String("path", r.URL.Path), slog.String("method", r.Method), slog.String("remote_addr", r.RemoteAddr), slog.Duration("duration", time.Since(startTime)))
	})
}

func ready(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rapid-go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.SetEnvPrefix("rapid_go")
	viper.SetDefault("port", "32400")
	viper.BindEnv("port")
	viper.BindEnv("postgres_connection_string")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".rapid-go" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".rapid-go")
	}

	viper.SetEnvPrefix("rapid_go") // set environment variable prefix
	viper.AutomaticEnv()           // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
