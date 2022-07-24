package noop_test

import (
	"testing"

	"github.com/pbten/graphql-go"
	"github.com/pbten/graphql-go/example/starwars"
	"github.com/pbten/graphql-go/trace/noop"
	"github.com/pbten/graphql-go/trace/tracer"
)

func TestInterfaceImplementation(t *testing.T) {
	var _ tracer.ValidationTracer = &noop.Tracer{}
	var _ tracer.Tracer = &noop.Tracer{}
}

func TestTracerOption(t *testing.T) {
	_, err := graphql.ParseSchema(starwars.Schema, nil, graphql.Tracer(noop.Tracer{}))
	if err != nil {
		t.Fatal(err)
	}
}
