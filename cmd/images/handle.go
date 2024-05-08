package images

import (
	"log"
	"orgf/internal"
)

var (
	defaultSuffixes = []string{".jpg", ".jpeg", ".png", ".svg", ".gif"}
	additionalSuffixes []string
	sourceFolder string
	folderName string
)

func HandleImages(sourceFolder string) {
	wd, err := internal.GetWorkingDir(sourceFolder)
	if err != nil {
		log.Fatal(err)
	}
	
	imgFolder, err := internal.CreateDestinationFolder(wd, folderName)
	if err != nil {
		log.Fatal(err)
	}

	defaultSuffixes = append(defaultSuffixes, additionalSuffixes...)
	if err = internal.HandleFileMove(wd, imgFolder, defaultSuffixes); err != nil {
		log.Fatal(err)
	}
}
