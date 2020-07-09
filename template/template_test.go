package template

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersionCompare(t *testing.T) {
	assert.Equal(t, true, VersionCompare("v1.2.8", []string{"v1.2.7"}))
	assert.Equal(t, true, VersionCompare("v1.2.8", []string{"v1.2.8"}))
	assert.Equal(t, true, VersionCompare("v1.2.8", []string{"v1.2.5", "v1.2.8"}))
	assert.Equal(t, false, VersionCompare("v1.2.7", []string{"v1.2.8"}))
	assert.Equal(t, true, VersionCompare("v0.0.30", []string{"v0.0.30"}))
	assert.Equal(t, true, VersionCompare("v0.0.30", []string{">=v0.0.30"}))
	assert.Equal(t, true, VersionCompare("v0.0.30", []string{">=v0.0.29"}))
	assert.Equal(t, false, VersionCompare("v0.0.30", []string{">=v0.1.1"}))
	assert.Equal(t, true, VersionCompare("v0.0.30", []string{"<=v0.1.1"}))
}
