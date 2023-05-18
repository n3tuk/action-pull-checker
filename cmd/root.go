package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	logger   *logrus.Logger
	logLevel string

	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:               "pull-requester",
		Short:             "A GitHub Action application for checking pull requests",
		PersistentPreRunE: initConfig,
	}
)

// Initialize the command-line settings for the root command
func init() {
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", logrus.InfoLevel.String(), "Logging level (debug, info, warn, error, fatal, panic)")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func initConfig(_ *cobra.Command, _ []string) error {
	var err error

	if err = setupLogging(); err != nil {
		return err
	}

	return nil
}

// Create an instance of the logging service and configure it
func setupLogging() error {
	logger = logrus.New()

	executable, err := os.Executable()
	if err != nil {
		executable = "pull-requester"
	}

	if envRunnerDebug := os.Getenv("RUNNER_DEBUG"); envRunnerDebug == "1" {
		logger.SetLevel(logrus.DebugLevel)
	} else if level, err := logrus.ParseLevel(logLevel); err == nil {
		logger.SetLevel(level)
	} else {
		return fmt.Errorf("failed to process log level %s: %w", logLevel, err)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})

	logger.
		WithFields(logrus.Fields{
			"version": Version,
			"commit":  Commit,
		}).
		Infof("%s starting...", path.Base(executable))

	return nil
}
