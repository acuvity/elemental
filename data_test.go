package elemental

import (
	"fmt"
	"slices"
	"time"

	"github.com/globalsign/mgo/bson"
	"github.com/mitchellh/copystructure"
)

//lint:file-ignore U1000 auto generated code.

// ListIdentity represents the Identity of the object.
var ListIdentity = Identity{
	Name:     "list",
	Category: "lists",
	Package:  "todo-list",
	Private:  false,
}

// ListsList represents a list of Lists
type ListsList []*List

// Identity returns the identity of the objects in the list.
func (o ListsList) Identity() Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the ListsList.
func (o ListsList) Copy() Identifiables {

	out := slices.Clone(o)
	return &out
}

// Append appends the objects to the a new copy of the ListsList.
func (o ListsList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*List))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o ListsList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o ListsList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the ListsList converted to SparseListsList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o ListsList) ToSparse(fields ...string) Identifiables {

	out := make(SparseListsList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToSparse(fields...).(*SparseList)
	}

	return out
}

// Version returns the version of the content.
func (o ListsList) Version() int {

	return 1
}

// List represents the model of a list
type List struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// This attribute is creation only.
	CreationOnly string `json:"creationOnly" msgpack:"creationOnly" bson:"creationonly" mapstructure:"creationOnly,omitempty"`

	// The date.
	Date time.Time `json:"date" msgpack:"date" bson:"date" mapstructure:"date,omitempty"`

	// The description.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The hash of the structure used to compare with new import version.
	ImportHash string `json:"importHash,omitempty" msgpack:"importHash,omitempty" bson:"importhash,omitempty" mapstructure:"importHash,omitempty"`

	// The user-defined import label that allows the system to group resources from the
	// same import operation.
	ImportLabel string `json:"importLabel,omitempty" msgpack:"importLabel,omitempty" bson:"importlabel,omitempty" mapstructure:"importLabel,omitempty"`

	// The name.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The namespace of the object.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// This attribute is readonly.
	ReadOnly string `json:"readOnly" msgpack:"readOnly" bson:"readonly" mapstructure:"readOnly,omitempty"`

	// This attribute is a ref to a single object.
	Ref *Task `json:"ref" msgpack:"ref" bson:"ref" mapstructure:"ref,omitempty"`

	// This attribute is a ref to a objects.
	RefList TasksList `json:"refList" msgpack:"refList" bson:"reflist" mapstructure:"refList,omitempty"`

	// This attribute is a ref map to a objects.
	RefMap map[string]*Task `json:"refMap" msgpack:"refMap" bson:"refmap" mapstructure:"refMap,omitempty"`

	// This attribute is secret.
	Secret string `json:"secret" msgpack:"secret" bson:"secret" mapstructure:"secret,omitempty"`

	// this is a slice.
	Slice []string `json:"slice" msgpack:"slice" bson:"slice" mapstructure:"slice,omitempty"`

	// This attribute is not exposed.
	Unexposed string `json:"-" msgpack:"-" bson:"unexposed" mapstructure:"-,omitempty"`

	// Hash of the object used to shard the data.
	ZHash int `json:"-" msgpack:"-" bson:"zhash" mapstructure:"-,omitempty"`

	// Sharding zone.
	Zone int `json:"-" msgpack:"-" bson:"zone" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewList returns a new *List
func NewList() *List {

	return &List{
		ModelVersion: 1,
		Ref:          NewTask(),
		RefList:      TasksList{},
		RefMap:       map[string]*Task{},
		Slice:        []string{},
	}
}

// Identity returns the Identity of the object.
func (o *List) Identity() Identity {

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

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *List) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesList{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.CreationOnly = o.CreationOnly
	s.Date = o.Date
	s.Description = o.Description
	s.ImportHash = o.ImportHash
	s.ImportLabel = o.ImportLabel
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType
	s.ReadOnly = o.ReadOnly
	s.Ref = o.Ref
	s.RefList = o.RefList
	s.RefMap = o.RefMap
	s.Secret = o.Secret
	s.Slice = o.Slice
	s.Unexposed = o.Unexposed
	s.ZHash = o.ZHash
	s.Zone = o.Zone

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *List) SetBSON(raw bson.Raw) error {

	if o == nil || raw.Kind == bson.ElementNil {
		return bson.ErrSetZero
	}

	s := &mongoAttributesList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.CreationOnly = s.CreationOnly
	o.Date = s.Date
	o.Description = s.Description
	o.ImportHash = s.ImportHash
	o.ImportLabel = s.ImportLabel
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType
	o.ReadOnly = s.ReadOnly
	o.Ref = s.Ref
	o.RefList = s.RefList
	o.RefMap = s.RefMap
	o.Secret = s.Secret
	o.Slice = s.Slice
	o.Unexposed = s.Unexposed
	o.ZHash = s.ZHash
	o.Zone = s.Zone

	return nil
}

// Version returns the hardcoded version of the model.
func (o *List) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *List) BleveType() string {

	return "list"
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

// GetImportHash returns the ImportHash of the receiver.
func (o *List) GetImportHash() string {

	return o.ImportHash
}

// SetImportHash sets the property ImportHash of the receiver using the given value.
func (o *List) SetImportHash(importHash string) {

	o.ImportHash = importHash
}

// GetImportLabel returns the ImportLabel of the receiver.
func (o *List) GetImportLabel() string {

	return o.ImportLabel
}

// SetImportLabel sets the property ImportLabel of the receiver using the given value.
func (o *List) SetImportLabel(importLabel string) {

	o.ImportLabel = importLabel
}

// GetName returns the Name of the receiver.
func (o *List) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *List) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *List) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *List) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *List) GetZHash() int {

	return o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the given value.
func (o *List) SetZHash(zHash int) {

	o.ZHash = zHash
}

// GetZone returns the Zone of the receiver.
func (o *List) GetZone() int {

	return o.Zone
}

// SetZone sets the property Zone of the receiver using the given value.
func (o *List) SetZone(zone int) {

	o.Zone = zone
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *List) ToSparse(fields ...string) SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseList{
			ID:           &o.ID,
			CreationOnly: &o.CreationOnly,
			Date:         &o.Date,
			Description:  &o.Description,
			ImportHash:   &o.ImportHash,
			ImportLabel:  &o.ImportLabel,
			Name:         &o.Name,
			Namespace:    &o.Namespace,
			ParentID:     &o.ParentID,
			ParentType:   &o.ParentType,
			ReadOnly:     &o.ReadOnly,
			Ref:          o.Ref,
			RefList:      &o.RefList,
			RefMap:       &o.RefMap,
			Secret:       &o.Secret,
			Slice:        &o.Slice,
			Unexposed:    &o.Unexposed,
			ZHash:        &o.ZHash,
			Zone:         &o.Zone,
		}
	}

	sp := &SparseList{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "creationOnly":
			sp.CreationOnly = &(o.CreationOnly)
		case "date":
			sp.Date = &(o.Date)
		case "description":
			sp.Description = &(o.Description)
		case "importHash":
			sp.ImportHash = &(o.ImportHash)
		case "importLabel":
			sp.ImportLabel = &(o.ImportLabel)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		case "readOnly":
			sp.ReadOnly = &(o.ReadOnly)
		case "ref":
			sp.Ref = o.Ref
		case "refList":
			sp.RefList = &(o.RefList)
		case "refMap":
			sp.RefMap = &(o.RefMap)
		case "secret":
			sp.Secret = &(o.Secret)
		case "slice":
			sp.Slice = &(o.Slice)
		case "unexposed":
			sp.Unexposed = &(o.Unexposed)
		case "zHash":
			sp.ZHash = &(o.ZHash)
		case "zone":
			sp.Zone = &(o.Zone)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseList to the object.
func (o *List) Patch(sparse SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseList)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.CreationOnly != nil {
		o.CreationOnly = *so.CreationOnly
	}
	if so.Date != nil {
		o.Date = *so.Date
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.ImportHash != nil {
		o.ImportHash = *so.ImportHash
	}
	if so.ImportLabel != nil {
		o.ImportLabel = *so.ImportLabel
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
	if so.ReadOnly != nil {
		o.ReadOnly = *so.ReadOnly
	}
	if so.Ref != nil {
		o.Ref = so.Ref
	}
	if so.RefList != nil {
		o.RefList = *so.RefList
	}
	if so.RefMap != nil {
		o.RefMap = *so.RefMap
	}
	if so.Secret != nil {
		o.Secret = *so.Secret
	}
	if so.Slice != nil {
		o.Slice = *so.Slice
	}
	if so.Unexposed != nil {
		o.Unexposed = *so.Unexposed
	}
	if so.ZHash != nil {
		o.ZHash = *so.ZHash
	}
	if so.Zone != nil {
		o.Zone = *so.Zone
	}
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *List) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Ref != nil {
		if err := o.Ref.EncryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to encrypt ref attribute 'Ref' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	for _, sub := range o.RefList {
		if sub == nil {
			continue
		}
		if err := sub.EncryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to encrypt refList/refMap attribute 'RefList' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	for _, sub := range o.RefMap {
		if sub == nil {
			continue
		}
		if err := sub.EncryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to encrypt refList/refMap attribute 'RefMap' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *List) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Ref != nil {
		if err := o.Ref.DecryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to decrypt ref attribute 'Ref' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	for _, sub := range o.RefList {
		if sub == nil {
			continue
		}
		if err := sub.DecryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to decrypt refList/refMap attribute 'RefList' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	for _, sub := range o.RefMap {
		if sub == nil {
			continue
		}
		if err := sub.DecryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to decrypt refList/refMap attribute 'RefMap' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	return nil
}

// DeepCopy returns a deep copy if the List.
func (o *List) DeepCopy() *List {

	if o == nil {
		return nil
	}

	out := &List{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *List.
func (o *List) DeepCopyInto(out *List) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy List: %s", err))
	}

	*out = *target.(*List)
}

// Validate valides the current information stored into the structure.
func (o *List) Validate() error {

	ResetDefaultForZeroValues(o)

	errors := Errors{}
	requiredErrors := Errors{}

	if err := ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if o.Ref != nil {
		if err := o.Ref.Validate(); err != nil {
			errors = errors.Append(err)
			InjectAttributePath(errors, "ref")
		}
	}

	for i, sub := range o.RefList {
		if sub == nil {
			continue
		}
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
			InjectAttributePath(errors, fmt.Sprintf("%s/%v", "refList", i))
		}
	}

	for i, sub := range o.RefMap {
		if sub == nil {
			continue
		}
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
			InjectAttributePath(errors, fmt.Sprintf("%s/%v", "refMap", i))
		}
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
func (*List) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := ListAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return ListLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*List) AttributeSpecifications() map[string]AttributeSpecification {

	return ListAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *List) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "creationOnly":
		return o.CreationOnly
	case "date":
		return o.Date
	case "description":
		return o.Description
	case "importHash":
		return o.ImportHash
	case "importLabel":
		return o.ImportLabel
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	case "readOnly":
		return o.ReadOnly
	case "ref":
		return o.Ref
	case "refList":
		return o.RefList
	case "refMap":
		return o.RefMap
	case "secret":
		return o.Secret
	case "slice":
		return o.Slice
	case "unexposed":
		return o.Unexposed
	case "zHash":
		return o.ZHash
	case "zone":
		return o.Zone
	}

	return nil
}

// ListAttributesMap represents the map of attribute for List.
var ListAttributesMap = map[string]AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"CreationOnly": {
		AllowedChoices: []string{},
		BSONFieldName:  "creationonly",
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
	"Date": {
		AllowedChoices: []string{},
		BSONFieldName:  "date",
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ImportHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "importhash",
		ConvertedName:  "ImportHash",
		CreationOnly:   true,
		Description:    `The hash of the structure used to compare with new import version.`,
		Exposed:        true,
		Getter:         true,
		Name:           "importHash",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ImportLabel": {
		AllowedChoices: []string{},
		BSONFieldName:  "importlabel",
		ConvertedName:  "ImportLabel",
		CreationOnly:   true,
		Description: `The user-defined import label that allows the system to group resources from the
