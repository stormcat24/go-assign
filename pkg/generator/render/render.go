package render

import (
	"bytes"
	"embed"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
)

//go:embed template/*.tmpl
var TemplateFile embed.FS

type GeneratedParam struct {
	EarlierGo116 bool
	Package      string
	Extend       Struct
	FileName     string
}

type Struct struct {
	Name   string
	Parent string
	Fields []*FileField
}

func (s *Struct) HasBase64Encoding() bool {
	for _, f := range s.Fields {
		if f.Base64 {
			return true
		}
	}
	return false
}

type FileField struct {
	Name             string
	LocalName        string
	AssignTargetName string
	Tag              string
	Base64           bool
}

func Render(param *GeneratedParam) ([]byte, error) {
	// TODO: Implement Go1.16 earlier
	tmpl, err := template.ParseFS(TemplateFile, "template/gen_go1.16.tmpl")
	if err != nil {
		return nil, err
	}
	rawFile, err := os.CreateTemp("", "*.go")
	if err != nil {
		return nil, err
	}
	defer rawFile.Close()

	buf := bytes.NewBufferString("")
	if err := tmpl.Execute(buf, param); err != nil {
		return nil, err
	}

	// Default imports setting
	data, err := imports.Process(rawFile.Name(), buf.Bytes(), nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}
