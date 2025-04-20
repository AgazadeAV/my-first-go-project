package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),

		field.String("first_name").
			NotEmpty(),

		field.String("last_name").
			NotEmpty(),

		field.String("username").
			NotEmpty().
			Unique(),

		field.String("email").
			NotEmpty().
			Unique(),

		field.String("phone_number").
			NotEmpty().
			Unique(),

		field.Time("birth_date").
			SchemaType(map[string]string{
				"postgres": "date",
			}),
	}
}

func (User) Table() string {
	return "myschema.users"
}
