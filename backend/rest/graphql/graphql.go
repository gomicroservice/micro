package graphql

import (
	"github.com/gomicroservice/micro/backend/rest/graphql/mutations"
	"github.com/gomicroservice/micro/backend/rest/graphql/queries"
	"github.com/graphql-go/graphql"
)

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.RootQuery,
	Mutation: mutations.RootMutation,
})
