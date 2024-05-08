package documents

import "github.com/spf13/cobra"

var DocCommand = &cobra.Command{
	Use: "doc",
	Aliases: []string{"docs", "document", "documents"},
	Short: "Organize document files",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		HandleDocuments(sourceFolder)
	},
}

func init() {
	DocCommand.Flags().StringVarP(&folderName, "destination-folder", "d", "docs", "Destination folder name")
	DocCommand.Flags().StringVarP(&sourceFolder, "source-folder", "s", ".", "Source folder that orgf will run. If not specified, orgf will be run on current directory")
	DocCommand.Flags().StringArrayVarP(&additionalSuffixes, "suffix", "x", []string{}, "Additional files to organize together. Default files are the files have .doc, .docx, .pdf and .txt extensions")
}