same import operation.`,
		Exposed: true,
		Getter:  true,
		Name:    "importLabel",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
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
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
	"ReadOnly": {
		AllowedChoices: []string{},
		BSONFieldName:  "readonly",
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
	"Ref": {
		AllowedChoices: []string{},
		BSONFieldName:  "ref",
		ConvertedName:  "Ref",
		Description:    `This attribute is a ref to a single object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ref",
		Orderable:      true,
		Stored:         true,
		SubType:        "task",
		Type:           "ref",
	},
	"RefList": {
		AllowedChoices: []string{},
		BSONFieldName:  "reflist",
		ConvertedName:  "RefList",
		Description:    `This attribute is a ref to a objects.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "refList",
		Orderable:      true,
		Stored:         true,
		SubType:        "task",
		Type:           "refList",
	},
	"RefMap": {
		AllowedChoices: []string{},
		BSONFieldName:  "refmap",
		ConvertedName:  "RefMap",
		Description:    `This attribute is a ref map to a objects.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "refMap",
		Orderable:      true,
		Stored:         true,
		SubType:        "task",
		Type:           "refMap",
	},
	"Secret": {
		AllowedChoices: []string{},
		BSONFieldName:  "secret",
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"Slice": {
		AllowedChoices: []string{},
		BSONFieldName:  "slice",
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
	"Unexposed": {
		AllowedChoices: []string{},
		BSONFieldName:  "unexposed",
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"ZHash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description:    `Hash of the object used to shard the data.`,
		Getter:         true,
		Name:           "zHash",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"Zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Sharding zone.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// ListLowerCaseAttributesMap represents the map of attribute for List.
var ListLowerCaseAttributesMap = map[string]AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"creationonly": {
		AllowedChoices: []string{},
		BSONFieldName:  "creationonly",
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
	"date": {
		AllowedChoices: []string{},
		BSONFieldName:  "date",
		ConvertedName:  "Date",
		Description:    `The date.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "date",
		Orderable:      true,
		Stored:         true,
		Type:           "time",
	},
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"importhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "importhash",
		ConvertedName:  "ImportHash",
		CreationOnly:   true,
		Description:    `The hash of the structure used to compare with new import version.`,
		Exposed:        true,
		Getter:         true,
		Name:           "importHash",
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"importlabel": {
		AllowedChoices: []string{},
		BSONFieldName:  "importlabel",
		ConvertedName:  "ImportLabel",
		CreationOnly:   true,
		Description: `The user-defined import label that allows the system to group resources from the
same import operation.`,
		Exposed: true,
		Getter:  true,
		Name:    "importLabel",
		Setter:  true,
		Stored:  true,
		Type:    "string",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
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
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
	"readonly": {
		AllowedChoices: []string{},
		BSONFieldName:  "readonly",
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
	"ref": {
		AllowedChoices: []string{},
		BSONFieldName:  "ref",
		ConvertedName:  "Ref",
		Description:    `This attribute is a ref to a single object.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "ref",
		Orderable:      true,
		Stored:         true,
		SubType:        "task",
		Type:           "ref",
	},
	"reflist": {
		AllowedChoices: []string{},
		BSONFieldName:  "reflist",
		ConvertedName:  "RefList",
		Description:    `This attribute is a ref to a objects.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "refList",
		Orderable:      true,
		Stored:         true,
		SubType:        "task",
		Type:           "refList",
	},
	"refmap": {
		AllowedChoices: []string{},
		BSONFieldName:  "refmap",
		ConvertedName:  "RefMap",
		Description:    `This attribute is a ref map to a objects.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "refMap",
		Orderable:      true,
		Stored:         true,
		SubType:        "task",
		Type:           "refMap",
	},
	"secret": {
		AllowedChoices: []string{},
		BSONFieldName:  "secret",
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"slice": {
		AllowedChoices: []string{},
		BSONFieldName:  "slice",
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
	"unexposed": {
		AllowedChoices: []string{},
		BSONFieldName:  "unexposed",
		ConvertedName:  "Unexposed",
		Description:    `This attribute is not exposed.`,
		Filterable:     true,
		Name:           "unexposed",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"zhash": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zhash",
		ConvertedName:  "ZHash",
		Description:    `Hash of the object used to shard the data.`,
		Getter:         true,
		Name:           "zHash",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "integer",
	},
	"zone": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "zone",
		ConvertedName:  "Zone",
		Description:    `Sharding zone.`,
		Getter:         true,
		Name:           "zone",
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Transient:      true,
		Type:           "integer",
	},
}

// SparseListsList represents a list of SparseLists
type SparseListsList []*SparseList

// Identity returns the identity of the objects in the list.
func (o SparseListsList) Identity() Identity {

	return ListIdentity
}

// Copy returns a pointer to a copy the SparseListsList.
func (o SparseListsList) Copy() Identifiables {

	copy := slices.Clone(o)
	return &copy
}

// Append appends the objects to the a new copy of the SparseListsList.
func (o SparseListsList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*SparseList))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SparseListsList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseListsList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseListsList converted to ListsList.
func (o SparseListsList) ToPlain() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseListsList) Version() int {

	return 1
}

// SparseList represents the sparse version of a list.
type SparseList struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// This attribute is creation only.
	CreationOnly *string `json:"creationOnly,omitempty" msgpack:"creationOnly,omitempty" bson:"creationonly,omitempty" mapstructure:"creationOnly,omitempty"`

	// The date.
	Date *time.Time `json:"date,omitempty" msgpack:"date,omitempty" bson:"date,omitempty" mapstructure:"date,omitempty"`

	// The description.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The hash of the structure used to compare with new import version.
	ImportHash *string `json:"importHash,omitempty" msgpack:"importHash,omitempty" bson:"importhash,omitempty" mapstructure:"importHash,omitempty"`

	// The user-defined import label that allows the system to group resources from the
	// same import operation.
	ImportLabel *string `json:"importLabel,omitempty" msgpack:"importLabel,omitempty" bson:"importlabel,omitempty" mapstructure:"importLabel,omitempty"`

	// The name.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// The namespace of the object.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	// This attribute is readonly.
	ReadOnly *string `json:"readOnly,omitempty" msgpack:"readOnly,omitempty" bson:"readonly,omitempty" mapstructure:"readOnly,omitempty"`

	// This attribute is a ref to a single object.
	Ref *Task `json:"ref,omitempty" msgpack:"ref,omitempty" bson:"ref,omitempty" mapstructure:"ref,omitempty"`

	// This attribute is a ref to a objects.
	RefList *TasksList `json:"refList,omitempty" msgpack:"refList,omitempty" bson:"reflist,omitempty" mapstructure:"refList,omitempty"`

	// This attribute is a ref map to a objects.
	RefMap *map[string]*Task `json:"refMap,omitempty" msgpack:"refMap,omitempty" bson:"refmap,omitempty" mapstructure:"refMap,omitempty"`

	// This attribute is secret.
	Secret *string `json:"secret,omitempty" msgpack:"secret,omitempty" bson:"secret,omitempty" mapstructure:"secret,omitempty"`

	// this is a slice.
	Slice *[]string `json:"slice,omitempty" msgpack:"slice,omitempty" bson:"slice,omitempty" mapstructure:"slice,omitempty"`

	// This attribute is not exposed.
	Unexposed *string `json:"-" msgpack:"-" bson:"unexposed,omitempty" mapstructure:"-,omitempty"`

	// Hash of the object used to shard the data.
	ZHash *int `json:"-" msgpack:"-" bson:"zhash,omitempty" mapstructure:"-,omitempty"`

	// Sharding zone.
	Zone *int `json:"-" msgpack:"-" bson:"zone,omitempty" mapstructure:"-,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseList returns a new  SparseList.
func NewSparseList() *SparseList {
	return &SparseList{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseList) Identity() Identity {

	return ListIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseList) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseList) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseList) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseList{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.CreationOnly != nil {
		s.CreationOnly = o.CreationOnly
	}
	if o.Date != nil {
		s.Date = o.Date
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.ImportHash != nil {
		s.ImportHash = o.ImportHash
	}
	if o.ImportLabel != nil {
		s.ImportLabel = o.ImportLabel
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}
	if o.ReadOnly != nil {
		s.ReadOnly = o.ReadOnly
	}
	if o.Ref != nil {
		s.Ref = o.Ref
	}
	if o.RefList != nil {
		s.RefList = o.RefList
	}
	if o.RefMap != nil {
		s.RefMap = o.RefMap
	}
	if o.Secret != nil {
		s.Secret = o.Secret
	}
	if o.Slice != nil {
		s.Slice = o.Slice
	}
	if o.Unexposed != nil {
		s.Unexposed = o.Unexposed
	}
	if o.ZHash != nil {
		s.ZHash = o.ZHash
	}
	if o.Zone != nil {
		s.Zone = o.Zone
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseList) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseList{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.CreationOnly != nil {
		o.CreationOnly = s.CreationOnly
	}
	if s.Date != nil {
		o.Date = s.Date
	}
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.ImportHash != nil {
		o.ImportHash = s.ImportHash
	}
	if s.ImportLabel != nil {
		o.ImportLabel = s.ImportLabel
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}
	if s.ReadOnly != nil {
		o.ReadOnly = s.ReadOnly
	}
	if s.Ref != nil {
		o.Ref = s.Ref
	}
	if s.RefList != nil {
		o.RefList = s.RefList
	}
	if s.RefMap != nil {
		o.RefMap = s.RefMap
	}
	if s.Secret != nil {
		o.Secret = s.Secret
	}
	if s.Slice != nil {
		o.Slice = s.Slice
	}
	if s.Unexposed != nil {
		o.Unexposed = s.Unexposed
	}
	if s.ZHash != nil {
		o.ZHash = s.ZHash
	}
	if s.Zone != nil {
		o.Zone = s.Zone
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseList) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseList) ToPlain() PlainIdentifiable {

	out := NewList()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.CreationOnly != nil {
		out.CreationOnly = *o.CreationOnly
	}
	if o.Date != nil {
		out.Date = *o.Date
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.ImportHash != nil {
		out.ImportHash = *o.ImportHash
	}
	if o.ImportLabel != nil {
		out.ImportLabel = *o.ImportLabel
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}
	if o.ReadOnly != nil {
		out.ReadOnly = *o.ReadOnly
	}
	if o.Ref != nil {
		out.Ref = o.Ref
	}
	if o.RefList != nil {
		out.RefList = *o.RefList
	}
	if o.RefMap != nil {
		out.RefMap = *o.RefMap
	}
	if o.Secret != nil {
		out.Secret = *o.Secret
	}
	if o.Slice != nil {
		out.Slice = *o.Slice
	}
	if o.Unexposed != nil {
		out.Unexposed = *o.Unexposed
	}
	if o.ZHash != nil {
		out.ZHash = *o.ZHash
	}
	if o.Zone != nil {
		out.Zone = *o.Zone
	}

	return out
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseList) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Ref != nil {
		if err := o.Ref.EncryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to encrypt ref attribute 'Ref' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	if o.RefList != nil {
		for _, sub := range *o.RefList {
			if sub == nil {
				continue
			}
			if err := sub.EncryptAttributes(encrypter); err != nil {
				return fmt.Errorf("unable to encrypt refList/refMap attribute 'RefList' for 'List' (%s): %w", o.Identifier(), err)
			}
		}
	}

	if o.RefMap != nil {
		for _, sub := range *o.RefMap {
			if sub == nil {
				continue
			}
			if err := sub.EncryptAttributes(encrypter); err != nil {
				return fmt.Errorf("unable to encrypt refList/refMap attribute 'RefMap' for 'List' (%s): %w", o.Identifier(), err)
			}
		}
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseList) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Ref != nil {
		if err := o.Ref.DecryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to decrypt ref attribute 'Ref' for 'List' (%s): %w", o.Identifier(), err)
		}
	}

	if o.RefList != nil {
		for _, sub := range *o.RefList {
			if sub == nil {
				continue
			}
			if err := sub.DecryptAttributes(encrypter); err != nil {
				return fmt.Errorf("unable to decrypt refList/refMap attribute 'RefList' for 'List' (%s): %w", o.Identifier(), err)
			}
		}
	}

	if o.RefMap != nil {
		for _, sub := range *o.RefMap {
			if sub == nil {
				continue
			}
			if err := sub.DecryptAttributes(encrypter); err != nil {
				return fmt.Errorf("unable to decrypt refList/refMap attribute 'RefMap' for 'List' (%s): %w", o.Identifier(), err)
			}
		}
	}

	return nil
}

// GetImportHash returns the ImportHash of the receiver.
func (o *SparseList) GetImportHash() (out string) {

	if o.ImportHash == nil {
		return
	}

	return *o.ImportHash
}

// SetImportHash sets the property ImportHash of the receiver using the address of the given value.
func (o *SparseList) SetImportHash(importHash string) {

	o.ImportHash = &importHash
}

// GetImportLabel returns the ImportLabel of the receiver.
func (o *SparseList) GetImportLabel() (out string) {

	if o.ImportLabel == nil {
		return
	}

	return *o.ImportLabel
}

// SetImportLabel sets the property ImportLabel of the receiver using the address of the given value.
func (o *SparseList) SetImportLabel(importLabel string) {

	o.ImportLabel = &importLabel
}

// GetName returns the Name of the receiver.
func (o *SparseList) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseList) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseList) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseList) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// GetZHash returns the ZHash of the receiver.
func (o *SparseList) GetZHash() (out int) {

	if o.ZHash == nil {
		return
	}

	return *o.ZHash
}

// SetZHash sets the property ZHash of the receiver using the address of the given value.
func (o *SparseList) SetZHash(zHash int) {

	o.ZHash = &zHash
}

// GetZone returns the Zone of the receiver.
func (o *SparseList) GetZone() (out int) {

	if o.Zone == nil {
		return
	}

	return *o.Zone
}

// SetZone sets the property Zone of the receiver using the address of the given value.
func (o *SparseList) SetZone(zone int) {

	o.Zone = &zone
}

// DeepCopy returns a deep copy if the SparseList.
func (o *SparseList) DeepCopy() *SparseList {

	if o == nil {
		return nil
	}

	out := &SparseList{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseList.
func (o *SparseList) DeepCopyInto(out *SparseList) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseList: %s", err))
	}

	*out = *target.(*SparseList)
}

type mongoAttributesList struct {
	ID           bson.ObjectId    `bson:"_id,omitempty"`
	CreationOnly string           `bson:"creationonly"`
	Date         time.Time        `bson:"date"`
	Description  string           `bson:"description"`
	ImportHash   string           `bson:"importhash,omitempty"`
	ImportLabel  string           `bson:"importlabel,omitempty"`
	Name         string           `bson:"name"`
	Namespace    string           `bson:"namespace,omitempty"`
	ParentID     string           `bson:"parentid"`
	ParentType   string           `bson:"parenttype"`
	ReadOnly     string           `bson:"readonly"`
	Ref          *Task            `bson:"ref"`
	RefList      TasksList        `bson:"reflist"`
	RefMap       map[string]*Task `bson:"refmap"`
	Secret       string           `bson:"secret"`
	Slice        []string         `bson:"slice"`
	Unexposed    string           `bson:"unexposed"`
	ZHash        int              `bson:"zhash"`
	Zone         int              `bson:"zone"`
}
type mongoAttributesSparseList struct {
	ID           bson.ObjectId     `bson:"_id,omitempty"`
	CreationOnly *string           `bson:"creationonly,omitempty"`
	Date         *time.Time        `bson:"date,omitempty"`
	Description  *string           `bson:"description,omitempty"`
	ImportHash   *string           `bson:"importhash,omitempty"`
	ImportLabel  *string           `bson:"importlabel,omitempty"`
	Name         *string           `bson:"name,omitempty"`
	Namespace    *string           `bson:"namespace,omitempty"`
	ParentID     *string           `bson:"parentid,omitempty"`
	ParentType   *string           `bson:"parenttype,omitempty"`
	ReadOnly     *string           `bson:"readonly,omitempty"`
	Ref          *Task             `bson:"ref,omitempty"`
	RefList      *TasksList        `bson:"reflist,omitempty"`
	RefMap       *map[string]*Task `bson:"refmap,omitempty"`
	Secret       *string           `bson:"secret,omitempty"`
	Slice        *[]string         `bson:"slice,omitempty"`
	Unexposed    *string           `bson:"unexposed,omitempty"`
	ZHash        *int              `bson:"zhash,omitempty"`
	Zone         *int              `bson:"zone,omitempty"`
}

// SubtaskIdentity represents the Identity of the object.
var SubtaskIdentity = Identity{
	Name:     "subtask",
	Category: "subtasks",
	Package:  "todo-list",
	Private:  false,
}

// SubtasksList represents a list of Subtasks
type SubtasksList []*Subtask

// Identity returns the identity of the objects in the list.
func (o SubtasksList) Identity() Identity {

	return SubtaskIdentity
}

// Copy returns a pointer to a copy the SubtasksList.
func (o SubtasksList) Copy() Identifiables {

	out := slices.Clone(o)
	return &out
}

// Append appends the objects to the a new copy of the SubtasksList.
func (o SubtasksList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*Subtask))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SubtasksList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SubtasksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the SubtasksList converted to SparseSubtasksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o SubtasksList) ToSparse(fields ...string) Identifiables {

	out := make(SparseSubtasksList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToSparse(fields...).(*SparseSubtask)
	}

	return out
}

// Version returns the version of the content.
func (o SubtasksList) Version() int {

	return 1
}

// Subtask represents the model of a subtask
type Subtask struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The Name.
	Name string `json:"Name" msgpack:"Name" bson:"name" mapstructure:"Name,omitempty"`

	// The namespace of the object.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSubtask returns a new *Subtask
func NewSubtask() *Subtask {

	return &Subtask{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Subtask) Identity() Identity {

	return SubtaskIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Subtask) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Subtask) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Subtask) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSubtask{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Subtask) SetBSON(raw bson.Raw) error {

	if o == nil || raw.Kind == bson.ElementNil {
		return bson.ErrSetZero
	}

	s := &mongoAttributesSubtask{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Subtask) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Subtask) BleveType() string {

	return "subtask"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Subtask) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Subtask) Doc() string {

	return `Represent a subtask object.`
}

func (o *Subtask) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetNamespace returns the Namespace of the receiver.
func (o *Subtask) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Subtask) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Subtask) ToSparse(fields ...string) SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseSubtask{
			ID:         &o.ID,
			Name:       &o.Name,
			Namespace:  &o.Namespace,
			ParentID:   &o.ParentID,
			ParentType: &o.ParentType,
		}
	}

	sp := &SparseSubtask{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "Name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseSubtask to the object.
func (o *Subtask) Patch(sparse SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseSubtask)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *Subtask) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *Subtask) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DeepCopy returns a deep copy if the Subtask.
func (o *Subtask) DeepCopy() *Subtask {

	if o == nil {
		return nil
	}

	out := &Subtask{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Subtask.
func (o *Subtask) DeepCopyInto(out *Subtask) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Subtask: %s", err))
	}

	*out = *target.(*Subtask)
}

// Validate valides the current information stored into the structure.
func (o *Subtask) Validate() error {

	ResetDefaultForZeroValues(o)

	errors := Errors{}
	requiredErrors := Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Subtask) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := SubtaskAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return SubtaskLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Subtask) AttributeSpecifications() map[string]AttributeSpecification {

	return SubtaskAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Subtask) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "Name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	}

	return nil
}

// SubtaskAttributesMap represents the map of attribute for Subtask.
var SubtaskAttributesMap = map[string]AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `The Name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "Name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
}

// SubtaskLowerCaseAttributesMap represents the map of attribute for Subtask.
var SubtaskLowerCaseAttributesMap = map[string]AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
		ConvertedName:  "Name",
		Description:    `The Name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "Name",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
}

// SparseSubtasksList represents a list of SparseSubtasks
type SparseSubtasksList []*SparseSubtask

// Identity returns the identity of the objects in the list.
func (o SparseSubtasksList) Identity() Identity {

	return SubtaskIdentity
}

// Copy returns a pointer to a copy the SparseSubtasksList.
func (o SparseSubtasksList) Copy() Identifiables {

	copy := slices.Clone(o)
	return &copy
}

// Append appends the objects to the a new copy of the SparseSubtasksList.
func (o SparseSubtasksList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*SparseSubtask))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SparseSubtasksList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseSubtasksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseSubtasksList converted to SubtasksList.
func (o SparseSubtasksList) ToPlain() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseSubtasksList) Version() int {

	return 1
}

// SparseSubtask represents the sparse version of a subtask.
type SparseSubtask struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The Name.
	Name *string `json:"Name,omitempty" msgpack:"Name,omitempty" bson:"name,omitempty" mapstructure:"Name,omitempty"`

	// The namespace of the object.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseSubtask returns a new  SparseSubtask.
func NewSparseSubtask() *SparseSubtask {
	return &SparseSubtask{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseSubtask) Identity() Identity {

	return SubtaskIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseSubtask) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseSubtask) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSubtask) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseSubtask{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseSubtask) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseSubtask{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseSubtask) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseSubtask) ToPlain() PlainIdentifiable {

	out := NewSubtask()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}

	return out
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseSubtask) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseSubtask) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseSubtask) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseSubtask) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseSubtask.
func (o *SparseSubtask) DeepCopy() *SparseSubtask {

	if o == nil {
		return nil
	}

	out := &SparseSubtask{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseSubtask.
func (o *SparseSubtask) DeepCopyInto(out *SparseSubtask) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseSubtask: %s", err))
	}

	*out = *target.(*SparseSubtask)
}

type mongoAttributesSubtask struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       string        `bson:"name"`
	Namespace  string        `bson:"namespace,omitempty"`
	ParentID   string        `bson:"parentid"`
	ParentType string        `bson:"parenttype"`
}
type mongoAttributesSparseSubtask struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Name       *string       `bson:"name,omitempty"`
	Namespace  *string       `bson:"namespace,omitempty"`
	ParentID   *string       `bson:"parentid,omitempty"`
	ParentType *string       `bson:"parenttype,omitempty"`
}

// TaskStatusValue represents the possible values for attribute "status".
type TaskStatusValue string

const (
	// TaskStatusDONE represents the value DONE.
	TaskStatusDONE TaskStatusValue = "DONE"

	// TaskStatusPROGRESS represents the value PROGRESS.
	TaskStatusPROGRESS TaskStatusValue = "PROGRESS"

	// TaskStatusTODO represents the value TODO.
	TaskStatusTODO TaskStatusValue = "TODO"
)

// TaskIdentity represents the Identity of the object.
var TaskIdentity = Identity{
	Name:     "task",
	Category: "tasks",
	Package:  "todo-list",
	Private:  false,
}

// TasksList represents a list of Tasks
type TasksList []*Task

// Identity returns the identity of the objects in the list.
func (o TasksList) Identity() Identity {

	return TaskIdentity
}

// Copy returns a pointer to a copy the TasksList.
func (o TasksList) Copy() Identifiables {

	out := slices.Clone(o)
	return &out
}

// Append appends the objects to the a new copy of the TasksList.
func (o TasksList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*Task))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o TasksList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o TasksList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the TasksList converted to SparseTasksList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o TasksList) ToSparse(fields ...string) Identifiables {

	out := make(SparseTasksList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToSparse(fields...).(*SparseTask)
	}

	return out
}

// Version returns the version of the content.
func (o TasksList) Version() int {

	return 1
}

// Task represents the model of a task
type Task struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// The description.
	Description string `json:"description" msgpack:"description" bson:"description" mapstructure:"description,omitempty"`

	// The name.
	Name string `json:"name" msgpack:"name" bson:"name" mapstructure:"name,omitempty"`

	// The namespace of the object.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// This attribute is secret.
	Secret string `json:"secret" msgpack:"secret" bson:"secret" mapstructure:"secret,omitempty"`

	// The status of the task.
	Status TaskStatusValue `json:"status" msgpack:"status" bson:"status" mapstructure:"status,omitempty"`

	// This is a nested ref.
	Subtask *Subtask `json:"-" msgpack:"-" bson:"subtask" mapstructure:"-,omitempty"`

	// This is a nested refList.
	SubtaskList SubtasksList `json:"subtaskList" msgpack:"subtaskList" bson:"subtasklist" mapstructure:"subtaskList,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewTask returns a new *Task
func NewTask() *Task {

	return &Task{
		ModelVersion: 1,
		Status:       TaskStatusTODO,
		Subtask:      NewSubtask(),
	}
}

// Identity returns the Identity of the object.
func (o *Task) Identity() Identity {

	return TaskIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Task) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Task) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Task) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesTask{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Description = o.Description
	s.Name = o.Name
	s.Namespace = o.Namespace
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType
	s.Secret = o.Secret
	s.Status = o.Status
	s.Subtask = o.Subtask
	s.SubtaskList = o.SubtaskList

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Task) SetBSON(raw bson.Raw) error {

	if o == nil || raw.Kind == bson.ElementNil {
		return bson.ErrSetZero
	}

	s := &mongoAttributesTask{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Description = s.Description
	o.Name = s.Name
	o.Namespace = s.Namespace
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType
	o.Secret = s.Secret
	o.Status = s.Status
	o.Subtask = s.Subtask
	o.SubtaskList = s.SubtaskList

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Task) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Task) BleveType() string {

	return "task"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Task) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Task) Doc() string {

	return `Represent a task to do in a listd.`
}

