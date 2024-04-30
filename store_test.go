package main

import (
	"bytes"
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
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "momsspecials"
	data := ([]byte("some jpeg bytes"))

	if err := s.writeStream(key, bytes.NewReader(data)); err != nil {
		t.Error(err)
	}

	r, err := s.Read(key)
	assert.Nil(t, err)

	b, _ := io.ReadAll(r)

	assert.Equal(t, string(b), string(data))

	err = s.Delete(key)
	assert.Nil(t, err)
}

func TestStoreDeleteKey(t *testing.T) {
	opts := StoreOpts{
		PathTransformFunc: CASPathTransformFunc,
	}

	s := NewStore(opts)
	key := "momsspecials"
	data := ([]byte("some jpeg bytes"))

	err := s.writeStream(key, bytes.NewReader(data))
	assert.Nil(t, err)

	err = s.Delete(key)

	assert.Nil(t, err)
}
