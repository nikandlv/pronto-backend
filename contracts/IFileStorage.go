package contracts

import "io"

type IFileStorage interface
{
	Copy(reader io.Reader, from string) (bool, error)
	Remove(name string) (bool, error)
}