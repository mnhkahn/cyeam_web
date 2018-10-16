// Package util
package util

import (
	"testing"

	"github.com/mnhkahn/gogogo/logger"

	"github.com/issue9/assert"
)

func TestJsonToThrift(t *testing.T) {
	res, err := JsonToThrift(`{"id":1,"type":"2"}`)
	assert.Nil(t, err)
	assert.NotEmpty(t, res)
	logger.Info(string(res))
}