func (o *Task) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetName returns the Name of the receiver.
func (o *Task) GetName() string {

	return o.Name
}

// SetName sets the property Name of the receiver using the given value.
func (o *Task) SetName(name string) {

	o.Name = name
}

// GetNamespace returns the Namespace of the receiver.
func (o *Task) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *Task) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *Task) ToSparse(fields ...string) SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseTask{
			ID:          &o.ID,
			Description: &o.Description,
			Name:        &o.Name,
			Namespace:   &o.Namespace,
			ParentID:    &o.ParentID,
			ParentType:  &o.ParentType,
			Secret:      &o.Secret,
			Status:      &o.Status,
			Subtask:     o.Subtask,
			SubtaskList: &o.SubtaskList,
		}
	}

	sp := &SparseTask{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "description":
			sp.Description = &(o.Description)
		case "name":
			sp.Name = &(o.Name)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		case "secret":
			sp.Secret = &(o.Secret)
		case "status":
			sp.Status = &(o.Status)
		case "subtask":
			sp.Subtask = o.Subtask
		case "subtaskList":
			sp.SubtaskList = &(o.SubtaskList)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseTask to the object.
func (o *Task) Patch(sparse SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseTask)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Description != nil {
		o.Description = *so.Description
	}
	if so.Name != nil {
		o.Name = *so.Name
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
	if so.Secret != nil {
		o.Secret = *so.Secret
	}
	if so.Status != nil {
		o.Status = *so.Status
	}
	if so.Subtask != nil {
		o.Subtask = so.Subtask
	}
	if so.SubtaskList != nil {
		o.SubtaskList = *so.SubtaskList
	}
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *Task) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Subtask != nil {
		if err := o.Subtask.EncryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to encrypt ref attribute 'Subtask' for 'Task' (%s): %w", o.Identifier(), err)
		}
	}

	for _, sub := range o.SubtaskList {
		if err := sub.EncryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to encrypt refList/refMap attribute 'SubtaskList' for 'Task' (%s): %w", o.Identifier(), err)
		}
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *Task) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Subtask != nil {
		if err := o.Subtask.DecryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to decrypt ref attribute 'Subtask' for 'Task' (%s): %w", o.Identifier(), err)
		}
	}

	for _, sub := range o.SubtaskList {
		if err := sub.DecryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to decrypt refList/refMap attribute 'SubtaskList' for 'Task' (%s): %w", o.Identifier(), err)
		}
	}

	return nil
}

