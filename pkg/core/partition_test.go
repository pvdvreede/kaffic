package core_test

import (
	"testing"

	"github.com/pvdvreede/kaffic/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestPartitionNextOffset(t *testing.T) {
	p := core.NewPartition()

	assert.Equal(t, p.NextOffset().ToInt64(), int64(1))
	assert.Equal(t, p.NextOffset().ToInt64(), int64(2))
	assert.Equal(t, p.NextOffset().ToInt64(), int64(3))
}
