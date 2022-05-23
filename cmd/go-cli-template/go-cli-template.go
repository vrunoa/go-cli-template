package main

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/vrunoa/go-cli-template/internal/greeter"
	"github.com/vrunoa/go-cli-template/internal/version"
	"os"
)

func setupLogging(verbose bool) {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "15:04:05"})
}

var lang string

func helloCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hello",
		Short: "say hello",
		Long:  "say hello in some language",
		Run: func(cmd *cobra.Command, args []string) {
			g := greeter.New(lang)
			g.Hello()
		},
	}
	cmd.Flags().StringVarP(&lang, "lang", "l", "", "--lang es")
	return cmd
}

func main() {
	setupLogging(true)
	mainCmd := &cobra.Command{
		Use:     "go-cli-template [command]",
		Short:   "A CLI written in Go",
		Version: fmt.Sprintf("%s\n(build %s)", version.Version, version.GitCommit),
	}
	mainCmd.AddCommand(helloCommand())
	if err := mainCmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("wops! seems like we messed up")
		os.Exit(1)
	}
}
