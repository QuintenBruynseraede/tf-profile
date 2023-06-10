package tfprofile

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicRun(t *testing.T) {
	err := Table([]string{}, 1, true, "tot_time=asc", true)
	assert.Nil(t, err)
}

func TestFileDoesntExist(t *testing.T) {
	err := Table([]string{"does-not-exist"}, 1, true, "tot_time=asc", true)
	assert.NotNil(t, err)
}
