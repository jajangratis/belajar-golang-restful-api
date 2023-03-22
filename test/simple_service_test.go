package test

import (
	"github.com/stretchr/testify/assert"
	"jajangratis/belajar-golang-restful-api/simple"
	"testing"
)

func TestSimpleServiceError(t *testing.T) {
	simpleService, err := simple.InitializedService(true)
	assert.Nil(t, simpleService)
	assert.NotNil(t, err)
}

func TestSimpleServiceSuccess(t *testing.T) {
	simpleService, err := simple.InitializedService(false)
	assert.NotNil(t, simpleService)
	assert.Nil(t, err)
}
