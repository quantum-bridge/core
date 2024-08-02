package env

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"math/big"
	"net/url"
	"reflect"
)

// Hook is a function type that takes a value and returns a reflect.Value and an error.
type Hook func(value interface{}) (reflect.Value, error)

// Hooks is a map where the key is a string and the value is a Hook function.
type Hooks map[string]Hook

// BaseHooks is a map of base hook functions for different types.
var BaseHooks = Hooks{
	"string": func(value interface{}) (reflect.Value, error) {
		// Convert the value to a string.
		result, err := cast.ToStringE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse string")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},
	"*string": func(value interface{}) (reflect.Value, error) {
		// Convert the value to a string.
		result, err := cast.ToStringE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse string")
		}
		// Return the converted value as a pointer.
		return reflect.ValueOf(&result), nil
	},
	"[]int64": func(value interface{}) (reflect.Value, error) {
		// Initialize an empty slice of int64.
		var a []int64

		// Switch on the type of the value.
		switch v := value.(type) {
		case []int64:
			// If the value is already a slice of int64, return it.
			return reflect.ValueOf(value), nil
		case []int:
			// If the value is a slice of int, convert it to a slice of int64.
			for _, intValue := range v {
				a = append(a, int64(intValue))
			}

			// Return the converted slice.
			return reflect.ValueOf(a), nil
		case []interface{}:
			// If the value is a slice of interface{}, try to convert each element to int64.
			for i, u := range v {
				int64Value, err := cast.ToInt64E(u)
				if err != nil {
					// If the conversion fails, return an error.
					return reflect.Value{}, fmt.Errorf("failed to cast slice element number %d: %#v of type %T into int64", i, value, value)
				}
				a = append(a, int64Value)
			}
			// Return the converted slice.
			return reflect.ValueOf(a), nil
		case interface{}:
			// If the value is a single interface{}, try to convert it to int64.
			int64Value, err := cast.ToInt64E(value)
			if err != nil {
				// If the conversion fails, return an error.
				return reflect.Value{}, fmt.Errorf("failed to cast %#v of type %T to int64", value, value)
			}

			// Return the converted value as a slice of int64.
			return reflect.ValueOf([]int64{int64Value}), nil
		default:
			// If the value is of an unsupported type, return an error.
			return reflect.Value{}, fmt.Errorf("failed to cast %#v of type %T to []int64", value, value)
		}
	},
	// Hook for []string type. It tries to convert the value to a slice of strings.
	"[]string": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToStringSliceE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse []string")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for int type. It tries to convert the value to an integer.
	"int": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToIntE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse int")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for int32 type. It tries to convert the value to an int32.
	"int32": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToInt32E(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse int32")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for int64 type. It tries to convert the value to an int64.
	"int64": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToInt64E(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse int64")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for uint type. It tries to convert the value to an unsigned integer.
	"uint": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToUintE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse uint")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for uint32 type. It tries to convert the value to an uint32.
	"uint32": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToUint32E(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse uint32")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for uint64 type. It tries to convert the value to an uint64.
	"uint64": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToUint64E(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse uint64")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for float64 type. It tries to convert the value to a float64.
	"float64": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToFloat64E(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse float64")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for bool type. It tries to convert the value to a boolean.
	"bool": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToBoolE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse bool")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for *bool type. It tries to convert the value to a boolean and returns a pointer to it.
	"*bool": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToBoolE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse bool")
		}

		// Return the converted value as a pointer.
		return reflect.ValueOf(&result), nil
	},

	// Hook for time.Time type. It tries to convert the value to a time.Time.
	"time.Time": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToTimeE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse time")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for *time.Time type. It tries to convert the value to a time.Time and returns a pointer to it.
	"*time.Time": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToTimeE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse time pointer")
		}

		// Return the converted value as a pointer.
		return reflect.ValueOf(&result), nil
	},

	// Hook for time.Duration type. It tries to convert the value to a time.Duration.
	"time.Duration": func(value interface{}) (reflect.Value, error) {
		result, err := cast.ToDurationE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse duration")
		}

		// Return the converted value.
		return reflect.ValueOf(result), nil
	},

	// Hook for *time.Duration type. It tries to convert the value to a time.Duration and returns a pointer to it.
	"*time.Duration": func(value interface{}) (reflect.Value, error) {
		if value == nil {
			// If the value is nil, return a nil reflect.Value.
			return reflect.ValueOf(nil), nil
		}
		result, err := cast.ToDurationE(value)
		if err != nil {
			// If the conversion fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse duration")
		}

		// Return the converted value as a pointer.
		return reflect.ValueOf(&result), nil
	},

	// Hook for *big.Int type. It tries to convert the value to a *big.Int.
	"*big.Int": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			// If the value is a string, try to convert it to a big integer.
			i, ok := new(big.Int).SetString(v, 10)
			if !ok {
				// If the conversion fails, return an error.
				return reflect.Value{}, errors.New("failed to parse")
			}

			// Return the converted value.
			return reflect.ValueOf(i), nil
		case int:
			// If the value is an integer, convert it to a big integer.
			return reflect.ValueOf(big.NewInt(int64(v))), nil
		default:
			// If the value is of an unsupported type, return an error.
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},

	// Hook for *uint64 type. It tries to convert the value to a *uint64.
	"*uint64": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			// If the value is a string, try to convert it to an unsigned integer.
			puint, err := cast.ToUint64E(v)
			if err != nil {
				// If the conversion fails, return an error.
				return reflect.Value{}, errors.New("failed to parse")
			}

			// Return the converted value as a pointer.
			return reflect.ValueOf(&puint), nil
		default:
			// If the value is of an unsupported type, return an error.
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},

	// Hook for *url.URL type. It tries to convert the value to a *url.URL.
	"*url.URL": func(value interface{}) (reflect.Value, error) {
		switch v := value.(type) {
		case string:
			// If the value is a string, try to parse it as a URL.
			u, err := url.Parse(v)
			if err != nil {
				// If the parsing fails, wrap and return the error.
				return reflect.Value{}, errors.Wrap(err, "failed to parse url")
			}

			// Return the parsed URL.
			return reflect.ValueOf(u), nil
		case nil:
			// If the value is nil, return a nil reflect.Value.
			return reflect.ValueOf(nil), nil
		default:
			// If the value is of an unsupported type, return an error.
			return reflect.Value{}, fmt.Errorf("unsupported conversion from %T", value)
		}
	},

	// Hook for json.RawMessage type. It tries to convert the value to a json.RawMessage.
	"json.RawMessage": func(value interface{}) (reflect.Value, error) {
		if value == nil {
			// If the value is nil, return a nil reflect.Value.
			return reflect.Value{}, nil
		}

		var params map[string]interface{}

		switch s := value.(type) {
		case map[interface{}]interface{}:
			// If the value is a map with interface{} keys, convert it to a map with string keys.
			params = make(map[string]interface{})
			for key, value := range s {
				params[key.(string)] = value
			}
		case map[string]interface{}:
			// If the value is already a map with string keys, use it as is.
			params = s
		default:
			// If the value is of an unsupported type, return an error.
			return reflect.Value{}, errors.New("unexpected type while figure []json.RawMessage")
		}

		// Try to marshal the map into a JSON string.
		result, err := json.Marshal(params)
		if err != nil {
			// If the marshaling fails, wrap and return the error.
			return reflect.Value{}, errors.Wrap(err, "failed to parse json.RawMessage")
		}

		// Return the marshaled JSON as a json.RawMessage.
		return reflect.ValueOf(json.RawMessage(result)), nil
	},
}
