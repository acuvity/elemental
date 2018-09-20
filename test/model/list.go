package testmodel

import (
	"fmt"
	"sync"

	"time"

	"go.aporeto.io/elemental"
)

// ListIdentity represents the Identity of the object.
var ListIdentity = elemental.Identity{
	Name:     "list",
	Category: "lists",
	Private:  false,
}

// ListsList represents a list of Lists
type ListsList []*List

// Identity returns the identity of the objects in the list.
func (o ListsList) Identity() elemental.Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the ListsList.
func (o ListsList) Copy() elemental.Identifiables {

	copy := append(ListsList{}, o...)
	return &copy
}

// Append appends the objects to the a new copy of the ListsList.
func (o ListsList) Append(objects ...elemental.Identifiable) elemental.Identifiables {

	out := append(ListsList{}, o...)
	for _, obj := range objects {
		out = append(out, obj.(*List))
	}

	return out
}

// List converts the object to an elemental.IdentifiablesList.
func (o ListsList) List() elemental.IdentifiablesList {

	out := elemental.IdentifiablesList{}
	for _, item := range o {
		out = append(out, item)
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ListsList) DefaultOrder() []string {

	return []string{}
}

// Version returns the version of the content.
func (o ListsList) Version() int {

	return 1
}

// List represents the model of a list
type List struct {
	// The identifier.
	ID string `json:"ID" bson:"_id" mapstructure:"ID,omitempty"`

	// This attribute is creation only.
	CreationOnly string `json:"creationOnly" bson:"creationonly" mapstructure:"creationOnly,omitempty"`

	// The date.
	Date time.Time `json:"date" bson:"date" mapstructure:"date,omitempty"`

	// The description.
	Description string `json:"description" bson:"description" mapstructure:"description,omitempty"`

	// The name.
	Name string `json:"name" bson:"name" mapstructure:"name,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// This attribute is readonly.
	ReadOnly string `json:"readOnly" bson:"readonly" mapstructure:"readOnly,omitempty"`

	// this is a slice.
	Slice []string `json:"slice" bson:"slice" mapstructure:"slice,omitempty"`

	// This attribute is not exposed.
	Unexposed string `json:"-" bson:"unexposed" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" bson:"_modelversion"`

	sync.Mutex `json:"-" bson:"-"`
}

// NewList returns a new *List
func NewList() *List {

	return &List{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *List) Identity() elemental.Identity {

	return ListIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *List) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *List) SetIdentifier(id string) {

	o.ID = id
}

// Version returns the hardcoded version of the model.
func (o *List) Version() int {

	return 1
}

// DefaultOrder returns the list of default ordering fields.
func (o *List) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *List) Doc() string {
	return `Represent a a list of task to do.`
}

func (o *List) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *List) GetName() string {

	return o.Name
}

// SetName sets the given Name of the receiver.
func (o *List) SetName(name string) {

	o.Name = name
}

// Validate valides the current information stored into the structure.
func (o *List) Validate() error {

	errors := elemental.Errors{}
	requiredErrors := elemental.Errors{}

	if err := elemental.ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = append(requiredErrors, err)
	}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*List) SpecificationForAttribute(name string) elemental.AttributeSpecification {

	if v, ok := ListAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ListLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*List) AttributeSpecifications() map[string]elemental.AttributeSpecification {

	return ListAttributesMap
}

// ListAttributesMap represents the map of attribute for List.
var ListAttributesMap = map[string]elemental.AttributeSpecification{
	"ID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"CreationOnly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CreationOnly",
		CreationOnly:   true,
		Description:    `This attribute is creation only.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "creationOnly",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ParentType": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"ReadOnly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReadOnly",
		Description:    `This attribute is readonly.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"Slice": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Slice",
		Description:    `this is a slice.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "slice",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"Unexposed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}

// ListLowerCaseAttributesMap represents the map of attribute for List.
var ListLowerCaseAttributesMap = map[string]elemental.AttributeSpecification{
	"id": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ID",
		Description:    `The identifier.`,
		Exposed:        true,
		Filterable:     true,
		Identifier:     true,
		Name:           "ID",
		Orderable:      true,
		PrimaryKey:     true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"creationonly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "CreationOnly",
		CreationOnly:   true,
		Description:    `This attribute is creation only.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "creationOnly",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"date": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"description": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Name",
		Description:    `The name.`,
		Exposed:        true,
		Filterable:     true,
		Getter:         true,
		Name:           "name",
		Orderable:      true,
		Required:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentID",
		Description:    `The identifier of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		ForeignKey:     true,
		Name:           "parentID",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"parenttype": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		Autogenerated:  true,
		ConvertedName:  "ParentType",
		Description:    `The type of the parent of the object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "parentType",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"readonly": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "ReadOnly",
		Description:    `This attribute is readonly.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "readOnly",
		Orderable:      true,
		ReadOnly:       true,
		Stored:         true,
		Type:           "string",
	},
	"slice": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Slice",
		Description:    `this is a slice.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "slice",
		Orderable:      true,
		Stored:         true,
		SubType:        "string",
		Type:           "list",
	},
	"unexposed": elemental.AttributeSpecification{
		AllowedChoices: []string{},
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
}
