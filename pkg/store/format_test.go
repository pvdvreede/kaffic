package store_test

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pvdvreede/ms-core/pkg/core"
	"github.com/pvdvreede/ms-core/pkg/store"
)

func TestToFileReader(t *testing.T) {
	input := &store.MessageFileFormat{}
	r, err := store.ToFileReader(input)
	assert.Nil(t, err)
	actual, err := ioutil.ReadAll(r)
	assert.Nil(t, err)
	assert.NotEmpty(t, actual)
}

func TestToMessageFileFormat(t *testing.T) {
	input := core.NewInboundMessage(core.FromOffset(12345), time.Unix(12345678, 11), []byte("the payload"))

	output := store.ToMessageFileFormat(input)

	assert.NotNil(t, output)
	assert.Equal(t, output, &store.MessageFileFormat{
		Offset:    12345,
		Payload:   []byte("the payload"),
		Version:   "v1",
		Timestamp: 12345678000000011,
	})
}
