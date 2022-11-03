package generator

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go/format"
	"strings"

	"dagger.io/dagger"
	"github.com/dagger/dagger/codegen/generator/go/templates"
	tsgenerator "github.com/dagger/dagger/codegen/generator/ts"
	"github.com/dagger/dagger/codegen/introspection"
)

var ErrUnknownSDK = errors.New("unknown sdk language")

type SDKLang string

const (
	SDKLangUnknown SDKLang = ""
	SDKLangGo      SDKLang = "Go"
	SDKLangTS      SDKLang = "TypeScript"
)

type Config struct {
	Lang    SDKLang
	Package string
}

type Generator interface {
	Generate(ctx context.Context) ([]byte, error)
}

func Generate(ctx context.Context, schema *introspection.Schema, cfg Config) ([]byte, error) {
	// Set parent objects for fields
	for _, t := range schema.Types {
		for _, f := range t.Fields {
			f.ParentObject = t
		}
	}

	var gen Generator
	switch cfg.Lang {
	case SDKLangGo:
		gen = &GoGenerator{
			cfg:    cfg,
			schema: schema,
		}

	default:
		return []byte{}, ErrUnknownSDK
	}

	return gen.Generate(ctx)
}

func IntrospectAndGenerate(ctx context.Context, c *dagger.Client, cfg Config) ([]byte, error) {
	var response introspection.Response
	err := c.Do(ctx,
		&dagger.Request{
			Query: introspection.Query,
		},
		&dagger.Response{Data: &response},
	)
	if err != nil {
		return nil, fmt.Errorf("error querying the API: %w", err)
	}

	return Generate(ctx, response.Schema, cfg)
}

type GoGenerator struct {
	cfg    Config
	schema *introspection.Schema
}

func (g *GoGenerator) Generate(_ context.Context) ([]byte, error) {
	headerData := struct {
		Package string
		Schema  *introspection.Schema
	}{
		Package: g.cfg.Package,
		Schema:  g.schema,
	}
	var header bytes.Buffer
	if err := templates.Header.Execute(&header, headerData); err != nil {
		return nil, err
	}

	render := []string{
		header.String(),
	}

	err := g.schema.Visit(introspection.VisitHandlers{
		Scalar: func(t *introspection.Type) error {
			var out bytes.Buffer
			if err := templates.Scalar.Execute(&out, t); err != nil {
				return err
			}
			render = append(render, out.String())
			return nil
		},
		Object: func(t *introspection.Type) error {
			var out bytes.Buffer
			if err := templates.Object.Execute(&out, t); err != nil {
				return err
			}
			render = append(render, out.String())
			return nil
		},
		Input: func(t *introspection.Type) error {
			var out bytes.Buffer
			if err := templates.Input.Execute(&out, t); err != nil {
				return err
			}
			render = append(render, out.String())
			return nil
		},
	})
	if err != nil {
		return nil, err
	}

	formatted, err := format.Source(
		[]byte(strings.Join(render, "\n")),
	)
	if err != nil {
		return nil, fmt.Errorf("error formatting generated code: %w", err)
	}
	return formatted, nil
}
