package documents

import (
	"log"
	"orgf/internal"
)

var (
	defaultSuffixes = []string{".doc",".docx",".pdf", ".txt"}
	additionalSuffixes []string
	sourceFolder string
	folderName string
)

func HandleDocuments(sourceFolder string) {
	wd, err := internal.GetWorkingDir(sourceFolder)
	if err != nil {
		log.Fatal(err)
	}

	docsFolder, err := internal.CreateDestinationFolder(wd, folderName)
	if err != nil {
		log.Fatal(err)
	}

	defaultSuffixes = append(defaultSuffixes, additionalSuffixes...)
	if err := internal.HandleFileMove(wd, docsFolder, defaultSuffixes); err != nil {
		log.Fatal(err)
	}
}