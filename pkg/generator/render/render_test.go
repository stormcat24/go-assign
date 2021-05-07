package render

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {

	p := GeneratedParam{
		Package: "example",
		Extend: Struct{
			Name:   "ConfigGenerated",
			Parent: "Config",
			Fields: []*FileField{
				{
					Name:             "HelloAPITokenFile",
					LocalName:        "helloAPITokenFile",
					AssignTargetName: "HelloAPIToken",
					Tag:              `json:"helloAPITokenFile"`,
					Base64:           false,
				},
				{
					Name:             "WorldAPITokenFile",
					LocalName:        "worldAPITokenFile",
					AssignTargetName: "WorldAPIToken",
					Tag:              `json:"worldAPITokenFile"`,
					Base64:           true,
				},
			},
		},
	}

	_, err := Render(&p)
	assert.NoError(t, err)
}
