package cmd

import (
	"log"
	"orgf/cmd/images"
	"orgf/cmd/documents"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "orgf <command>",
	Short: "Folder organizer",
}

func init() {
	rootCmd.AddCommand(images.ImgCommand, documents.DocCommand)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		return
	}
}