// DeepCopy returns a deep copy if the Task.
func (o *Task) DeepCopy() *Task {

	if o == nil {
		return nil
	}

	out := &Task{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Task.
func (o *Task) DeepCopyInto(out *Task) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Task: %s", err))
	}

	*out = *target.(*Task)
}

// Validate valides the current information stored into the structure.
func (o *Task) Validate() error {

	ResetDefaultForZeroValues(o)

	errors := Errors{}
	requiredErrors := Errors{}

	if err := ValidateRequiredString("name", o.Name); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateStringInList("status", string(o.Status), []string{"DONE", "PROGRESS", "TODO"}, false); err != nil {
		errors = errors.Append(err)
	}

	for i, sub := range o.SubtaskList {
		if err := sub.Validate(); err != nil {
			errors = errors.Append(err)
			InjectAttributePath(errors, fmt.Sprintf("%s/%v", "subtaskList", i))
		}
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
func (*Task) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := TaskAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return TaskLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Task) AttributeSpecifications() map[string]AttributeSpecification {

	return TaskAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Task) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "description":
		return o.Description
	case "name":
		return o.Name
	case "namespace":
		return o.Namespace
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	case "secret":
		return o.Secret
	case "status":
		return o.Status
	case "subtask":
		return o.Subtask
	case "subtaskList":
		return o.SubtaskList
	}

	return nil
}

