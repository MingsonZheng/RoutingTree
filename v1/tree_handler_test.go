package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHandlerBasedOnTree(t *testing.T) {
	handler := NewHandlerBasedOnTree().(*HandlerBasedOnTree)
	assert.NotNil(t, handler.root)

	handler.Route(http.MethodPost, "/user", func(c *Context) {})

	// 开始做断言，这个时候我们应该确认，在根节点之下只有一个user节点
	assert.Equal(t, 1, len(handler.root.children))

	n := handler.root.children[0]
	assert.NotNil(t, n)
	assert.Equal(t, "user", n.path)
	assert.NotNil(t, n.handler)
	assert.Empty(t, n.children)
}
