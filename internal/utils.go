package internal

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

func GetWorkingDir(sourceFolder string) (string, error) {
	wd, err := os.Getwd() // Get the current workign directory
	if err != nil {
		return "", err
	}

	if filepath.IsAbs(sourceFolder) { // Check if the path is absolute
		wd = sourceFolder
	} else {
		wd = filepath.Join(wd, sourceFolder) // if not absolute, concatenate the paths
	}

	return wd, nil
}

func CreateDestinationFolder(wd, folderName string) (string, error) {
	folderToMove := filepath.Join(wd, folderName)
	if err := os.Mkdir(folderToMove, 0777); err != nil {
		return "", err
	}

	return folderToMove, nil
}

func HandleFileMove(wd, folderToMove string, suffixes []string) error {
	files, err := os.ReadDir(wd)
	if err != nil {
		return err
	}

	for _, value := range files {
		for _, suffix := range suffixes {
			if strings.HasSuffix(value.Name(), suffix) {
				moveFile(wd, folderToMove, value.Name())
			}
		}
	}
	return nil
}

func moveFile(wd, folderToMove, fileName string) error {
	sourcePath := filepath.Join(wd, fileName)
	destinationPath := filepath.Join(folderToMove, fileName)
	fileToMove, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer fileToMove.Close()

	move, err := os.Create(destinationPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(move, fileToMove)
	if err != nil {
		return err
	}

	err = os.Remove(sourcePath)
	if err != nil {
		return err
	}

	return nil
}