// TaskAttributesMap represents the map of attribute for Task.
var TaskAttributesMap = map[string]AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"Description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"Name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
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
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
	"Secret": {
		AllowedChoices: []string{},
		BSONFieldName:  "secret",
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"Status": {
		AllowedChoices: []string{"DONE", "PROGRESS", "TODO"},
		BSONFieldName:  "status",
		ConvertedName:  "Status",
		DefaultValue:   TaskStatusTODO,
		Description:    `The status of the task.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"Subtask": {
		AllowedChoices: []string{},
		BSONFieldName:  "subtask",
		ConvertedName:  "Subtask",
		Description:    `This is a nested ref.`,
		Name:           "subtask",
		Stored:         true,
		SubType:        "subtask",
		Type:           "ref",
	},
	"SubtaskList": {
		AllowedChoices: []string{},
		BSONFieldName:  "subtasklist",
		ConvertedName:  "SubtaskList",
		Description:    `This is a nested refList.`,
		Exposed:        true,
		Name:           "subtaskList",
		Stored:         true,
		SubType:        "subtask",
		Type:           "refList",
	},
}

// TaskLowerCaseAttributesMap represents the map of attribute for Task.
var TaskLowerCaseAttributesMap = map[string]AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"description": {
		AllowedChoices: []string{},
		BSONFieldName:  "description",
		ConvertedName:  "Description",
		Description:    `The description.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "description",
		Orderable:      true,
		Stored:         true,
		Type:           "string",
	},
	"name": {
		AllowedChoices: []string{},
		BSONFieldName:  "name",
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
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
	"secret": {
		AllowedChoices: []string{},
		BSONFieldName:  "secret",
		ConvertedName:  "Secret",
		Description:    `This attribute is secret.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "secret",
		Orderable:      true,
		Secret:         true,
		Stored:         true,
		Type:           "string",
	},
	"status": {
		AllowedChoices: []string{"DONE", "PROGRESS", "TODO"},
		BSONFieldName:  "status",
		ConvertedName:  "Status",
		DefaultValue:   TaskStatusTODO,
		Description:    `The status of the task.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "status",
		Orderable:      true,
		Stored:         true,
		Type:           "enum",
	},
	"subtask": {
		AllowedChoices: []string{},
		BSONFieldName:  "subtask",
		ConvertedName:  "Subtask",
		Description:    `This is a nested ref.`,
		Name:           "subtask",
		Stored:         true,
		SubType:        "subtask",
		Type:           "ref",
	},
	"subtasklist": {
		AllowedChoices: []string{},
		BSONFieldName:  "subtasklist",
		ConvertedName:  "SubtaskList",
		Description:    `This is a nested refList.`,
		Exposed:        true,
		Name:           "subtaskList",
		Stored:         true,
		SubType:        "subtask",
		Type:           "refList",
	},
}

// SparseTasksList represents a list of SparseTasks
type SparseTasksList []*SparseTask

// Identity returns the identity of the objects in the list.
func (o SparseTasksList) Identity() Identity {

	return TaskIdentity
}

// Copy returns a pointer to a copy the SparseTasksList.
func (o SparseTasksList) Copy() Identifiables {

	copy := slices.Clone(o)
	return &copy
}

// Append appends the objects to the a new copy of the SparseTasksList.
func (o SparseTasksList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*SparseTask))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SparseTasksList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseTasksList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseTasksList converted to TasksList.
func (o SparseTasksList) ToPlain() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseTasksList) Version() int {

	return 1
}

