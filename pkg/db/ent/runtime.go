// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/applang"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/country"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/lang"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/message"
	"github.com/NpoolPlatform/internationalization/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	applangFields := schema.AppLang{}.Fields()
	_ = applangFields
	// applangDescMainLang is the schema descriptor for main_lang field.
	applangDescMainLang := applangFields[3].Descriptor()
	// applang.DefaultMainLang holds the default value on creation for the main_lang field.
	applang.DefaultMainLang = applangDescMainLang.Default.(bool)
	// applangDescCreateAt is the schema descriptor for create_at field.
	applangDescCreateAt := applangFields[4].Descriptor()
	// applang.DefaultCreateAt holds the default value on creation for the create_at field.
	applang.DefaultCreateAt = applangDescCreateAt.Default.(func() uint32)
	// applangDescUpdateAt is the schema descriptor for update_at field.
	applangDescUpdateAt := applangFields[5].Descriptor()
	// applang.DefaultUpdateAt holds the default value on creation for the update_at field.
	applang.DefaultUpdateAt = applangDescUpdateAt.Default.(func() uint32)
	// applang.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	applang.UpdateDefaultUpdateAt = applangDescUpdateAt.UpdateDefault.(func() uint32)
	// applangDescDeleteAt is the schema descriptor for delete_at field.
	applangDescDeleteAt := applangFields[6].Descriptor()
	// applang.DefaultDeleteAt holds the default value on creation for the delete_at field.
	applang.DefaultDeleteAt = applangDescDeleteAt.Default.(func() uint32)
	// applangDescID is the schema descriptor for id field.
	applangDescID := applangFields[0].Descriptor()
	// applang.DefaultID holds the default value on creation for the id field.
	applang.DefaultID = applangDescID.Default.(func() uuid.UUID)
	countryFields := schema.Country{}.Fields()
	_ = countryFields
	// countryDescCreateAt is the schema descriptor for create_at field.
	countryDescCreateAt := countryFields[5].Descriptor()
	// country.DefaultCreateAt holds the default value on creation for the create_at field.
	country.DefaultCreateAt = countryDescCreateAt.Default.(func() uint32)
	// countryDescUpdateAt is the schema descriptor for update_at field.
	countryDescUpdateAt := countryFields[6].Descriptor()
	// country.DefaultUpdateAt holds the default value on creation for the update_at field.
	country.DefaultUpdateAt = countryDescUpdateAt.Default.(func() uint32)
	// country.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	country.UpdateDefaultUpdateAt = countryDescUpdateAt.UpdateDefault.(func() uint32)
	// countryDescDeleteAt is the schema descriptor for delete_at field.
	countryDescDeleteAt := countryFields[7].Descriptor()
	// country.DefaultDeleteAt holds the default value on creation for the delete_at field.
	country.DefaultDeleteAt = countryDescDeleteAt.Default.(func() uint32)
	// countryDescID is the schema descriptor for id field.
	countryDescID := countryFields[0].Descriptor()
	// country.DefaultID holds the default value on creation for the id field.
	country.DefaultID = countryDescID.Default.(func() uuid.UUID)
	langFields := schema.Lang{}.Fields()
	_ = langFields
	// langDescCreateAt is the schema descriptor for create_at field.
	langDescCreateAt := langFields[5].Descriptor()
	// lang.DefaultCreateAt holds the default value on creation for the create_at field.
	lang.DefaultCreateAt = langDescCreateAt.Default.(func() uint32)
	// langDescUpdateAt is the schema descriptor for update_at field.
	langDescUpdateAt := langFields[6].Descriptor()
	// lang.DefaultUpdateAt holds the default value on creation for the update_at field.
	lang.DefaultUpdateAt = langDescUpdateAt.Default.(func() uint32)
	// lang.UpdateDefaultUpdateAt holds the default value on update for the update_at field.
	lang.UpdateDefaultUpdateAt = langDescUpdateAt.UpdateDefault.(func() uint32)
	// langDescDeleteAt is the schema descriptor for delete_at field.
	langDescDeleteAt := langFields[7].Descriptor()
	// lang.DefaultDeleteAt holds the default value on creation for the delete_at field.
	lang.DefaultDeleteAt = langDescDeleteAt.Default.(func() uint32)
	// langDescID is the schema descriptor for id field.
	langDescID := langFields[0].Descriptor()
	// lang.DefaultID holds the default value on creation for the id field.
	lang.DefaultID = langDescID.Default.(func() uuid.UUID)
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescMessage is the schema descriptor for message field.
	messageDescMessage := messageFields[4].Descriptor()
	// message.MessageValidator is a validator for the "message" field. It is called by the builders before save.
	message.MessageValidator = messageDescMessage.Validators[0].(func(string) error)
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
