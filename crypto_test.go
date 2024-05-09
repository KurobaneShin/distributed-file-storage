package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCopyEncryptDecrypt(t *testing.T) {
	src := bytes.NewReader([]byte("Foo not Bar"))
	dst := new(bytes.Buffer)
	key := newEncryptionKey()
	_, err := copyEncrypt(key, src, dst)

	assert.Nil(t, err)

	fmt.Println(dst.Bytes())
}

func Test_newEncryptionKey(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newEncryptionKey(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEncryptionKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
