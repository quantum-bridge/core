package env

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/quantum-bridge/core/pkg/convert"
	bridgeErrors "github.com/quantum-bridge/core/pkg/errors"
	"reflect"
)

// Configuration is a struct that holds the configuration data and the target struct to be filled with the data.
type Configuration struct {
	sourceData interface{}
	config     interface{}
	dataHooks  Hooks
}

// NewConfiguration creates a new configuration instance with the given target struct.
func NewConfiguration(config interface{}) *Configuration {
	// Create a new configuration instance with the given target struct.
	hooks := make(Hooks)

	// Iterate over the BaseHooks map and add each key-value pair to the hooks map.
	for key, hook := range BaseHooks {
		hooks[key] = hook
	}

	// Return the configuration instance with the target struct and the hooks map.
	return &Configuration{
		config:    config,
		dataHooks: hooks,
	}
}

// From sets the source data for the configuration.
func (c *Configuration) From(sourceData map[string]interface{}) *Configuration {
	c.sourceData = sourceData

	return c
}

// Load fills the target struct with the source data.
func (c *Configuration) Load() error {
	// Check if the source data is nil.
	if c.sourceData == nil {
		return errors.New("source map is not set")
	}

	// Check if the config is nil.
	if c.config == nil {
		return errors.New("cannot figure out to nil target")
	}

	// Get the target value of the configuration.
	targetValue := reflect.ValueOf(c.config)
	if targetValue.Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	// Check if the target value is nil and if it can be set.
	if targetValue.IsNil() && !targetValue.CanSet() {
		return errors.New("target is not settable")
	}

	// Set the value of the target value.
	return c.setValue(targetValue.Elem(), c.sourceData)
}

// setValue sets the value of the target value.
func (c *Configuration) setValue(valueRef reflect.Value, sourceValue interface{}) error {
	// Check if the source value is nil.
	if sourceValue == nil {
		return nil
	}

	// Get the hook for the value reference type and check if it exists.
	hook, hasHook := c.dataHooks[valueRef.Type().String()]
	if hasHook {
		value, err := hook(sourceValue)
		if err != nil {
			return err
		}
		valueRef.Set(value)

		return nil
	}

	// Check if the value reference kind is a pointer and set the pointer value if it is to the type of the source value.
	var err error
	switch valueRef.Kind() {
	case reflect.Bool:
		err = convert.SetBool(valueRef, sourceValue)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		err = convert.SetInt(valueRef, sourceValue)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		err = convert.SetUint(valueRef, sourceValue)
	case reflect.Float32, reflect.Float64:
		err = convert.SetFloat(valueRef, sourceValue)
	case reflect.String:
		err = convert.SetString(valueRef, sourceValue)
	case reflect.Pointer:
		err = c.setPointer(valueRef, sourceValue)
	case reflect.Array:
		err = c.setArray(valueRef, sourceValue)
	case reflect.Slice:
		err = c.setSlice(valueRef, sourceValue)
	case reflect.Map:
		err = c.setMap(valueRef, sourceValue)
	case reflect.Struct:
		err = c.setStruct(valueRef, sourceValue)
	default:
		return errors.New(fmt.Sprintf("%s types are not supported", valueRef.Type().String()))
	}

	if val, ok := c.config.(Validatable); ok {
		return val.Validate()
	}

	return err
}

// setPointer sets the pointer value of the target value.
func (c *Configuration) setPointer(valueRef reflect.Value, sourceValue interface{}) error {
	// Allocate memory if the pointer is nil.
	if err := allocateMemoryIfPointerIsNil(&valueRef); err != nil {
		return err
	}

	// Set the value of the pointer.
	return c.setValue(valueRef.Elem(), sourceValue)
}

// allocateMemoryIfPointerIsNil allocates memory if the pointer is nil.
func allocateMemoryIfPointerIsNil(valuePtr *reflect.Value) error {
	// Check if the value pointer kind is not a pointer.
	if valuePtr.Kind() != reflect.Ptr {
		return errors.New("value is not a pointer")
	}

	// Check if the value pointer is nil.
	if valuePtr.IsNil() {
		// Set the value pointer to a new value.
		valuePtr.Set(reflect.New(valuePtr.Type().Elem()))
	}

	return nil
}

// setArray sets the array value of the target value.
func (c *Configuration) setArray(arrayRef reflect.Value, sourceValue interface{}) error {
	// Get the source array value and check if it is an array or a slice.
	sourceArray := reflect.ValueOf(sourceValue)
	// Check if the source array kind is not an array or a slice.
	if sourceArray.Kind() != reflect.Array && sourceArray.Kind() != reflect.Slice {
		return errors.New(fmt.Sprintf("can't set array from non-array value: expected type %s, actual type: %s",
			arrayRef.Type().String(), sourceArray.Type().String()))
	}

	// Check if the array reference length is not equal to the source array length.
	if arrayRef.Len() != sourceArray.Len() {
		return errors.New("array length mismatch")
	}

	// Iterate over the source array and set the value of the array reference.
	for i := 0; i < sourceArray.Len(); i++ {
		if err := c.setValue(arrayRef.Index(i), sourceArray.Index(i).Interface()); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set array element %d", i))
		}
	}

	return nil
}

