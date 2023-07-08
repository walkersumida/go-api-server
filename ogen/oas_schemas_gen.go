// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"fmt"

	"github.com/go-faster/errors"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// DeleteItemOK is response for DeleteItem operation.
type DeleteItemOK struct{}

// Ref: #/components/schemas/error
type Error struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

// Ref: #/components/schemas/Item
type Item struct {
	ID     OptInt64      `json:"id"`
	Name   string        `json:"name"`
	Status OptItemStatus `json:"status"`
}

// GetID returns the value of ID.
func (s *Item) GetID() OptInt64 {
	return s.ID
}

// GetName returns the value of Name.
func (s *Item) GetName() string {
	return s.Name
}

// GetStatus returns the value of Status.
func (s *Item) GetStatus() OptItemStatus {
	return s.Status
}

// SetID sets the value of ID.
func (s *Item) SetID(val OptInt64) {
	s.ID = val
}

// SetName sets the value of Name.
func (s *Item) SetName(val string) {
	s.Name = val
}

// SetStatus sets the value of Status.
func (s *Item) SetStatus(val OptItemStatus) {
	s.Status = val
}

// Item status.
// Ref: #/components/schemas/ItemStatus
type ItemStatus string

const (
	ItemStatusNotReady ItemStatus = "not_ready"
	ItemStatusReady    ItemStatus = "ready"
	ItemStatusDoing    ItemStatus = "doing"
	ItemStatusDone     ItemStatus = "done"
)

// MarshalText implements encoding.TextMarshaler.
func (s ItemStatus) MarshalText() ([]byte, error) {
	switch s {
	case ItemStatusNotReady:
		return []byte(s), nil
	case ItemStatusReady:
		return []byte(s), nil
	case ItemStatusDoing:
		return []byte(s), nil
	case ItemStatusDone:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *ItemStatus) UnmarshalText(data []byte) error {
	switch ItemStatus(data) {
	case ItemStatusNotReady:
		*s = ItemStatusNotReady
		return nil
	case ItemStatusReady:
		*s = ItemStatusReady
		return nil
	case ItemStatusDoing:
		*s = ItemStatusDoing
		return nil
	case ItemStatusDone:
		*s = ItemStatusDone
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// NewOptInt64 returns new OptInt64 with value set to v.
func NewOptInt64(v int64) OptInt64 {
	return OptInt64{
		Value: v,
		Set:   true,
	}
}

// OptInt64 is optional int64.
type OptInt64 struct {
	Value int64
	Set   bool
}

// IsSet returns true if OptInt64 was set.
func (o OptInt64) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt64) Reset() {
	var v int64
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt64) SetTo(v int64) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt64) Get() (v int64, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt64) Or(d int64) int64 {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptItemStatus returns new OptItemStatus with value set to v.
func NewOptItemStatus(v ItemStatus) OptItemStatus {
	return OptItemStatus{
		Value: v,
		Set:   true,
	}
}

// OptItemStatus is optional ItemStatus.
type OptItemStatus struct {
	Value ItemStatus
	Set   bool
}

// IsSet returns true if OptItemStatus was set.
func (o OptItemStatus) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptItemStatus) Reset() {
	var v ItemStatus
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptItemStatus) SetTo(v ItemStatus) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptItemStatus) Get() (v ItemStatus, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptItemStatus) Or(d ItemStatus) ItemStatus {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// UpdateItemOK is response for UpdateItem operation.
type UpdateItemOK struct{}