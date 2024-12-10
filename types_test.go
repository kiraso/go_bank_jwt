package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewAccount(t *testing.T) {
	acc, err := NewAccount("a", "b", "annie")
	assert.Nil(t, err)
	fmt.Printf("%+v\n", acc) // ใช้ Printf อย่างถูกต้อง
}
