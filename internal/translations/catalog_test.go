// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package translations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestLookup tests the Lookup method of the dictionary type.
func TestLookup(t *testing.T) {
	d := &dictionary{
		index: []uint32{0, 10, 11, 12, 13, 14, 15},
		data:  "abcdefghijklmnopqrstuvwxyz",
	}

	testCases := []struct {
		name     string
		key      string
		expected string
		ok       bool
	}{
		{
			name:     "non-existing key",
			key:      "non-existing key",
			expected: "",
			ok:       false,
		},
		{
			name:     "existing key",
			key:      "'-w %d': value must be greater than 8 and less than 65536.",
			expected: "k",
			ok:       true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, ok := d.Lookup(tc.key)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.ok, ok)
		})
	}
}
