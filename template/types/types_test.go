package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFieldOptions_SetSelected(t *testing.T) {
	var fo = FieldOptions{
		{"value": "123", "selected": ""},
		{"value": "234", "selected": ""},
	}
	fo.SetSelected("123", []string{"selected", ""})
	assert.Equal(t, fo[0]["selected"], "selected")

	var fo1 = FieldOptions{
		{"value": "123"},
		{"value": "234"},
	}
	fo1.SetSelected("123", []string{"selected", ""})
	assert.Equal(t, fo1[0]["selected"], "selected")
}
