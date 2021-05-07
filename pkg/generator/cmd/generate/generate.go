package generate

import (
	"context"
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/hori-ryota/go-genutil/genutil"
	"github.com/spf13/cobra"

	"github.com/stormcat24/go-assign/pkg/cli"
	"github.com/stormcat24/go-assign/pkg/generator/render"
)

type generator struct {
}

func NewCommand() *cobra.Command {
	v := generator{}
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate an inherited struct with go-assign tag.",
		RunE:  cli.WithContext(v.run),
	}
	return cmd
}

func (g *generator) run(ctx context.Context) error {
	cd, err := os.Getwd()
	if err != nil {
		return err
	}
	return filepath.WalkDir(cd, func(path string, d os.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}

		genItems, err := g.process(path)
		if err != nil {
			return err
		}

		for _, gen := range genItems {
			data, err := render.Render(&gen)
			if err != nil {
				return err
			}
			dstPath := filepath.Join(cd, gen.FileName)
			if err := os.WriteFile(dstPath, data, 0644); err != nil {
				return err
			}
		}
		return nil
	})
}

func (g *generator) process(dir string) ([]render.GeneratedParam, error) {
	walkers, err := genutil.DirToAstWalker(dir, func(fi os.FileInfo) bool {
		return !fi.IsDir()
	})
	if err != nil {
		return nil, err
	}

	generatedParams := make([]render.GeneratedParam, 0)
	for _, walker := range walkers {
		for _, spec := range walker.AllStructSpecs() {
			structType := spec.Type.(*ast.StructType)

			fileFields := make([]*render.FileField, 0)
			for _, field := range structType.Fields.List {
				if field.Tag == nil {
					continue
				}
				tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
				value, exist := tag.Lookup(TagName)
				if !exist {
					continue
				}
				ff, err := extractFileField(field.Names[0].Name, value)
				if err != nil {
					return nil, err
				}
				fileFields = append(fileFields, ff)
			}

			if len(fileFields) > 0 {
				p := render.GeneratedParam{
					Package: walker.Pkg.Name,
					Extend: render.Struct{
						Name:   fmt.Sprintf("%sGenerated", spec.Name.Name),
						Parent: spec.Name.Name,
						Fields: fileFields,
					},
					FileName: fmt.Sprintf("%s_gen.go", strings.ToLower(spec.Name.Name)),
				}
				generatedParams = append(generatedParams, p)
			}
		}
	}
	return generatedParams, nil
}
