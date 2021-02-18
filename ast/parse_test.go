package ast

import (
	"strings"
	"testing"

	"github.com/hexops/autogold"
	"github.com/vektah/gqlparser/v2/ast"
)

func Test_filesToAstSources(t *testing.T) {
	got, err := filesToAstSources("testdata/simple-schema.graphql")
	if err != nil {
		t.Errorf("expected to parse schema file: %#v", err)
		return
	}
	autogold.Equal(t, got)
}

func Test_astSchemaDocToAstSchema(t *testing.T) {

	t.Run("translates schema doc into schema struct", func(t *testing.T) {
		doc := &ast.SchemaDocument{
			Definitions: []*ast.Definition{
				{
					Kind: "OBJECT",
					Name: "User",
					Fields: []*ast.FieldDefinition{
						{
							Name: "id",
							Type: &ast.Type{
								NamedType: "ID",
								NonNull:   true,
							},
						},
						{
							Name: "name",
							Type: &ast.Type{
								NamedType: "String",
								NonNull:   false,
							},
						},
					},
				},
			},
		}

		got, err := astSchemaDocToAstSchema(doc)
		if err != nil {
			t.Errorf("TestParseSchemaDoc() error = %v", err)
		}
		autogold.Equal(t, got, autogold.Dir(""))
	})
}

// uncomment the below tests out. Will get stuck in an infinite loop
//
func Test_astSourceToAstSchema(t *testing.T) {

	t.Run("multiple sources", func(t *testing.T) {
		sources := []*ast.Source{
			{
				Name: "scalars.graphql",
				Input: `
					# resolves to time.Time
					scalar Time

					# resolves to map[string]interface{}
					scalar Map

					# resolves to interface{}
					scalar Any

					# resolves to string of type Password
					scalar Password
				`,
				BuiltIn: false,
			},
			{
				Name: "schema.graphql",
				Input: `
					schema {
						query: Query
						mutation: Mutation
					}

					type Query {
						user(id: ID!): User
						users: [User]
						me: User
					}

					type Mutation {
						createUser(input: UserInput!): User!
						updateUser(id: ID!, input: UserInput!): User!
					}

					type User {
						id: ID!
						createdAt: Time!
						updatedAt: Time!
						name: String
					}

					input UserInput {
						name: String
					}
				`,
				BuiltIn: false,
			},
		}
		got, err := astSourceToAstSchema(sources...)
		if err != nil {
			t.Errorf("expected to load schema: %#v", err)
			return
		}
		autogold.Equal(t, got)
	})
}

func Test_sourceToSchemaDoc(t *testing.T) {

	t.Run("parse simple source file", func(t *testing.T) {
		source := []*ast.Source{
			{
				Name: "schema.graphql",
				Input: strings.TrimSpace(`
type Query {
  me: User
}

type Mutation {
  createUser(input: UserInput!): User!
}

type User{
  id: ID!
  createdAt: Time!
  updatedAt: Time!
  name: String
}

input UserInput{
  name: String
}
`),
				BuiltIn: false,
			},
		}

		got, err := sourceToSchemaDoc(source...)
		if err != nil {
			t.Errorf("TestSourceToSchemaDoc() error = %v", err)
		}
		autogold.Equal(t, got)
	})
}
