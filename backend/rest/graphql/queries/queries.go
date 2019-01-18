package queries

import (
	"github.com/gomicroservice/micro/backend/rest/graphql/queries/fields"
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getNotTodos": fields.GetNotTodos,
	},
})
