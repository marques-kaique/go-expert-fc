//go:build tools
// +build tools

package tools // por conversão se utiliza o tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)