package topic_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pvdvreede/kaffic/pkg/mocks"
	"github.com/pvdvreede/kaffic/pkg/topic"

	"github.com/stretchr/testify/assert"
)

func TestTopicWriter(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	storeMock := mocks.NewMockStoreManager(ctrl)

	w := new(bytes.Buffer)

	storeMock.EXPECT().GetMessageWriter().Return(w, nil).Times(1)

	subj := topic.NewTopicWriter("test-topic", storeMock)
	off, err := subj.WriteMessage([]byte("test message"))

	assert.Nil(t, err)

	_, err = ioutil.ReadAll(w)

	assert.Nil(t, err)
	assert.Equal(t, off.ToInt64(), int64(1))
}
