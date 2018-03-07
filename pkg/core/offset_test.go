package core_test

import (
	"testing"

	"github.com/pvdvreede/ms-core/pkg/core"
	"github.com/stretchr/testify/assert"
)

func TestInitialOffset(t *testing.T) {
	o := core.InitialOffset()
	n1 := o.Next()

	assert.NotEqual(t, o, n1)
	assert.Equal(t, o.ToInt64(), int64(1))

	assert.Equal(t, n1.ToInt64(), int64(2))

	assert.Equal(t, n1.String(), "000000000000000000000000000002")
}

func TestFromOffset(t *testing.T) {
	o := core.FromOffset(11111)
	n1 := o.Next()

	assert.NotEqual(t, o, n1)
	assert.Equal(t, o.ToInt64(), int64(11111))

	assert.Equal(t, n1.ToInt64(), int64(11112))

	assert.Equal(t, n1.String(), "000000000000000000000000011112")
}
