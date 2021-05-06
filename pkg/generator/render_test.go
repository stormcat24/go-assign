package generator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender(t *testing.T) {
	p := Param{
		Package: "example",
		Extend: Struct{
			Name:   "ConfigGenerated",
			Parent: "Config",
			Fields: []Field{
				{
					Name:       "HelloAPITokenFile",
					LocalName:  "helloAPITokenFile",
					SourceName: "HelloAPIToken",
					Tag:        `json:"helloAPITokenFile"`,
					Base64:     false,
				},
				{
					Name:       "WorldAPITokenFile",
					LocalName:  "worldAPITokenFile",
					SourceName: "WorldAPIToken",
					Tag:        `json:"worldAPITokenFile"`,
					Base64:     true,
				},
			},
		},
	}

	_, err := Render(&p)
	assert.NoError(t, err)
}