// SparseTask represents the sparse version of a task.
type SparseTask struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// The description.
	Description *string `json:"description,omitempty" msgpack:"description,omitempty" bson:"description,omitempty" mapstructure:"description,omitempty"`

	// The name.
	Name *string `json:"name,omitempty" msgpack:"name,omitempty" bson:"name,omitempty" mapstructure:"name,omitempty"`

	// The namespace of the object.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	// This attribute is secret.
	Secret *string `json:"secret,omitempty" msgpack:"secret,omitempty" bson:"secret,omitempty" mapstructure:"secret,omitempty"`

	// The status of the task.
	Status *TaskStatusValue `json:"status,omitempty" msgpack:"status,omitempty" bson:"status,omitempty" mapstructure:"status,omitempty"`

	// This is a nested ref.
	Subtask *Subtask `json:"-" msgpack:"-" bson:"subtask,omitempty" mapstructure:"-,omitempty"`

	// This is a nested refList.
	SubtaskList *SubtasksList `json:"subtaskList,omitempty" msgpack:"subtaskList,omitempty" bson:"subtasklist,omitempty" mapstructure:"subtaskList,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseTask returns a new  SparseTask.
func NewSparseTask() *SparseTask {
	return &SparseTask{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseTask) Identity() Identity {

	return TaskIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseTask) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseTask) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTask) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseTask{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Description != nil {
		s.Description = o.Description
	}
	if o.Name != nil {
		s.Name = o.Name
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}
	if o.Secret != nil {
		s.Secret = o.Secret
	}
	if o.Status != nil {
		s.Status = o.Status
	}
	if o.Subtask != nil {
		s.Subtask = o.Subtask
	}
	if o.SubtaskList != nil {
		s.SubtaskList = o.SubtaskList
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseTask) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseTask{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Description != nil {
		o.Description = s.Description
	}
	if s.Name != nil {
		o.Name = s.Name
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}
	if s.Secret != nil {
		o.Secret = s.Secret
	}
	if s.Status != nil {
		o.Status = s.Status
	}
	if s.Subtask != nil {
		o.Subtask = s.Subtask
	}
	if s.SubtaskList != nil {
		o.SubtaskList = s.SubtaskList
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseTask) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseTask) ToPlain() PlainIdentifiable {

	out := NewTask()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Description != nil {
		out.Description = *o.Description
	}
	if o.Name != nil {
		out.Name = *o.Name
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}
	if o.Secret != nil {
		out.Secret = *o.Secret
	}
	if o.Status != nil {
		out.Status = *o.Status
	}
	if o.Subtask != nil {
		out.Subtask = o.Subtask
	}
	if o.SubtaskList != nil {
		out.SubtaskList = *o.SubtaskList
	}

	return out
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseTask) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Subtask != nil {
		if err := o.Subtask.EncryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to encrypt ref attribute 'Subtask' for 'Task' (%s): %w", o.Identifier(), err)
		}
	}

	if o.SubtaskList != nil {
		for _, sub := range *o.SubtaskList {
			if err := sub.EncryptAttributes(encrypter); err != nil {
				return fmt.Errorf("unable to encrypt refList/refMap attribute 'SubtaskList' for 'Task' (%s): %w", o.Identifier(), err)
			}
		}
	}

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseTask) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	if o.Subtask != nil {
		if err := o.Subtask.DecryptAttributes(encrypter); err != nil {
			return fmt.Errorf("unable to decrypt ref attribute 'Subtask' for 'Task' (%s): %w", o.Identifier(), err)
		}
	}

	if o.SubtaskList != nil {
		for _, sub := range *o.SubtaskList {
			if err := sub.DecryptAttributes(encrypter); err != nil {
				return fmt.Errorf("unable to decrypt refList/refMap attribute 'SubtaskList' for 'Task' (%s): %w", o.Identifier(), err)
			}
		}
	}

	return nil
}

// GetName returns the Name of the receiver.
func (o *SparseTask) GetName() (out string) {

	if o.Name == nil {
		return
	}

	return *o.Name
}

// SetName sets the property Name of the receiver using the address of the given value.
func (o *SparseTask) SetName(name string) {

	o.Name = &name
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseTask) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseTask) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseTask.
func (o *SparseTask) DeepCopy() *SparseTask {

	if o == nil {
		return nil
	}

	out := &SparseTask{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseTask.
func (o *SparseTask) DeepCopyInto(out *SparseTask) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseTask: %s", err))
	}

	*out = *target.(*SparseTask)
}

type mongoAttributesTask struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Description string          `bson:"description"`
	Name        string          `bson:"name"`
	Namespace   string          `bson:"namespace,omitempty"`
	ParentID    string          `bson:"parentid"`
	ParentType  string          `bson:"parenttype"`
	Secret      string          `bson:"secret"`
	Status      TaskStatusValue `bson:"status"`
	Subtask     *Subtask        `bson:"subtask"`
	SubtaskList SubtasksList    `bson:"subtasklist"`
}
type mongoAttributesSparseTask struct {
	ID          bson.ObjectId    `bson:"_id,omitempty"`
	Description *string          `bson:"description,omitempty"`
	Name        *string          `bson:"name,omitempty"`
	Namespace   *string          `bson:"namespace,omitempty"`
	ParentID    *string          `bson:"parentid,omitempty"`
	ParentType  *string          `bson:"parenttype,omitempty"`
	Secret      *string          `bson:"secret,omitempty"`
	Status      *TaskStatusValue `bson:"status,omitempty"`
	Subtask     *Subtask         `bson:"subtask,omitempty"`
	SubtaskList *SubtasksList    `bson:"subtasklist,omitempty"`
}

// UnmarshalableListIdentity represents the Identity of the object.
var UnmarshalableListIdentity = Identity{Name: "list", Category: "lists"}

// UnmarshalableListsList represents a list of UnmarshalableLists
type UnmarshalableListsList []*UnmarshalableList

// Identity returns the identity of the objects in the list.
func (o UnmarshalableListsList) Identity() Identity {

	return UnmarshalableListIdentity
}

// Copy returns a pointer to a copy the UnmarshalableListsList.
func (o UnmarshalableListsList) Copy() Identifiables {

	out := slices.Clone(o)
	return &out
}

// Append appends the objects to the a new copy of the UnmarshalableListsList.
func (o UnmarshalableListsList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*UnmarshalableList))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o UnmarshalableListsList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range o {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o UnmarshalableListsList) DefaultOrder() []string {

	return []string{
		"flagDefaultOrderingKey",
	}
}

// Version returns the version of the content.
func (o UnmarshalableListsList) Version() int {

	return 1
}

// An UnmarshalableList is a List that cannot be marshalled  or unmarshalled.
type UnmarshalableList struct {
	List
}

// NewUnmarshalableList returns a new UnmarshalableList.
func NewUnmarshalableList() *UnmarshalableList {
	return &UnmarshalableList{List: List{}}
}

// Identity returns the identity.
func (o *UnmarshalableList) Identity() Identity { return UnmarshalableListIdentity }

// UnmarshalJSON makes the UnmarshalableList not unmarshalable.
func (o *UnmarshalableList) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalJSON makes the UnmarshalableList not marshalable.
func (o *UnmarshalableList) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// UnmarshalMsgpack makes the UnmarshalableList not unmarshalable.
func (o *UnmarshalableList) UnmarshalMsgpack([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalMsgpack makes the UnmarshalableList not marshalable.
func (o *UnmarshalableList) MarshalMsgpack() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// Validate validates the data
func (o *UnmarshalableList) Validate() Errors { return nil }

// An UnmarshalableError is a List that cannot be marshalled or unmarshalled.
type UnmarshalableError struct {
	Error
}

// UnmarshalJSON makes the UnmarshalableError not unmarshalable.
func (o *UnmarshalableError) UnmarshalJSON([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalJSON makes the UnmarshalableError not marshalable.
func (o *UnmarshalableError) MarshalJSON() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// UnmarshalMsgpack makes the UnmarshalableError not unmarshalable.
func (o *UnmarshalableError) UnmarshalMsgpack([]byte) error {
	return fmt.Errorf("error unmarshalling")
}

// MarshalMsgpack makes the UnmarshalableError not marshalable.
func (o *UnmarshalableError) MarshalMsgpack() ([]byte, error) {
	return nil, fmt.Errorf("error marshalling")
}

// UserIdentity represents the Identity of the object.
var UserIdentity = Identity{
	Name:     "user",
	Category: "users",
	Package:  "todo-list",
	Private:  false,
}

// UsersList represents a list of Users
type UsersList []*User

// Identity returns the identity of the objects in the list.
func (o UsersList) Identity() Identity {

	return UserIdentity
}

// Copy returns a pointer to a copy the UsersList.
func (o UsersList) Copy() Identifiables {

	out := slices.Clone(o)
	return &out
}

// Append appends the objects to the a new copy of the UsersList.
func (o UsersList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*User))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o UsersList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o UsersList) DefaultOrder() []string {

	return []string{}
}

// ToSparse returns the UsersList converted to SparseUsersList.
// Objects in the list will only contain the given fields. No field means entire field set.
func (o UsersList) ToSparse(fields ...string) Identifiables {

	out := make(SparseUsersList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToSparse(fields...).(*SparseUser)
	}

	return out
}

// Version returns the version of the content.
func (o UsersList) Version() int {

	return 1
}

// User represents the model of a user
type User struct {
	// The identifier.
	ID string `json:"ID" msgpack:"ID" bson:"-" mapstructure:"ID,omitempty"`

	// the object is archived and not deleted.
	Archived bool `json:"archived" msgpack:"archived" bson:"archived" mapstructure:"archived,omitempty"`

	// The first name.
	FirstName string `json:"firstName" msgpack:"firstName" bson:"firstname" mapstructure:"firstName,omitempty"`

	// The last name.
	LastName string `json:"lastName" msgpack:"lastName" bson:"lastname" mapstructure:"lastName,omitempty"`

	// The namespace of the object.
	Namespace string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID string `json:"parentID" msgpack:"parentID" bson:"parentid" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType string `json:"parentType" msgpack:"parentType" bson:"parenttype" mapstructure:"parentType,omitempty"`

	// the login.
	UserName string `json:"userName" msgpack:"userName" bson:"username" mapstructure:"userName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewUser returns a new *User
func NewUser() *User {

	return &User{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *User) Identity() Identity {

	return UserIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *User) Identifier() string {

	return o.ID
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *User) SetIdentifier(id string) {

	o.ID = id
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *User) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesUser{}

	if o.ID != "" {
		s.ID = bson.ObjectIdHex(o.ID)
	}
	s.Archived = o.Archived
	s.FirstName = o.FirstName
	s.LastName = o.LastName
	s.Namespace = o.Namespace
	s.ParentID = o.ParentID
	s.ParentType = o.ParentType
	s.UserName = o.UserName

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *User) SetBSON(raw bson.Raw) error {

	if o == nil || raw.Kind == bson.ElementNil {
		return bson.ErrSetZero
	}

	s := &mongoAttributesUser{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	o.ID = s.ID.Hex()
	o.Archived = s.Archived
	o.FirstName = s.FirstName
	o.LastName = s.LastName
	o.Namespace = s.Namespace
	o.ParentID = s.ParentID
	o.ParentType = s.ParentType
	o.UserName = s.UserName

	return nil
}

// Version returns the hardcoded version of the model.
func (o *User) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *User) BleveType() string {

	return "user"
}

// DefaultOrder returns the list of default ordering fields.
func (o *User) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *User) Doc() string {

	return `Represent a user.`
}

func (o *User) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// GetArchived returns the Archived of the receiver.
func (o *User) GetArchived() bool {

	return o.Archived
}

// SetArchived sets the property Archived of the receiver using the given value.
func (o *User) SetArchived(archived bool) {

	o.Archived = archived
}

// GetNamespace returns the Namespace of the receiver.
func (o *User) GetNamespace() string {

	return o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the given value.
func (o *User) SetNamespace(namespace string) {

	o.Namespace = namespace
}

// ToSparse returns the sparse version of the model.
// The returned object will only contain the given fields. No field means entire field set.
func (o *User) ToSparse(fields ...string) SparseIdentifiable {

	if len(fields) == 0 {
		// nolint: goimports
		return &SparseUser{
			ID:         &o.ID,
			Archived:   &o.Archived,
			FirstName:  &o.FirstName,
			LastName:   &o.LastName,
			Namespace:  &o.Namespace,
			ParentID:   &o.ParentID,
			ParentType: &o.ParentType,
			UserName:   &o.UserName,
		}
	}

	sp := &SparseUser{}
	for _, f := range fields {
		switch f {
		case "ID":
			sp.ID = &(o.ID)
		case "archived":
			sp.Archived = &(o.Archived)
		case "firstName":
			sp.FirstName = &(o.FirstName)
		case "lastName":
			sp.LastName = &(o.LastName)
		case "namespace":
			sp.Namespace = &(o.Namespace)
		case "parentID":
			sp.ParentID = &(o.ParentID)
		case "parentType":
			sp.ParentType = &(o.ParentType)
		case "userName":
			sp.UserName = &(o.UserName)
		}
	}

	return sp
}

// Patch apply the non nil value of a *SparseUser to the object.
func (o *User) Patch(sparse SparseIdentifiable) {
	if !sparse.Identity().IsEqual(o.Identity()) {
		panic("cannot patch from a parse with different identity")
	}

	so := sparse.(*SparseUser)
	if so.ID != nil {
		o.ID = *so.ID
	}
	if so.Archived != nil {
		o.Archived = *so.Archived
	}
	if so.FirstName != nil {
		o.FirstName = *so.FirstName
	}
	if so.LastName != nil {
		o.LastName = *so.LastName
	}
	if so.Namespace != nil {
		o.Namespace = *so.Namespace
	}
	if so.ParentID != nil {
		o.ParentID = *so.ParentID
	}
	if so.ParentType != nil {
		o.ParentType = *so.ParentType
	}
	if so.UserName != nil {
		o.UserName = *so.UserName
	}
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *User) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *User) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DeepCopy returns a deep copy if the User.
func (o *User) DeepCopy() *User {

	if o == nil {
		return nil
	}

	out := &User{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *User.
func (o *User) DeepCopyInto(out *User) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy User: %s", err))
	}

	*out = *target.(*User)
}

// Validate valides the current information stored into the structure.
func (o *User) Validate() error {

	ResetDefaultForZeroValues(o)

	errors := Errors{}
	requiredErrors := Errors{}

	if err := ValidateRequiredString("firstName", o.FirstName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateRequiredString("lastName", o.LastName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	if err := ValidateRequiredString("userName", o.UserName); err != nil {
		requiredErrors = requiredErrors.Append(err)
	}

	// Custom object validation.

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*User) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := UserAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return UserLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*User) AttributeSpecifications() map[string]AttributeSpecification {

	return UserAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *User) ValueForAttribute(name string) any {

	switch name {
	case "ID":
		return o.ID
	case "archived":
		return o.Archived
	case "firstName":
		return o.FirstName
	case "lastName":
		return o.LastName
	case "namespace":
		return o.Namespace
	case "parentID":
		return o.ParentID
	case "parentType":
		return o.ParentType
	case "userName":
		return o.UserName
	}

	return nil
}

// UserAttributesMap represents the map of attribute for User.
var UserAttributesMap = map[string]AttributeSpecification{
	"ID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"Archived": {
		AllowedChoices: []string{},
		BSONFieldName:  "archived",
		ConvertedName:  "Archived",
		Description:    `the object is archived and not deleted.`,
		Exposed:        true,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"FirstName": {
		AllowedChoices: []string{},
		BSONFieldName:  "firstname",
		ConvertedName:  "FirstName",
		Description:    `The first name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "firstName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"LastName": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastname",
		ConvertedName:  "LastName",
		Description:    `The last name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"Namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"ParentID": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"ParentType": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
	"UserName": {
		AllowedChoices: []string{},
		BSONFieldName:  "username",
		ConvertedName:  "UserName",
		Description:    `the login.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "userName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// UserLowerCaseAttributesMap represents the map of attribute for User.
