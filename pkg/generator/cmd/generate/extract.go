package generate

import (
	"fmt"
	"strings"

	"github.com/stoewer/go-strcase"

	"github.com/stormcat24/go-assign/pkg/generator/render"
)

const (
	TagName = "go-assign"

	AttributeFileFieldName = "fileFieldName"
	AttributeFileFieldTag  = "fileFieldTag"
	AttributeBase64        = "base64"
)

func extractFileField(assignTargetName, value string) (*render.FileField, error) {
	ff := render.FileField{
		AssignTargetName: assignTargetName,
	}
	pairs := strings.Split(value, ",")
	for _, pair := range pairs {
		tokens := strings.Split(pair, "=")
		if len(tokens) != 2 {
			return nil, fmt.Errorf("invalid go-assign tag value: %s", pair)
		}

		k := tokens[0]
		v := tokens[1]
		switch k {
		case AttributeFileFieldName:
			ff.Name = v
			ff.LocalName = strcase.LowerCamelCase(v)
		case AttributeFileFieldTag:
			ff.Tag = v
		case AttributeBase64:
			ff.Base64 = v == "true"
		default:
			return nil, fmt.Errorf("invalid go-assign attribute name: %s", k)
		}
	}
	return &ff, nil
}
