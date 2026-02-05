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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"strings"
)

// IsErrorWithCode returns true if the given error is an elemental.Error
// or elemental.Errors with the status set to the given code.
func IsErrorWithCode(err error, code int) bool {

	var c int
	var e1 Error
	var e2 Errors

	if errors.As(err, &e1) {
		c = e1.Code
	} else if errors.As(err, &e2) {
		c = e2.Code()
	}

	return c == code
}

// An Error represents a computational error.
//
// They can be encoded and sent back to the clients.
type Error struct {
	Code        int    `msgpack:"code" json:"code,omitempty"`
	Description string `msgpack:"description" json:"description"`
	Subject     string `msgpack:"subject" json:"subject"`
	Title       string `msgpack:"title" json:"title"`
	Data        any    `msgpack:"data" json:"data,omitempty"`
	Trace       string `msgpack:"trace" json:"trace,omitempty"`
}

// NewError returns a new Error.
func NewError(title, description, subject string, code int) Error {

	return NewErrorWithData(title, description, subject, code, nil)
}

// NewErrorWithData returns a new Error with the given opaque data.
func NewErrorWithData(title, description, subject string, code int, data any) Error {

	return Error{
		Code:        code,
		Description: description,
		Subject:     subject,
		Title:       title,
		Data:        data,
	}
}

func (e Error) Error() string {

	if e.Trace != "" {
		return fmt.Sprintf("error %d (%s): %s: %s [trace: %s]", e.Code, e.Subject, e.Title, e.Description, e.Trace)
	}

	return fmt.Sprintf("error %d (%s): %s: %s", e.Code, e.Subject, e.Title, e.Description)
}

// Errors represents a list of Error.
type Errors []Error

// NewErrors creates a new Errors.
func NewErrors(errors ...error) Errors {

	out := Errors{}
	if len(errors) == 0 {
		return out
	}

	return out.Append(errors...)
}

func (e Errors) Error() string {

	strs := make([]string, len(e))

	for i := range e {
		strs[i] = e[i].Error()
	}

	return strings.Join(strs, ", ")
}

// Code returns the code of the first error code in the Errors.
func (e Errors) Code() int {

	if len(e) == 0 {
		return -1
	}

	return e[0].Code
}

// Append returns returns a copy of the receiver containing
// also the given errors.
func (e Errors) Append(errs ...error) Errors {

	out := slices.Clone(e)

	for _, err := range errs {

		var e1 Error
		var e2 Errors
		if errors.As(err, &e1) {
			out = append(out, e1)
		} else if errors.As(err, &e2) {
			out = append(out, e2...)
		} else {
			out = append(out, NewError("Internal Server Error", err.Error(), "elemental", http.StatusInternalServerError))
		}
	}

	return out
}

// Trace returns Errors with all inside Error marked with the
// given trace ID.
func (e Errors) Trace(id string) Errors {

	out := make(Errors, 0, len(e))
	for _, err := range e {
		err.Trace = id
		out = append(out, err)
	}

	return out
}

// DecodeErrors decodes the given bytes into a en elemental.Errors.
func DecodeErrors(data []byte) (Errors, error) {

	es := []Error{}
	if err := json.Unmarshal(data, &es); err != nil {
		return nil, err
	}

	e := NewErrors()
	for _, err := range es {
		e = append(e, err)
	}

	return e, nil
}

// IsValidationError returns true if the given error is a validation error
// with the given title for the given attribute.
func IsValidationError(err error, title string, attribute string) bool {

	var elementalError Error

	var e1 Errors
	var e2 Error

	if errors.As(err, &e1) {
		if e1.Code() != http.StatusUnprocessableEntity {
			return false
		}
		if len(e1) != 1 {
			return false
		}
		elementalError = e1[0]
	} else if errors.As(err, &e2) {
		if e2.Code != http.StatusUnprocessableEntity {
			return false
		}
		elementalError = e2
	} else {
		return false
	}

	if elementalError.Title != title {
		return false
	}

	if elementalError.Data == nil {
		return false
	}

	m, ok := elementalError.Data.(map[string]any)
	if !ok {
		return false
	}

	return m["attribute"].(string) == attribute
}

func InjectAttributePath(err error, path string) {

	if path == "" {
		return
	}

	eerrs := Errors{}
	if !errors.As(err, &eerrs) {
		eerr := Error{}
		if !errors.As(err, &eerr) {
			return
		}
		eerrs = append(eerrs, eerr)
	}

	for _, eerr := range eerrs {

		if data, ok := eerr.Data.(map[string]any); ok {
			if v, ok := data["attribute"]; ok {
				if vs, ok := v.(string); ok {
					data["attribute"] = fmt.Sprintf("%s/%s", path, vs)
					eerr.Data = data
				}
			}
		}

		if data, ok := eerr.Data.(map[string]string); ok {
			if v, ok := data["attribute"]; ok {
				data["attribute"] = fmt.Sprintf("%s/%s", path, v)
				eerr.Data = data
			}
		}
	}
}
