package api

import "github.com/obarbier/custom-app/core/pkg/models"

const (
	Read = 1 << iota
	Create
	Update
	Delete
)

var CapabilityBit = map[models.Capability]int{
	models.CapabilityCreate: Create,
	models.CapabilityUpdate: Update,
	models.CapabilityDelete: Delete,
	models.CapabilityRead:   Read,
}

var HTTPMethodMatch = map[string]int{
	"GET":    Read,
	"PUT":    Update,
	"DELETE": Delete,
	"POST":   Create,
}

func SetPolicyBit(cs []models.Capability) int {
	var res int
	for _, c := range cs {
		res |= CapabilityBit[c]
	}
	return res
}
