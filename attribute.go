// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

// AttributeSpecificationNameKey represents an attribute name.
type AttributeSpecificationNameKey string

// AttributeSpecifiable is the interface of an object that can access specification of its attributes.
type AttributeSpecifiable interface {
	SpecificationForAttribute(AttributeSpecificationNameKey) AttributeSpecification
}

// UniqueScope is a the type used to define uniqueness.
type UniqueScope int

const (
	// LocallyUnique represents the uniqueness in a particular context.
	LocallyUnique UniqueScope = iota + 1

	// GloballyUnique represents the absolute uniqueness.
	GloballyUnique
)

// AttributeSpecification represents all the metadata of an attribute.
type AttributeSpecification struct {
	AllowedChars   string
	AllowedChoices []string
	Autogenerated  bool
	Availability   string
	Channel        string
	CreationOnly   bool
	DefaultOrder   bool
	Deprecated     bool
	Exposed        bool
	Filterable     bool
	ForeignKey     bool
	Format         string
	Getter         bool
	Identifier     bool
	Index          bool
	MaxLength      uint
	MaxValue       float64
	MinLength      uint
	MinValue       float64
	Name           string
	Orderable      bool
	PrimaryKey     bool
	ReadOnly       bool
	Required       bool
	Setter         bool
	Stored         bool
	SubType        string
	Transient      bool
	Type           string
	Unique         bool
	UniqueScope    UniqueScope
}