var UserLowerCaseAttributesMap = map[string]AttributeSpecification{
	"id": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "_id",
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
	"archived": {
		AllowedChoices: []string{},
		BSONFieldName:  "archived",
		ConvertedName:  "Archived",
		Description:    `the object is archived and not deleted.`,
		Exposed:        true,
		Getter:         true,
		Name:           "archived",
		Setter:         true,
		Stored:         true,
		Type:           "boolean",
	},
	"firstname": {
		AllowedChoices: []string{},
		BSONFieldName:  "firstname",
		ConvertedName:  "FirstName",
		Description:    `The first name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "firstName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"lastname": {
		AllowedChoices: []string{},
		BSONFieldName:  "lastname",
		ConvertedName:  "LastName",
		Description:    `The last name.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "lastName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
	"namespace": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "namespace",
		ConvertedName:  "Namespace",
		Description:    `The namespace of the object.`,
		Exposed:        true,
		Getter:         true,
		Name:           "namespace",
		Orderable:      true,
		ReadOnly:       true,
		Setter:         true,
		Stored:         true,
		Type:           "string",
	},
	"parentid": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parentid",
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
	"parenttype": {
		AllowedChoices: []string{},
		Autogenerated:  true,
		BSONFieldName:  "parenttype",
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
	"username": {
		AllowedChoices: []string{},
		BSONFieldName:  "username",
		ConvertedName:  "UserName",
		Description:    `the login.`,
		Exposed:        true,
		Filterable:     true,
		Name:           "userName",
		Orderable:      true,
		Required:       true,
		Stored:         true,
		Type:           "string",
	},
}

// SparseUsersList represents a list of SparseUsers
type SparseUsersList []*SparseUser

// Identity returns the identity of the objects in the list.
func (o SparseUsersList) Identity() Identity {

	return UserIdentity
}

// Copy returns a pointer to a copy the SparseUsersList.
func (o SparseUsersList) Copy() Identifiables {

	copy := slices.Clone(o)
	return &copy
}

// Append appends the objects to the a new copy of the SparseUsersList.
func (o SparseUsersList) Append(objects ...Identifiable) Identifiables {

	out := slices.Clone(o)
	for _, obj := range objects {
		out = append(out, obj.(*SparseUser))
	}

	return out
}

// List converts the object to an IdentifiablesList.
func (o SparseUsersList) List() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i]
	}

	return out
}

// DefaultOrder returns the default ordering fields of the content.
func (o SparseUsersList) DefaultOrder() []string {

	return []string{}
}

// ToPlain returns the SparseUsersList converted to UsersList.
func (o SparseUsersList) ToPlain() IdentifiablesList {

	out := make(IdentifiablesList, len(o))
	for i := range len(o) {
		out[i] = o[i].ToPlain()
	}

	return out
}

// Version returns the version of the content.
func (o SparseUsersList) Version() int {

	return 1
}

// SparseUser represents the sparse version of a user.
type SparseUser struct {
	// The identifier.
	ID *string `json:"ID,omitempty" msgpack:"ID,omitempty" bson:"-" mapstructure:"ID,omitempty"`

	// the object is archived and not deleted.
	Archived *bool `json:"archived,omitempty" msgpack:"archived,omitempty" bson:"archived,omitempty" mapstructure:"archived,omitempty"`

	// The first name.
	FirstName *string `json:"firstName,omitempty" msgpack:"firstName,omitempty" bson:"firstname,omitempty" mapstructure:"firstName,omitempty"`

	// The last name.
	LastName *string `json:"lastName,omitempty" msgpack:"lastName,omitempty" bson:"lastname,omitempty" mapstructure:"lastName,omitempty"`

	// The namespace of the object.
	Namespace *string `json:"namespace,omitempty" msgpack:"namespace,omitempty" bson:"namespace,omitempty" mapstructure:"namespace,omitempty"`

	// The identifier of the parent of the object.
	ParentID *string `json:"parentID,omitempty" msgpack:"parentID,omitempty" bson:"parentid,omitempty" mapstructure:"parentID,omitempty"`

	// The type of the parent of the object.
	ParentType *string `json:"parentType,omitempty" msgpack:"parentType,omitempty" bson:"parenttype,omitempty" mapstructure:"parentType,omitempty"`

	// the login.
	UserName *string `json:"userName,omitempty" msgpack:"userName,omitempty" bson:"username,omitempty" mapstructure:"userName,omitempty"`

	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewSparseUser returns a new  SparseUser.
func NewSparseUser() *SparseUser {
	return &SparseUser{}
}

// Identity returns the Identity of the sparse object.
func (o *SparseUser) Identity() Identity {

	return UserIdentity
}

// Identifier returns the value of the sparse object's unique identifier.
func (o *SparseUser) Identifier() string {

	if o.ID == nil {
		return ""
	}
	return *o.ID
}

// SetIdentifier sets the value of the sparse object's unique identifier.
func (o *SparseUser) SetIdentifier(id string) {

	if id != "" {
		o.ID = &id
	} else {
		o.ID = nil
	}
}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseUser) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesSparseUser{}

	if o.ID != nil {
		s.ID = bson.ObjectIdHex(*o.ID)
	}
	if o.Archived != nil {
		s.Archived = o.Archived
	}
	if o.FirstName != nil {
		s.FirstName = o.FirstName
	}
	if o.LastName != nil {
		s.LastName = o.LastName
	}
	if o.Namespace != nil {
		s.Namespace = o.Namespace
	}
	if o.ParentID != nil {
		s.ParentID = o.ParentID
	}
	if o.ParentType != nil {
		s.ParentType = o.ParentType
	}
	if o.UserName != nil {
		s.UserName = o.UserName
	}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *SparseUser) SetBSON(raw bson.Raw) error {

	if o == nil {
		return nil
	}

	s := &mongoAttributesSparseUser{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	id := s.ID.Hex()
	o.ID = &id
	if s.Archived != nil {
		o.Archived = s.Archived
	}
	if s.FirstName != nil {
		o.FirstName = s.FirstName
	}
	if s.LastName != nil {
		o.LastName = s.LastName
	}
	if s.Namespace != nil {
		o.Namespace = s.Namespace
	}
	if s.ParentID != nil {
		o.ParentID = s.ParentID
	}
	if s.ParentType != nil {
		o.ParentType = s.ParentType
	}
	if s.UserName != nil {
		o.UserName = s.UserName
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *SparseUser) Version() int {

	return 1
}

// ToPlain returns the plain version of the sparse model.
func (o *SparseUser) ToPlain() PlainIdentifiable {

	out := NewUser()
	if o.ID != nil {
		out.ID = *o.ID
	}
	if o.Archived != nil {
		out.Archived = *o.Archived
	}
	if o.FirstName != nil {
		out.FirstName = *o.FirstName
	}
	if o.LastName != nil {
		out.LastName = *o.LastName
	}
	if o.Namespace != nil {
		out.Namespace = *o.Namespace
	}
	if o.ParentID != nil {
		out.ParentID = *o.ParentID
	}
	if o.ParentType != nil {
		out.ParentType = *o.ParentType
	}
	if o.UserName != nil {
		out.UserName = *o.UserName
	}

	return out
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *SparseUser) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *SparseUser) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// GetArchived returns the Archived of the receiver.
func (o *SparseUser) GetArchived() (out bool) {

	if o.Archived == nil {
		return
	}

	return *o.Archived
}

// SetArchived sets the property Archived of the receiver using the address of the given value.
func (o *SparseUser) SetArchived(archived bool) {

	o.Archived = &archived
}

// GetNamespace returns the Namespace of the receiver.
func (o *SparseUser) GetNamespace() (out string) {

	if o.Namespace == nil {
		return
	}

	return *o.Namespace
}

// SetNamespace sets the property Namespace of the receiver using the address of the given value.
func (o *SparseUser) SetNamespace(namespace string) {

	o.Namespace = &namespace
}

// DeepCopy returns a deep copy if the SparseUser.
func (o *SparseUser) DeepCopy() *SparseUser {

	if o == nil {
		return nil
	}

	out := &SparseUser{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *SparseUser.
func (o *SparseUser) DeepCopyInto(out *SparseUser) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy SparseUser: %s", err))
	}

	*out = *target.(*SparseUser)
}

type mongoAttributesUser struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Archived   bool          `bson:"archived"`
	FirstName  string        `bson:"firstname"`
	LastName   string        `bson:"lastname"`
	Namespace  string        `bson:"namespace,omitempty"`
	ParentID   string        `bson:"parentid"`
	ParentType string        `bson:"parenttype"`
	UserName   string        `bson:"username"`
}
type mongoAttributesSparseUser struct {
	ID         bson.ObjectId `bson:"_id,omitempty"`
	Archived   *bool         `bson:"archived,omitempty"`
	FirstName  *string       `bson:"firstname,omitempty"`
	LastName   *string       `bson:"lastname,omitempty"`
	Namespace  *string       `bson:"namespace,omitempty"`
	ParentID   *string       `bson:"parentid,omitempty"`
	ParentType *string       `bson:"parenttype,omitempty"`
	UserName   *string       `bson:"username,omitempty"`
}

// Root represents the model of a root
type Root struct {
	ModelVersion int `json:"-" msgpack:"-" bson:"_modelversion"`
}

// NewRoot returns a new *Root
func NewRoot() *Root {

	return &Root{
		ModelVersion: 1,
	}
}

// Identity returns the Identity of the object.
func (o *Root) Identity() Identity {

	return RootIdentity
}

// Identifier returns the value of the object's unique identifier.
func (o *Root) Identifier() string {

	return ""
}

// SetIdentifier sets the value of the object's unique identifier.
func (o *Root) SetIdentifier(id string) {

}

// GetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Root) GetBSON() (any, error) {

	if o == nil {
		return nil, nil
	}

	s := &mongoAttributesRoot{}

	return s, nil
}

// SetBSON implements the bson marshaling interface.
// This is used to transparently convert ID to MongoDBID as ObectID.
func (o *Root) SetBSON(raw bson.Raw) error {

	if o == nil || raw.Kind == bson.ElementNil {
		return bson.ErrSetZero
	}

	s := &mongoAttributesRoot{}
	if err := raw.Unmarshal(s); err != nil {
		return err
	}

	return nil
}

// Version returns the hardcoded version of the model.
func (o *Root) Version() int {

	return 1
}

// BleveType implements the bleve.Classifier Interface.
func (o *Root) BleveType() string {

	return "root"
}

// DefaultOrder returns the list of default ordering fields.
func (o *Root) DefaultOrder() []string {

	return []string{}
}

