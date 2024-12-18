// Code generated by protoc-gen-go-gorm-enum. DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"fmt"
)

func (x *TransactionType) Scan(value any) error {
	*x = TransactionType(TransactionType_value[string(value.([]byte))])

	return nil
}

func (x TransactionType) Value() (driver.Value, error) {
	return x.String(), nil
}

func TransactionType_FromString(str string) (TransactionType, error) {
	value, ok := TransactionType_value[str]
	if !ok {
		return TransactionType(0), fmt.Errorf("cannot parse TransactionType from %s", str)
	}

	return TransactionType(value), nil
}

func TransactionType_MustParseFromString(str string) TransactionType {
	value, ok := TransactionType_value[str]
	if !ok {
		panic(fmt.Errorf("cannot parse TransactionType from %s", str))
	}

	return TransactionType(value)
}
