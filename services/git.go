package services

import (
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
	"go.uber.org/zap"
)

func GetRepoFileSystem(url string) (billy.Filesystem, error) {
	fs := memfs.New()
	zap.L().Info(fs.Root())

	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func RecursivePrintAllFiles(fs billy.Filesystem) {
	zap.L().Info(fs.Root())
	if files, err := fs.ReadDir("/" + fs.Root()); err != nil {
		for _, file := range files {
			if file.IsDir() {
				if next, err := fs.Chroot(file.Name()); err != nil {
					RecursivePrintAllFiles(next)
				}
			} else {
				zap.L().Info(file.Name())
			}
		}
	}
}
