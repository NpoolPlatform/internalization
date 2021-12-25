package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"

	constant "github.com/NpoolPlatform/internationalization/pkg/const" //nolint
)

// Message holds the schema definition for the Message entity.
type Message struct {
	ent.Schema
}

// Fields of the Message.
func (Message) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).
			Unique(),
		field.UUID("app_id", uuid.UUID{}),
		field.UUID("message_id", uuid.UUID{}),
		field.Enum("lang").
			Values(
				constant.LangZhCN,
				constant.LangEnUS,
			),
		field.String("message"),
		field.Uint32("create_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("update_at").
			DefaultFunc(func() uint32 {
				return uint32(time.Now().Unix())
			}).
			UpdateDefault(func() uint32 {
				return uint32(time.Now().Unix())
			}),
		field.Uint32("delete_at").
			DefaultFunc(func() uint32 {
				return 0
			}),
	}
}

// Edges of the Message.
func (Message) Edges() []ent.Edge {
	return nil
}

// Indexes of the Message
func (Message) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("app_id", "message_id").
			Unique(),
	}
}
