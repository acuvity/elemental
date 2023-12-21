// Copyright 2019 Aporeto Inc.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package elemental

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"reflect"
)

// An AttributeSpecifiable is the interface an object must implement in order to access specification of its attributes.
type AttributeSpecifiable interface {

	// SpecificationForAttribute returns the AttributeSpecification for
	// given attribute name
	SpecificationForAttribute(string) AttributeSpecification

	// AttributeSpecifications returns all the AttributeSpecification mapped by
	// attribute name
	AttributeSpecifications() map[string]AttributeSpecification

	// ValueForAttribute returns the value for the given attribute
	ValueForAttribute(name string) any
}

// AttributeEncrypter is the interface that must be
// implement to manage encrypted attributes.
type AttributeEncrypter interface {

	// EncryptString encrypts the given string and returns the encrypted version.
	EncryptString(string) (string, error)

	// DecryptString decrypts the given string and returns the encrypted version.
	DecryptString(string) (string, error)
}

// An AttributeSpecification represents all the metadata of an attribute.
//
// This information is coming from the Monolithe Specifications.
type AttributeSpecification struct {

	// AllowedChars is a regexp that will be used to validate
	// what value a string attribute can take.
	//
	// This is enforced by elemental.
	AllowedChars string

	// AllowedChoices is a list of possible values for an attribute.
	//
	// This is enforced by elemental.
	AllowedChoices []string

	// Autogenerated defines if the attribute is autogenerated by the server.
	// It can be used in conjunction with ReadOnly.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Autogenerated bool

	// Availability is reserved for later use.
	Availability string

	// BSONFieldName is the name of the field that will be used when encoding/decoding the field into Binary JSON format.
	BSONFieldName string

	// ConvertedName contains the name after local conversion.
	ConvertedName string

	// Channel is reserved for later use.
	Channel string

	// CreationOnly defines if the attribute can be set only during creation.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	CreationOnly bool

	// DefaultValue holds the default value declared in specification.
	DefaultValue any

	// Deprecated defines if the attribute is deprecated.
	Deprecated bool

	// Description contains the description of the attribute.
	Description string

	// Exposed defines if the attribute is exposed through the north bound API.
	Exposed bool

	// Filterable defines if it is possible to filter based on this attribute.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Filterable bool

	// ForeignKey defines if the attribute is a foreign key.
	ForeignKey bool

	// Getter defines if the attribute needs to define a getter method.
	// This is useful if you can to define an Interface based on this attribute.
	Getter bool

	// Identifier defines if the attribute is used the access key from the
	// northbound API.
	Identifier bool

	// Index defines if the attribute is indexed or not.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Index bool

	// MaxLength defines what is the maximun length of the attribute.
	// This only makes sense if the type is a string.
	//
	// This is enforced by elemental.
	MaxLength uint

	// MaxValue defines what is the maximun value of the attribute.
	// This only makes sense if the type has a numeric type.
	//
	// This is enforced by elemental.
	MaxValue float64

	// MinLength defines what is the minimum length of the attribute.
	// This only makes sense if the type is a string.
	//
	// This is enforced by elemental.
	MinLength uint

	// MinValue defines what is the minimum value of the attribute.
	// This only makes sense if the type has a numeric type.
	//
	// This is enforced by elemental.
	MinValue float64

	// Name defines what is the name of the attribute.
	// This will be the raw Monolithe Specification name, without
	// Go translation.
	Name string

	// Orderable defines if it is possible to order based on the value of this attribute.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Orderable bool

	// PrimaryKey defines if the attribute is used as a primary key.
	PrimaryKey bool

	// ReadOnly defines if the attribute is read only.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	ReadOnly bool

	// Required defines is the attribute must be set or not.
	//
	// This is enforced by elemental.
	Required bool

	// Secret defines if the attribute is secret.
	// This is useful if you can to define perform sanity check on this field to be sure it
	// is not sent for instance.
	Secret bool

	// Setter defines if the attribute needs to define a setter method.
	// This is useful if you can to define an Interface based on this attribute.
	Setter bool

	// Signed indicates if the attribute's value should be used when
	// generating a signature for the object.
	Signed bool

	// Stored defines if the attribute will be stored in the northbound API.
	//
	// If this is true, the backend tags will be generated by Monolithe.
	Stored bool

	// SubType defines the Monolithe Subtype.
	SubType string

	// Transient defines if the attributes is transient or not.
	//
	// This is not enforced by elemental. You must write your own business logic to honor this.
	Transient bool

	// Type defines the raw Monolithe type.
	Type string

	// Encrypted defines if the attribute needs encryption.
	Encrypted bool
}

// ResetSecretAttributesValues will reset any attributes marked
// as `secret` in the given obj if it is an elemental.Identifiable
// or an elemental.Identifiables.
// The given Identifiables must implement the elemental.AttributeSpecifiable
// interface or this function will have no effect.
//
// If you pass anything else, this function does nothing.
func ResetSecretAttributesValues(obj any) {

	if obj == nil {
		return
	}

	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr && v.IsNil() {
		return
	}

	strip := func(o Identifiable) {

		oo := o
		if sp, ok := o.(SparseIdentifiable); ok {
			oo = sp.ToPlain()
		}

		if attrspec, ok := oo.(AttributeSpecifiable); ok {

			var rv, val reflect.Value

			for _, aspec := range attrspec.AttributeSpecifications() {

				if !aspec.Secret {
					continue
				}

				rv = reflect.Indirect(reflect.ValueOf(o))
				val = rv.FieldByName(aspec.ConvertedName)
				val.Set(reflect.Zero(val.Type()))
			}
		}
	}

	switch o := obj.(type) {

	case Identifiable:
		strip(o)

	case Identifiables:
		for _, i := range o.List() {
			strip(i)
		}
	}
}

// aesAttributeEncrypter is an elemental.AttributeEncrypter
// using AES encryption.
type aesAttributeEncrypter struct {
	passphrase []byte
}

// NewAESAttributeEncrypter returns a new elemental.AttributeEncrypter
// implementing AES encryption.
func NewAESAttributeEncrypter(passphrase string) (AttributeEncrypter, error) {

	passbytes := []byte(passphrase)

	switch len(passbytes) {
	case 16:
		break
	case 24:
		break
	case 32:
		break
	default:
		return nil, fmt.Errorf("the length for AES Attribute encrypter must be 16, 24 or 32 bytes long (current length %d) to select AES-128, AES-192 or AES-256", len(passbytes))
	}

	return &aesAttributeEncrypter{
		passphrase: passbytes,
	}, nil
}

// EncryptString encrypts the given string.
func (e *aesAttributeEncrypter) EncryptString(value string) (string, error) {

	if value == "" {
		return "", nil
	}

	data := []byte(value)

	c, err := aes.NewCipher(e.passphrase)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(gcm.Seal(nonce, nonce, data, nil)), nil
}

// DecryptString decrypts the given string.
func (e *aesAttributeEncrypter) DecryptString(value string) (string, error) {

	if value == "" {
		return "", nil
	}

	data, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	c, err := aes.NewCipher(e.passphrase)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", fmt.Errorf("data is too small")
	}

	out, err := gcm.Open(nil, data[:nonceSize], data[nonceSize:], nil)
	if err != nil {
		return "", err
	}

	return string(out), nil
}
