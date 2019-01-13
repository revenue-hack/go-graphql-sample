package main

import (
	"errors"

	"github.com/graphql-go/graphql"
)

// define schema, with our rootQuery and rootMutation
var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    rootQuery,
	Mutation: rootMutation,
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user":     UserField,
		"userList": UserListField,
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": CreateUserField,
	},
})

var UserField = &graphql.Field{
	Type:        UserType, // 返り値の型
	Description: "Get single user",
	Args: graphql.FieldConfigArgument{ //引数の定義
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) { //実行関数
		userId, isOK := params.Args["id"].(string) // 引数取り出し
		if isOK {
			return FindUserById(userId)
		}

		return nil, errors.New("no userId")
	},
}

var UserListField = &graphql.Field{
	Type:        graphql.NewList(UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		users, err := UserList()
		if err != nil {
			panic(err)
		}
		return users, nil
	},
}

var CreateUserField = &graphql.Field{
	Type:        UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"userName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"photoURL": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userName, _ := params.Args["userName"].(string)
		description, _ := params.Args["description"].(string)
		photoURL, _ := params.Args["photoURL"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := NewUser(userName, description, photoURL, email)
		if err != nil {
			panic(err)
		}
		if err := Store(newUser); err != nil {
			panic(err)
		}
		return newUser, nil
	},
}

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userId": &graphql.Field{
			Type: graphql.String,
		},
		"userName": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"photoURL": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})
