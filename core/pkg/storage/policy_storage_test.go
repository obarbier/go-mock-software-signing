package storage

import (
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrieTree(t *testing.T) {
	// create tree
	paths := []string{"/user/*", "/user/1", "/user/all", "/key/*"}
	r := newNode()

	for _, path := range paths {
		r.insert(path, models.PolicyAnon{})
	}

	type Test struct {
		name    string
		data    []string
		wants   []bool
		wantErr bool
	}

	tests := []Test{
		{
			name:    "Positive test",
			data:    []string{"/user/*", "/user/1", "/user/all", "/key/*"},
			wants:   []bool{true, true, true, true},
			wantErr: false,
		},

		{
			name:    "Negative test",
			data:    []string{"/abc/*", "/user1/1", "/user", "/key*/*"},
			wants:   []bool{false, false, false, false},
			wantErr: false,
		},

		{
			name:    "Mixed test",
			data:    []string{"/user/*", "/user/1", "/user", "/key*/*"},
			wants:   []bool{true, true, false, false},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for idx := range tt.data {
				got := r.isInTree(tt.data[idx])
				assert.Equal(t, tt.wants[idx], got)
			}
		})
	}
}
