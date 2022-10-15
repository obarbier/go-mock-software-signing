package storage

import (
	"github.com/obarbier/custom-app/core/pkg/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetPolicyBit(t *testing.T) {
	type args struct {
		cs []models.Capability
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "positive test",
			args: args{cs: []models.Capability{models.CapabilityRead, models.CapabilityDelete}},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SetPolicyBit(tt.args.cs), "SetPolicyBit(%v)", tt.args.cs)
		})
	}
}
