package main

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathTransformFnc(t *testing.T) {
	key := "momsbestpicture"

	PathName := CASPathTransformFunc(key)

	assert.Equal(t, PathName.PathName, "68044/29f74/181a6/3c50c/3d81d/733a1/2f14a/353ff")
	assert.Equal(t, PathName.Filename, "6804429f74181a63c50c3d81d733a12f14a353ff")
}

func TestStore(t *testing.T) {
	s := newStore()
	defer teardown(t, s)
	for i := 0; i < 50; i++ {

		key := fmt.Sprintf("key_%d", i)
		data := ([]byte("some jpeg bytes"))

		if _, err := s.writeStream(key, bytes.NewReader(data)); err != nil {
			t.Error(err)
		}

		ok := s.Has(key)
		assert.True(t, ok)

		_, r, err := s.Read(key)
		assert.Nil(t, err)

		b, _ := io.ReadAll(r)

		assert.Equal(t, string(b), string(data))

		err = s.Delete(key)
		assert.Nil(t, err)

		has := s.Has(key)
		assert.False(t, has)
	}
}

func newStore() *Store {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
		Root:              "root",
	}

	return NewStore(opts)
}

func teardown(t *testing.T, s *Store) {
	err := s.Clear()

	assert.Nil(t, err)
}
