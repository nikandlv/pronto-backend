package storage

import (
	"fmt"
	"io"
	"nikan.dev/pronto/internals/dependencies"
	"os"
)

type localFileStorage struct {
	deps dependencies.CommonDependencies
	
}

func (l localFileStorage) Copy(reader io.Reader,from string) (bool, error) {
	dst, err := os.Create(from)
	if err != nil {
		return false, err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, reader); err != nil {
		return false, err
	}
	return true, nil
}

func (l localFileStorage) Remove(path string)  (bool, error) {
	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, nil
}

func NewLocalFileStorage(deps dependencies.CommonDependencies) localFileStorage {
	return localFileStorage{
		deps,
	}
}