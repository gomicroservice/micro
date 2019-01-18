package fields

import (
	"context"

	"github.com/gomicroservice/micro/backend/rest/graphql/types"
	"github.com/gomicroservice/micro/backend/rest/mongo"
	"github.com/graphql-go/graphql"
	"github.com/mongodb/mongo-go-driver/bson"
)

type todoStruct struct {
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var GetNotTodos = &graphql.Field{
	Type:        graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		notTodoCollection := mongo.Client.Database("medium-app").Collection("Not_Todos")
		todos, err := notTodoCollection.Find(context.Background(), nil)
		if err != nil {
			panic(err)
		}
		var todosList []todoStruct
		for todos.Next(context.Background()) {
			doc := bson.NewDocument()
			err := todos.Decode(doc)
			if err != nil {
				panic(err)
			}
			keys, err := doc.Keys(false)
			if err != nil {
				panic(err)
			}
			// convert BSON to struct
			todo := todoStruct{}
			for _, key := range keys {
				keyString := key.String()
				elm := doc.Lookup(keyString)
				switch keyString {
				case "name":
					todo.NAME = elm.StringValue()
				case "description":
					todo.DESCRIPTION = elm.StringValue()
				default:
				}
			}
			todosList = append(todosList, todo)
		}

		return todosList, nil
	},
}