// Doc returns the documentation for the object
func (o *Root) Doc() string {

	return `Root object of the API.`
}

func (o *Root) String() string {

	return fmt.Sprintf("<%s:%s>", o.Identity().Name, o.Identifier())
}

// EncryptAttributes encrypts the attributes marked as `encrypted` using the given encrypter.
func (o *Root) EncryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DecryptAttributes decrypts the attributes marked as `encrypted` using the given decrypter.
func (o *Root) DecryptAttributes(encrypter AttributeEncrypter) (err error) {

	return nil
}

// DeepCopy returns a deep copy if the Root.
func (o *Root) DeepCopy() *Root {

	if o == nil {
		return nil
	}

	out := &Root{}
	o.DeepCopyInto(out)

	return out
}

// DeepCopyInto copies the receiver into the given *Root.
func (o *Root) DeepCopyInto(out *Root) {

	target, err := copystructure.Copy(o)
	if err != nil {
		panic(fmt.Sprintf("Unable to deepcopy Root: %s", err))
	}

	*out = *target.(*Root)
}

// Validate valides the current information stored into the structure.
func (o *Root) Validate() error {

	ResetDefaultForZeroValues(o)

	errors := Errors{}
	requiredErrors := Errors{}

	if len(requiredErrors) > 0 {
		return requiredErrors
	}

	if len(errors) > 0 {
		return errors
	}

	return nil
}

// SpecificationForAttribute returns the AttributeSpecification for the given attribute name key.
func (*Root) SpecificationForAttribute(name string) AttributeSpecification {

	if v, ok := RootAttributesMap[name]; ok {
		return v
	}

	// We could not find it, so let's check on the lower case indexed spec map
	return RootLowerCaseAttributesMap[name]
}

// AttributeSpecifications returns the full attribute specifications map.
func (*Root) AttributeSpecifications() map[string]AttributeSpecification {

	return RootAttributesMap
}

// ValueForAttribute returns the value for the given attribute.
// This is a very advanced function that you should not need but in some
// very specific use cases.
func (o *Root) ValueForAttribute(name string) any {

	switch name {
	}

	return nil
}

// RootAttributesMap represents the map of attribute for Root.
var RootAttributesMap = map[string]AttributeSpecification{}

// RootLowerCaseAttributesMap represents the map of attribute for Root.
var RootLowerCaseAttributesMap = map[string]AttributeSpecification{}

type mongoAttributesRoot struct {
}

var (
	identityNamesMap = map[string]Identity{
		"list":    ListIdentity,
		"root":    RootIdentity,
		"subtask": SubtaskIdentity,
		"task":    TaskIdentity,
		"user":    UserIdentity,
	}

	identitycategoriesMap = map[string]Identity{
		"lists":    ListIdentity,
		"root":     RootIdentity,
		"subtasks": SubtaskIdentity,
		"tasks":    TaskIdentity,
		"users":    UserIdentity,
	}

	aliasesMap = map[string]Identity{
		"lst": ListIdentity,
		"tsk": TaskIdentity,
		"usr": UserIdentity,
	}

	indexesMap = map[string][][]string{
		"list": {
			{":shard", ":unique", "zone", "zHash"},
			{"namespace", "importLabel"},
		},
		"root":    nil,
		"subtask": nil,
		"task":    nil,
		"user":    nil,
	}
)

// ModelVersion returns the current version of the model.
func ModelVersion() float64 { return 1 }

type modelManager struct{}

func (f modelManager) IdentityFromName(name string) Identity {

	return identityNamesMap[name]
}

func (f modelManager) IdentityFromCategory(category string) Identity {

	return identitycategoriesMap[category]
}

func (f modelManager) IdentityFromAlias(alias string) Identity {

	return aliasesMap[alias]
}

func (f modelManager) IdentityFromAny(any string) (i Identity) {

	if i = f.IdentityFromName(any); !i.IsEmpty() {
		return i
	}

	if i = f.IdentityFromCategory(any); !i.IsEmpty() {
		return i
	}

	return f.IdentityFromAlias(any)
}

func (f modelManager) Identifiable(identity Identity) Identifiable {

	switch identity {

	case ListIdentity:
		return NewList()
	case RootIdentity:
		return NewRoot()
	case SubtaskIdentity:
		return NewSubtask()
	case TaskIdentity:
		return NewTask()
	case UserIdentity:
		return NewUser()
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiable(identity Identity) SparseIdentifiable {

	switch identity {

	case ListIdentity:
		return NewSparseList()
	case SubtaskIdentity:
		return NewSparseSubtask()
	case TaskIdentity:
		return NewSparseTask()
	case UserIdentity:
		return NewSparseUser()
	default:
		return nil
	}
}

func (f modelManager) Indexes(identity Identity) [][]string {

	return indexesMap[identity.Name]
}

func (f modelManager) IdentifiableFromString(any string) Identifiable {

	return f.Identifiable(f.IdentityFromAny(any))
}

func (f modelManager) Identifiables(identity Identity) Identifiables {

	switch identity {

	case ListIdentity:
		return &ListsList{}
	case SubtaskIdentity:
		return &SubtasksList{}
	case TaskIdentity:
		return &TasksList{}
	case UserIdentity:
		return &UsersList{}
	default:
		return nil
	}
}

func (f modelManager) SparseIdentifiables(identity Identity) SparseIdentifiables {

	switch identity {

	case ListIdentity:
		return &SparseListsList{}
	case SubtaskIdentity:
		return &SparseSubtasksList{}
	case TaskIdentity:
		return &SparseTasksList{}
	case UserIdentity:
		return &SparseUsersList{}
	default:
		return nil
	}
}

func (f modelManager) IdentifiablesFromString(any string) Identifiables {

	return f.Identifiables(f.IdentityFromAny(any))
}

func (f modelManager) Relationships() RelationshipsRegistry {

	return relationshipsRegistry
}

func (f modelManager) AllIdentities() []Identity {
	return AllIdentities()
}

func (f modelManager) DetachedFromString(name string) any {

	switch name {

	default:
		return nil
	}
}

var manager = modelManager{}

// Manager returns the model ModelManager.
func Manager() ModelManager { return manager }

// AllIdentities returns all existing identities.
func AllIdentities() []Identity {

	return []Identity{
		ListIdentity,
		RootIdentity,
		SubtaskIdentity,
		TaskIdentity,
		UserIdentity,
	}
}

// AliasesForIdentity returns all the aliases for the given identity.
func AliasesForIdentity(identity Identity) []string {

	switch identity {
	case ListIdentity:
		return []string{
			"lst",
		}
	case RootIdentity:
		return []string{}
	case SubtaskIdentity:
		return []string{}
	case TaskIdentity:
		return []string{
			"tsk",
		}
	case UserIdentity:
		return []string{
			"usr",
		}
	}

	return nil
}

var relationshipsRegistry RelationshipsRegistry

func init() {

	relationshipsRegistry = RelationshipsRegistry{}

	relationshipsRegistry[ListIdentity] = &Relationship{
		Create: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rlcp1",
						Type: "string",
					},
					{
						Name: "rlcp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "lup1",
						Type: "string",
					},
					{
						Name: "lup2",
						Type: "boolean",
					},
				},
			},
		},
		Patch: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "lup1",
						Type: "string",
					},
					{
						Name: "lup2",
						Type: "boolean",
					},
				},
			},
		},
		Delete: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "ldp1",
						Type: "string",
					},
					{
						Name: "ldp2",
						Type: "boolean",
					},
				},
			},
		},
		Retrieve: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "lgp1",
						Type: "string",
					},
					{
						Name: "lgp2",
						Type: "boolean",
					},
					{
						Name: "sAp1",
						Type: "string",
					},
					{
						Name: "sAp2",
						Type: "boolean",
					},
					{
						Name: "sBp1",
						Type: "string",
					},
					{
						Name: "sBp2",
						Type: "boolean",
					},
				},
			},
		},
		RetrieveMany: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rlgmp1",
						Type: "string",
					},
					{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rlgmp1",
						Type: "string",
					},
					{
						Name: "rlgmp2",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[RootIdentity] = &Relationship{}

	relationshipsRegistry[SubtaskIdentity] = &Relationship{
		Update: map[string]*RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*RelationshipInfo{
			"root": {},
		},
	}

	relationshipsRegistry[TaskIdentity] = &Relationship{
		Create: map[string]*RelationshipInfo{
			"list": {
				Parameters: []ParameterDefinition{
					{
						Name: "ltcp1",
						Type: "string",
					},
					{
						Name: "ltcp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*RelationshipInfo{
			"root": {},
		},
		Retrieve: map[string]*RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*RelationshipInfo{
			"list": {
				Parameters: []ParameterDefinition{
					{
						Name: "ltgp1",
						Type: "string",
					},
					{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*RelationshipInfo{
			"list": {
				Parameters: []ParameterDefinition{
					{
						Name: "ltgp1",
						Type: "string",
					},
					{
						Name: "ltgp2",
						Type: "boolean",
					},
				},
			},
		},
	}

	relationshipsRegistry[UserIdentity] = &Relationship{
		Create: map[string]*RelationshipInfo{
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rucp1",
						Type: "string",
					},
					{
						Name: "rucp2",
						Type: "boolean",
					},
				},
			},
		},
		Update: map[string]*RelationshipInfo{
			"root": {},
		},
		Patch: map[string]*RelationshipInfo{
			"root": {},
		},
		Delete: map[string]*RelationshipInfo{
			"root": {
				RequiredParameters: NewParametersRequirement(
					[][][]string{
						{
							{
								"confirm",
							},
						},
					},
				),
				Parameters: []ParameterDefinition{
					{
						Name: "confirm",
						Type: "boolean",
					},
				},
			},
		},
		Retrieve: map[string]*RelationshipInfo{
			"root": {},
		},
		RetrieveMany: map[string]*RelationshipInfo{
			"list": {},
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rugmp1",
						Type: "string",
					},
					{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			},
		},
		Info: map[string]*RelationshipInfo{
			"list": {},
			"root": {
				Parameters: []ParameterDefinition{
					{
						Name: "rugmp1",
						Type: "string",
					},
					{
						Name: "rugmp2",
						Type: "boolean",
					},
				},
			},
		},
	}

}
