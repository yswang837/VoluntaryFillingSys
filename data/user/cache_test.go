package user

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestClient_SetTTL(t *testing.T) {
	c, err := NewClient()
	assert.Equal(t, nil, err)
	b := c.SetTTL("AA0111111", "aaa", "10")
	fmt.Println("result...", b)
}
