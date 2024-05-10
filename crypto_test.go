package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyEncryptDecrypt(t *testing.T) {
	payload := "Foo not Bar"
	src := bytes.NewReader([]byte(payload))
	dst := new(bytes.Buffer)
	key := newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)

	assert.Nil(t, err)

	fmt.Println(dst.Bytes())

	out := new(bytes.Buffer)
	nw, err := copyDecrypt(key, dst, out)

	assert.Nil(t, err)

	assert.Equal(t, 16+len(payload), nw)
}
