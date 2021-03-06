package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

// ParseFromQueryParam parse url query string to struct target (with multiple data type in struct field), target must in pointer
func ParseFromQueryParam(query url.Values, target interface{}) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()

	var parseDataTypeValue func(typ reflect.Type, val reflect.Value)

	var errs = NewMultiError()

	pValue := reflect.ValueOf(target)
	if pValue.Kind() != reflect.Ptr {
		panic(fmt.Errorf("%v is not pointer", pValue.Kind()))
	}
	pValue = pValue.Elem()
	pType := reflect.TypeOf(target).Elem()
	for i := 0; i < pValue.NumField(); i++ {
		field := pValue.Field(i)
		typ := pType.Field(i)
		if typ.Anonymous { // embedded struct
			if e, ok := ParseFromQueryParam(query, field.Addr().Interface()).(*MultiError); ok {
				errs.Merge(e)
			}
		}

		key := strings.TrimSuffix(typ.Tag.Get("json"), ",omitempty")
		if key == "-" {
			continue
		}

		var v string
		if val := query[key]; len(val) > 0 && val[0] != "" {
			v = val[0]
		} else {
			v = typ.Tag.Get("default")
		}

		parseDataTypeValue = func(sourceType reflect.Type, targetField reflect.Value) {
			switch sourceType.Kind() {
			case reflect.String:
				if ok, _ := strconv.ParseBool(typ.Tag.Get("lower")); ok {
					v = strings.ToLower(v)
				}
				targetField.SetString(v)
			case reflect.Int32, reflect.Int, reflect.Int64:
				vInt, err := strconv.Atoi(v)
				if v != "" && err != nil {
					errs.Append(key, fmt.Errorf("Cannot parse '%s' (%T) to type number", v, v))
				}
				targetField.SetInt(int64(vInt))
			case reflect.Bool:
				vBool, err := strconv.ParseBool(v)
				if v != "" && err != nil {
					errs.Append(key, fmt.Errorf("Cannot parse '%s' (%T) to type boolean", v, v))
				}
				targetField.SetBool(vBool)
			case reflect.Ptr:
				if v != "" {
					// allocate new value to pointer targetField
					targetField.Set(reflect.ValueOf(reflect.New(sourceType.Elem()).Interface()))
					parseDataTypeValue(sourceType.Elem(), targetField.Elem())
				}
			}
		}

		parseDataTypeValue(field.Type(), field)
	}

	if errs.HasError() {
		return errs
	}

	return
}

// GenerateHMAC to encrypt string
func GenerateHMAC(str string) string {
	key := []byte("1983671-12093847-002972-12332")
	h := hmac.New(sha256.New, key)
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// IsEmail check format is email or not
func IsEmail(str string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(str)
}

// TryCatch helper
func TryCatch(try func(), catch func(err error)) {
	defer func() {
		if r := recover(); r != nil {
			catch(fmt.Errorf("%v", r))
		}
	}()

	try()
}

// ToBoolPtr helper
func ToBoolPtr(b bool) *bool {
	return &b
}

// ToStringPtr helper
func ToStringPtr(str string) *string {
	return &str
}

// IsValidPassword helper
func IsValidPassword(pass string) bool {
	var uppercase, lowercase, num bool
	for _, r := range pass {
		if r >= 65 && r <= 90 { //code ascii for [A-Z]
			uppercase = true
		} else if r >= 97 && r <= 122 { //code ascii for [a-z]
			lowercase = true
		} else if r >= 48 && r <= 57 { //code ascii for [0-9]
			num = true
		}
	}
	return uppercase && lowercase && num
}
