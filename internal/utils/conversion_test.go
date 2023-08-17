package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToByteArrayWithByteArray(t *testing.T) {
	v := []byte{100, 101, 102}
	byteArray, _ := ToByteArray(v)
	assert.Equal(t, []byte{100, 101, 102}, byteArray)
}

func TestToByteArrayWithInt(t *testing.T) {
	v := 999
	byteArray, _ := ToByteArray(v)
	assert.Equal(t, []byte{0, 0, 3, 231}, byteArray)
}

func TestToByteArrayWithString(t *testing.T) {
	v := "hello"
	byteArray, _ := ToByteArray(v)
	assert.Equal(t, []byte{104, 101, 108, 108, 111}, byteArray)
}
