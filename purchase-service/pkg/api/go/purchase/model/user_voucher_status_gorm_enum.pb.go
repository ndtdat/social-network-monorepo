// Code generated by protoc-gen-go-gorm-enum. DO NOT EDIT.

package model

import (
	"database/sql/driver"
	"fmt"
)

func (x *UserVoucherStatus) Scan(value any) error {
	*x = UserVoucherStatus(UserVoucherStatus_value[string(value.([]byte))])

	return nil
}

func (x UserVoucherStatus) Value() (driver.Value, error) {
	return x.String(), nil
}

func UserVoucherStatus_FromString(str string) (UserVoucherStatus, error) {
	value, ok := UserVoucherStatus_value[str]
	if !ok {
		return UserVoucherStatus(0), fmt.Errorf("cannot parse UserVoucherStatus from %s", str)
	}

	return UserVoucherStatus(value), nil
}

func UserVoucherStatus_MustParseFromString(str string) UserVoucherStatus {
	value, ok := UserVoucherStatus_value[str]
	if !ok {
		panic(fmt.Errorf("cannot parse UserVoucherStatus from %s", str))
	}

	return UserVoucherStatus(value)
}
