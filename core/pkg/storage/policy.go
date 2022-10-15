package storage

import (
	"encoding/json"
	"fmt"
	"github.com/obarbier/custom-app/core/pkg/log_utils"
	"github.com/obarbier/custom-app/core/pkg/models"
)

const DefaultPolicyFormat = `
{
	"user/%d":{
				"capabilities": [ "read", "update"]
			}
}
`

func SetDefaultPolicy(id int64) string {
	return fmt.Sprintf(DefaultPolicyFormat, id)
}

// TODO(obarbier): can we user golang custom decoder to do this
func UnmarshalPolicy(data []byte) *models.Policy {
	var v models.Policy
	err := json.Unmarshal(data, &v)
	if err != nil {
		log_utils.Error(err)
	}
	return &v
}
