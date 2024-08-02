package env

import (
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/pkg/common"
	"github.com/quantum-bridge/core/pkg/convert"
	"reflect"
	"strings"
)

const (
	// keyTag is the key tag that holds the key name in the struct for the config.
	keyTag = "config"
	// ignore is the ignore tag.
	ignore = "-"
	// required is the required tag.
	required = "required"
	// nonZero is the non-zero tag.
	nonZero = "nonzero"
)

// Validatable is the interface that implements the Validate method.
type Validatable interface {
	// Validate validates the struct.
	Validate() error
}

// IsZeroer is the interface that implements the IsZero method.
type IsZeroer interface {
	// IsZero returns true if the struct is zero.
	IsZero() bool
}

// Validate validates the struct.
func isZero(value reflect.Value) bool {
	// Set the value to the variable.
	kind := value.Kind()
	// Check if the value is zero.
	if z, ok := value.Interface().(IsZeroer); ok {
		// Check if the value is a pointer or an interface and is nil.
		if (kind == reflect.Ptr || kind == reflect.Interface) && value.IsNil() {
			return true
		}

		// Check if the value is zero.
		return z.IsZero()
	}

	// Check if the value is zero.
	switch kind {
	case reflect.String:
		return len(value.String()) == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	case reflect.Slice:
		return value.Len() == 0
	case reflect.Map:
		return value.Len() == 0
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Struct:
		// Check if the struct is zero.
		vt := value.Type()

		// Iterate over the fields of the struct.
		for i := value.NumField() - 1; i >= 0; i-- {
			// Check if the field is unexported.
			if vt.Field(i).PkgPath != "" {
				continue // Private field
			}

			// Check if the field is not zero.
			if !isZero(value.Field(i)) {
				return false
			}
		}

		return true
	}

	return false
}

// parseFieldTag parses the field tag and returns the tag.
func parseFieldTag(field reflect.StructField, key string) (*Tag, error) {
	tag := &Tag{}

	// Get the tag from the field.
	fieldTag := field.Tag.Get(key)
	splitedTag := strings.Split(fieldTag, ",")

	// Check if the tag is ignored.
	if len(splitedTag) == 1 && splitedTag[0] == ignore {
		return nil, nil
	}

	// Check if the tag is empty.
	if len(splitedTag) == 0 {
		tag.Key = ""
	} else {
		tag.Key = splitedTag[0]
	}

	// Check if the tag key is empty.
	if tag.Key == "" {
		tag.Key = convert.ToSnakeCase(field.Name)
	}

	// Check if the tag has more than one rule.
	if len(splitedTag) > 1 {
		if common.Contains(splitedTag, ignore) {
			return nil, errors.New("ignore tag must be the only tag")
		}

		// Iterate over the rules of the tag and set the tag if it is required or non-zero.
		for _, rule := range splitedTag[1:] {
			switch rule {
			case required:
				tag.Required = true
			case nonZero:
				tag.NonZero = true
			default:
				return nil, errors.Errorf("unknown tag %s", rule)
			}
		}
	}

	return tag, nil
}
