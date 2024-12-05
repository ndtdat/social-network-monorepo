// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: common/sorter.proto

package common

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Sorter with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Sorter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Sorter with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in SorterMultiError, or nil if none found.
func (m *Sorter) ValidateAll() error {
	return m.validate(true)
}

func (m *Sorter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Field

	if _, ok := _Sorter_Order_InLookup[m.GetOrder()]; !ok {
		err := SorterValidationError{
			field:  "Order",
			reason: "value must be in list [desc asc]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SorterMultiError(errors)
	}

	return nil
}

// SorterMultiError is an error wrapping multiple validation errors returned by
// Sorter.ValidateAll() if the designated constraints aren't met.
type SorterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SorterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SorterMultiError) AllErrors() []error { return m }

// SorterValidationError is the validation error returned by Sorter.Validate if
// the designated constraints aren't met.
type SorterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SorterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SorterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SorterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SorterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SorterValidationError) ErrorName() string { return "SorterValidationError" }

// Error satisfies the builtin error interface
func (e SorterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSorter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SorterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SorterValidationError{}

var _Sorter_Order_InLookup = map[string]struct{}{
	"desc": {},
	"asc":  {},
}