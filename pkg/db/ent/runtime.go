// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/lang"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/message"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	langFields := schema.Lang{}.Fields()
	_ = langFields
	// langDescCreateAt is the schema descriptor for create_at field.
	langDescCreateAt := langFields[4].Descriptor()
	// lang.DefaultCreateAt holds the default value on creation for the create_at field.
	lang.DefaultCreateAt = langDescCreateAt.Default.(func() uint32)
	// langDescUpdateAt is the schema descriptor for update_at field.
	langDescUpdateAt := langFields[5].Descriptor()
	// lang.DefaultUpdateAt holds the default value on creation for the update_at field.
	lang.DefaultUpdateAt = langDescUpdateAt.Default.(func() uint32)
	// lang.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	lang.UpdateDefaultUpdateAt = langDescUpdateAt.UpdateDefault.(func() uint32)
	// langDescDeleteAt is the schema descriptor for delete_at field.
	langDescDeleteAt := langFields[6].Descriptor()
	// lang.DefaultDeleteAt holds the default value on creation for the delete_at field.
	lang.DefaultDeleteAt = langDescDeleteAt.Default.(func() uint32)
	// langDescID is the schema descriptor for id field.
	langDescID := langFields[0].Descriptor()
	// lang.DefaultID holds the default value on creation for the id field.
	lang.DefaultID = langDescID.Default.(func() uuid.UUID)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescCreateAt is the schema descriptor for create_at field.
	messageDescCreateAt := messageFields[6].Descriptor()
	// message.DefaultCreateAt holds the default value on creation for the create_at field.
	message.DefaultCreateAt = messageDescCreateAt.Default.(func() uint32)
	// messageDescUpdateAt is the schema descriptor for update_at field.
	messageDescUpdateAt := messageFields[7].Descriptor()
	// message.DefaultUpdateAt holds the default value on creation for the update_at field.
	message.DefaultUpdateAt = messageDescUpdateAt.Default.(func() uint32)
	// message.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	message.UpdateDefaultUpdateAt = messageDescUpdateAt.UpdateDefault.(func() uint32)
	// messageDescDeleteAt is the schema descriptor for delete_at field.
	messageDescDeleteAt := messageFields[8].Descriptor()
	// message.DefaultDeleteAt holds the default value on creation for the delete_at field.
	message.DefaultDeleteAt = messageDescDeleteAt.Default.(func() uint32)
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
}
