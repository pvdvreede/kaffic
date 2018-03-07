package core

import (
	"errors"
	"io"
	"time"

	"github.com/pvdvreede/ms-core/pkg/core"
	"github.com/pvdvreede/ms-core/pkg/store"
)

type TopicWriter struct {
	Topic        string
	StoreManager store.StoreManager
	// This is always set to one for now, TODO scaling later
	Partitions []*core.Partition
}

func (topicWriter *TopicWriter) WriteMessage(msg []byte) (*core.Offset, error) {
	ts := time.Now()

	offset := topicWriter.Partitions[0].NextOffset()

	r, err := store.ToFileReader(store.ToMessageFileFormat(core.NewInboundMessage(offset, ts, msg)))

	if err != nil {
		return nil, err
	}

	mw, err := topicWriter.StoreManager.GetMessageWriter()

	if err != nil {
		return nil, err
	}

	bw, err := io.Copy(mw, r)

	if err != nil {
		return nil, err
	}

	if bw == 0 {
		return nil, errors.New("Nothing was written to the underlying storage.")
	}

	// TODO: Update index

	// TODO: Check if new Segment should be created

	return &offset, nil
}

func NewTopicWriter(topicName string, storage store.StoreManager) *TopicWriter {
	return &TopicWriter{
		Partitions:   []*core.Partition{core.NewPartition()},
		Topic:        topicName,
		StoreManager: storage,
	}
}
