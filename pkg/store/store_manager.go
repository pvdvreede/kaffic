package store

import "io"

//go:generate mockgen -source=$GOFILE -destination=../mocks/mock_store.go -package=mocks

type StoreManager interface {
	GetMessageWriter() (io.Writer, error)
	GetIndexWriter() (io.Writer, error)
	GetMetaDataReadWriter() (io.ReadWriter, error)
}
