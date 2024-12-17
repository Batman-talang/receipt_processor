package storage

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryStore(t *testing.T) {
	store := New()
	prints := 40
	id := store.Save(prints)
	assert.NotEqual(t, id, "")

	getPoints, found := store.Get(id)
	assert.Equal(t, found, true)
	assert.Equal(t, getPoints, prints)

	_, found = store.Get(uuid.NewString())
	assert.Equal(t, found, false)

}
