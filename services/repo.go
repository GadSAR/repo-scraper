package services

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/storage/memory"
)

func GetRepoFileSystem(url string) (billy.Filesystem, error) {
	fs := memfs.New()

	if _, err := git.Clone(memory.NewStorage(), fs,
		&git.CloneOptions{
			URL: url,
		}); err != nil {
		return nil, err
	}

	return fs, nil
}
