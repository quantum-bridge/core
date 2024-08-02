package convert

import (
	"github.com/spf13/cast"
	"reflect"
)

// SetBool sets a boolean value to a reflect.Value.
func SetBool(valueRef reflect.Value, sourceValue interface{}) error {
	boolValue, err := cast.ToBoolE(sourceValue)
	valueRef.SetBool(boolValue)

	return err
}

// SetInt sets an integer value to a reflect.Value.
func SetInt(valueRef reflect.Value, sourceValue interface{}) error {
	intValue, err := cast.ToInt64E(sourceValue)
	valueRef.SetInt(intValue)

	return err
}

// SetUint sets an unsigned integer value to a reflect.Value.
func SetUint(valueRef reflect.Value, sourceValue interface{}) error {
	uintValue, err := cast.ToUint64E(sourceValue)
	valueRef.SetUint(uintValue)

	return err
}

// SetFloat sets a float value to a reflect.Value.
func SetFloat(valueRef reflect.Value, sourceValue interface{}) error {
	floatValue, err := cast.ToFloat64E(sourceValue)
	valueRef.SetFloat(floatValue)

	return err
}

// SetString sets a string value to a reflect.Value.
func SetString(valueRef reflect.Value, sourceValue interface{}) error {
	strValue, err := cast.ToStringE(sourceValue)
	valueRef.SetString(strValue)

	return err
}
