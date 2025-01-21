package services

import (
	"os"

	"github.com/GadSAR/repo-scraper/models"
	"github.com/go-git/go-billy/v5"
)

func ExceedThreshold(files []os.FileInfo, threshold int64) []os.FileInfo {
	var excceedFiles []os.FileInfo

	for _, file := range files {
		if file.Size() > threshold {
			excceedFiles = append(excceedFiles, file)
		}
	}

	return excceedFiles
}

func FormatFiles(files []os.FileInfo) []models.File {
	var formattedFiles []models.File

	for _, file := range files {
		formattedFiles = append(formattedFiles, models.File{
			Name: file.Name(),
			Size: file.Size(),
		})
	}

	return formattedFiles
}

func TraverseFiles(fs billy.Filesystem, array []os.FileInfo) ([]os.FileInfo, error) {
	files, err := fs.ReadDir("/")

	if err != nil {
		return array, err
	}

	for _, file := range files {
		if file.IsDir() {
			dirfs, err := fs.Chroot(file.Name())

			if err != nil {
				return array, err
			}

			array, err = TraverseFiles(dirfs, array)

			if err != nil {
				return array, err
			}
		} else {
			array = append(array, file)
		}
	}

	return array, nil
}
