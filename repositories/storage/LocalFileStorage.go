package storage

import (
	"fmt"
	"io"
	"nikan.dev/pronto/internals/dependencies"
	"nikan.dev/pronto/internals/entity"
	"os"
)

type localFileStorage struct {
	deps dependencies.CommonDependencies
	
}

func (l localFileStorage) Store(file entity.FileEntity, path string) error {
	dst, err := os.Create(fmt.Sprintf("%v/%v",path,file.Name))
	if err != nil {
		return err
	}
	defer dst.Close()
	if _, err = io.Copy(dst, file.Reader); err != nil {
		return  err
	}
	return nil

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