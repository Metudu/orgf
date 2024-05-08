package images

import (
	"github.com/spf13/cobra"
)

var ImgCommand = &cobra.Command{
	Use: "img",
	Aliases: []string{"imgs","image", "images"},
	Short: "Organize image files",
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		HandleImages(sourceFolder)
	},
}

func init() {
	ImgCommand.Flags().StringVarP(&folderName, "destination-folder", "d", "imgs", "Destination folder name")
	ImgCommand.Flags().StringVarP(&sourceFolder, "source-folder", "s", ".", "Source folder that orgf will run. If not specified, orgf will be run on current directory")
	ImgCommand.Flags().StringArrayVarP(&additionalSuffixes, "suffix", "x", []string{}, "Files will be organized based on suffixes")
}