// setSlice sets the slice value of the target value.
func (c *Configuration) setSlice(sliceRef reflect.Value, sourceValue interface{}) error {
	// Get the source array value and check if it is an array or a slice.
	sourceArray := reflect.ValueOf(sourceValue)
	// Check if the source array kind is not an array or a slice.
	if sourceArray.Kind() != reflect.Array && sourceArray.Kind() != reflect.Slice {
		return errors.Wrap(bridgeErrors.ErrNotValid, fmt.Sprintf("can't set slice from non-array value: expected type %s, actual type: %s",
			sliceRef.Type().String(), sourceArray.Type().String()))
	}

	// Create a new slice with the source array length.
	slice := reflect.MakeSlice(sliceRef.Type(), sourceArray.Len(), sourceArray.Len())
	for i := 0; i < sourceArray.Len(); i++ {
		if err := c.setValue(slice.Index(i), sourceArray.Index(i).Interface()); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set slice element %d", i))
		}
	}

	// Set the value of the slice reference.
	sliceRef.Set(slice)

	return nil
}

// setMap sets the map value of the target value.
func (c *Configuration) setMap(mapRef reflect.Value, sourceValue interface{}) error {
	// Check if the map reference key kind is not a string.
	if mapRef.Type().Key().Kind() != reflect.String {
		return errors.Wrap(bridgeErrors.ErrNotValid, fmt.Sprintf("map key type must be string or its alias, actual is %s", mapRef.Type().Key().String()))
	}

	// Get the source map value.
	sourceMap := reflect.ValueOf(sourceValue)
	// Check if the source map kind is not a map.
	if sourceMap.Kind() != reflect.Map {
		return errors.Wrap(bridgeErrors.ErrNotValid, fmt.Sprintf("can't set map from non-map value: expected type %s, actual type: %s",
			mapRef.Type().String(), sourceMap.Type().String()))
	}

	// Create a new map with the source map length.
	mapping := reflect.MakeMap(mapRef.Type())

	// Iterate over the source map and set the value of the map reference.
	for _, key := range sourceMap.MapKeys() {
		// Create a new map element value and set the value of the map reference.
		val := reflect.New(mapRef.Type().Elem()).Elem()
		if err := c.setValue(val, sourceMap.MapIndex(key).Interface()); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set map element %s", key.String()))
		}

		// Create a new map key value and set the value of the map reference.
		keyStr := reflect.New(mapRef.Type().Key()).Elem()

		// Convert the key string to a string and check if there is an error.
		if err := convert.SetString(keyStr, key.Interface()); err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't set map key %s", key.String()))
		}

		// Set the value of the map reference.
		mapping.SetMapIndex(keyStr, val)
	}

	// Set the value of the map reference.
	mapRef.Set(mapping)

	return nil
}

// setStruct sets the struct value of the target value.
func (c *Configuration) setStruct(structRef reflect.Value, sourceValue interface{}) error {
	// Convert the source value to a map and check if there is all okay.
	sourceMap, ok := sourceValue.(map[string]interface{})
	if !ok {
		// Convert the source value to a map and check if there is all okay.
		rawSourceMap, ok := sourceValue.(map[interface{}]interface{})
		if !ok {
			return errors.Wrap(bridgeErrors.ErrNotValid, fmt.Sprintf("can't cast to map %s", reflect.TypeOf(sourceValue).String()))
		}

		// Create a new source map and set the value of the source map.
		sourceMap = make(map[string]interface{})
		for key, value := range rawSourceMap {
			sourceMap[fmt.Sprintf("%v", key)] = value
		}
	}

	// Iterate over the struct reference and set the value of the struct reference.
	for i := 0; i < structRef.NumField(); i++ {
		// Get the field reference of the struct reference.
		fieldRef := structRef.Field(i)

		// Get the field tag of the struct reference and check if there is an error.
		tag, err := parseFieldTag(structRef.Type().Field(i), keyTag)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("can't parse field tag for field %s", structRef.Type().Field(i).Name))
		}

		// Check if the tag is nil.
		if tag == nil {
			continue
		}

		// Get the value of the source map and check if there is a value.
		value, hasValue := sourceMap[tag.Key]
		if hasValue {
			// Set the value of the field reference.
			if err := c.setValue(fieldRef, value); err != nil {
				return errors.Wrap(err, fmt.Sprintf("can't set field %s", tag.Key))
			}
		}

		// Check if the tag is required and if there is a value.
		if tag.Required && !hasValue {
			return errors.Wrap(bridgeErrors.ErrRequiredValue, fmt.Sprintf("required field '%s' is missing", tag.Key))
		}

		// Check if the tag is non-zero and if the field reference is zero.
		if tag.NonZero && isZero(fieldRef) {
			return errors.Wrap(bridgeErrors.ErrNonZeroValue, fmt.Sprintf("field %s has a zero value", tag.Key))
		}
	}

	return nil
}
